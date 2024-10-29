/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Bedrock-Technology/regen3/config"
	"github.com/Bedrock-Technology/regen3/log"
	"github.com/Bedrock-Technology/regen3/models"
	"github.com/Bedrock-Technology/regen3/scanner"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// cpstartCmd represents the cpstart command
var cpstartCmd = &cobra.Command{
	Use:   "cpstart",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cpstart called")
		configPath := cmd.Flag("config").Value.String()
		logPath := cmd.Flag("log").Value.String()
		config := config.LoadConfig(configPath)
		err := log.Init(config.LogLevel, logPath, fmt.Sprintf("%s/%s/%s", cmd.Root().Name(), config.Network, cmd.Name()), config.SlackUrl)
		if err != nil {
			return
		}
		podIndex := cmd.Flag("podIndex").Value.String()
		if podIndex == "" {
			fmt.Println("podIndex nil")
			return
		}
		podIndexInt, err := strconv.ParseInt(podIndex, 0, 64)
		if err != nil {
			fmt.Println(err)
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
		scanner := scanner.New(config, quit)
		var pod models.Pod
		rest := scanner.DBEngine.First(&pod, "pod_index = ?", podIndex)
		if rest.Error != nil {
			logrus.Errorln("Get Pod Error", rest.Error)
			return
		}
		// force confirm
		fmt.Printf("podIndex: %v, podOwner:%v, podAddress:%v press YES to continue\n",
			podIndexInt, pod.Owner, pod.Address)
		confirm := ""
		fmt.Scanln(&confirm)
		if confirm != "YES" {
			return
		}
		// send startCheckPoint
		timestamp, err := scanner.SendCheckPoint(big.NewInt(int64(pod.PodIndex)), pod.Address)
		if err != nil {
			logrus.Errorf("send checkpoint pod %v error:%v", pod.Address, err)
			panic("send checkpoint err")
		}
		if timestamp != 0 {
			logrus.Infof("need do FillProofs pod %v timestamp:%v", pod.Address, timestamp)
			proofs, err := scanner.FillProofs(pod.Address, timestamp)
			if err != nil {
				logrus.Errorf("FillProofs pod %v timestamp:%v", pod.Address, timestamp)
				panic("FillProofs")
			}
			// write to db
			rest := scanner.DBEngine.Model(&models.CheckPoint{}).Where("pod = ?", pod.Address).
				Where("checkpoint_timestamp = ?", timestamp).Where("checkpoint_finalized = ?", 0).
				Update("proofs", string(proofs))
			if rest.Error != nil {
				logrus.Errorf("update pod %v timestamp:%v", pod.Address, timestamp)
				panic("update")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cpstartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpstartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpstartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cpstartCmd.PersistentFlags().StringP("podIndex", "p", "", "pod")
}
