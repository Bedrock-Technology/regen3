/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Bedrock-Technology/regen3/config"
	"github.com/Bedrock-Technology/regen3/log"
	"github.com/Bedrock-Technology/regen3/models"
	s "github.com/Bedrock-Technology/regen3/scanner"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/sirupsen/logrus"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// completeQueueWithdrawalCmd represents the completeQueueWithdrawal command
var completeQueueWithdrawalCmd = &cobra.Command{
	Use:   "completeQueueWithdrawal",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("completeQueueWithdrawal called")
		root := cmd.Flag("root").Value.String()
		asToken := cmd.Flag("asToken").Value.String()
		if root == "" || asToken == "" {
			fmt.Println("root or asToken nil")
			return
		}
		configPath := cmd.Flag("config").Value.String()
		logPath := cmd.Flag("log").Value.String()
		config := config.LoadConfig(configPath)
		err := log.Init(config.LogLevel, logPath, fmt.Sprintf("%s/%s/%s", cmd.Root().Name(), config.Network, cmd.Name()), config.SlackUrl)
		if err != nil {
			return
		}
		//force confirm
		fmt.Printf("root: %v, asToken: %s press YES to continue\n", root, asToken)
		confirm := ""
		fmt.Scanln(&confirm)
		if confirm != "YES" {
			return
		}
		sigs := make(chan os.Signal, 1)
		quit := make(chan struct{}, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
		go func() {
			for sig := range sigs {
				switch sig {
				case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP:
					logrus.Infoln("Get Signal", sig)
					close(quit)
					return
				default:
					logrus.Infoln("Get Signal ", sig)
				}
			}
		}()
		scanner := s.New(config, quit)
		queueWithdrawalsNotCompleted := make([]models.QueueWithdrawals, 0)
		rest := scanner.DBEngine.Model(&models.QueueWithdrawals{}).Where("completed = ?", 0).
			Where("withdrawal_root = ?", root).Limit(1).Find(&queueWithdrawalsNotCompleted)
		if rest.Error != nil {
			logrus.Errorf("Get QueueWithdrawals failed: %v", rest.Error)
			return
		}
		if len(queueWithdrawalsNotCompleted) != 0 {
			pod := make([]models.Pod, 0)
			rest := scanner.DBEngine.Model(&models.Pod{}).Where("address = ?", queueWithdrawalsNotCompleted[0].Pod).Limit(1).Find(&pod)
			if rest.Error != nil {
				logrus.Errorf("Get QueueWithdrawals failed: %v", rest.Error)
				return
			}
			queue := queueWithdrawalsNotCompleted[0]
			withdrawal := s.DelegationManagerWithdrawalQueued{}
			err := json.Unmarshal([]byte(queue.Withdrawal), &withdrawal.Withdrawal)
			if err != nil {
				logrus.Errorf("Unmarshal QueueWithdrawals failed: %v", err)
				return
			}
			token := false
			if asToken == "true" {
				token = true
			}
			realTx, err := scanner.SendCompleteQueuedWithdrawals(pod[0], withdrawal, token)
			if err != nil {
				logrus.Errorf("SendCompleteQueuedWithdrawals failed: %v", err)
				return
			}
			logrus.Infoln("waiting sendCompleteQueuedWithdrawals tx:", realTx.Hash())
			txReceipt, err := bind.WaitMined(context.Background(), scanner.EthClient, realTx)
			if err != nil {
				logrus.Errorf("waiting sendCompleteQueuedWithdrawals index %v error:%v", pod[0].Address, err)
				panic("waiting error")
			}
			logrus.WithField("Report", "true").Infof("sendCompleteQueuedWithdrawals tx:%s", txReceipt.TxHash)
			//write to db
			fee := big.NewInt(0).Mul(txReceipt.EffectiveGasPrice, big.NewInt(int64(txReceipt.GasUsed)))
			txRecord := models.Transaction{
				TxHash: txReceipt.TxHash.String(),
				Status: txReceipt.Status,
				TxType: s.TxCompleteQueueWithdrawals,
				Fee:    fee.String(),
			}
			rest = scanner.DBEngine.Create(&txRecord)
			if rest.Error != nil {
				logrus.Errorf("Insert txRecord[%s] err:%v", txReceipt.TxHash.String(), rest.Error)
			}
			if txReceipt.Status != 1 {
				logrus.Errorf("sendCompleteQueuedWithdrawals tx: %v status failed:%v", txReceipt.TxHash, txRecord.Status)
				panic("sendQueueWithdrawals")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completeQueueWithdrawalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeQueueWithdrawalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeQueueWithdrawalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	completeQueueWithdrawalCmd.PersistentFlags().StringP("root", "r", "", "root")
	completeQueueWithdrawalCmd.PersistentFlags().StringP("asToken", "t", "", "asToken")
}
