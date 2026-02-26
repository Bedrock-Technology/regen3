/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
var fillCpProofCmd = &cobra.Command{
	Use:   "fillCpProof",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fillCpProof called")
		configPath := cmd.Flag("config").Value.String()
		logPath := cmd.Flag("log").Value.String()
		config := config.LoadConfig(configPath)
		err := log.Init(config.LogLevel, logPath, fmt.Sprintf("%s/%s/%s", cmd.Root().Name(), config.Network, cmd.Name()), config.SlackUrl)
		if err != nil {
			return
		}
		podId := cmd.Flag("podId").Value.String()
		if podId == "" {
			fmt.Println("podId nil")
			return
		}
		podIdInt, err := strconv.ParseInt(podId, 0, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		timestamp := cmd.Flag("timestamp").Value.String()
		if timestamp == "" {
			fmt.Println("timestamp nil")
			return
		}
		timestampInt, err := strconv.ParseInt(timestamp, 0, 64)
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
		rest := scanner.DBEngine.First(&pod, "id = ?", podId)
		if rest.Error != nil {
			logrus.Errorln("Get Pod Error", rest.Error)
			return
		}
		// force confirm
		fmt.Printf("podId: %v, podAddress:%v, timestamp: %v press YES to continue\n",
			podIdInt, pod.Address, timestampInt)
		confirm := ""
		fmt.Scanln(&confirm)
		if confirm != "YES" {
			return
		}
		logrus.Infof("need do FillProofs pod %v timestamp:%v", pod.Address, timestampInt)
		proofs, err := scanner.FillProofs(pod.Address, uint64(timestampInt))
		if err != nil {
			logrus.Errorf("FillProofs pod %v timestamp:%v", pod.Address, timestampInt)
			panic("FillProofs")
		}
		// write to db
		rest = scanner.DBEngine.Model(&models.CheckPoint{}).Where("pod = ?", pod.Address).
			Where("checkpoint_timestamp = ?", timestampInt).Where("checkpoint_finalized = ?", 0).
			Update("proofs", string(proofs))
		if rest.Error != nil {
			logrus.Errorf("update pod %v timestamp:%v", pod.Address, timestamp)
			panic("update")
		}
	},
}

func init() {
	rootCmd.AddCommand(fillCpProofCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpstartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpstartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	fillCpProofCmd.PersistentFlags().StringP("podId", "p", "", "pod")
	fillCpProofCmd.PersistentFlags().StringP("timestamp", "t", "", "timestamp")
}
