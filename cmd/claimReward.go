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
	"github.com/Bedrock-Technology/regen3/scanner"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// delegateToCmd represents the delegateTo command
var claimRewardCmd = &cobra.Command{
	Use:   "claimReward",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("claimReward called")
		configPath := cmd.Flag("config").Value.String()
		logPath := cmd.Flag("log").Value.String()
		config := config.LoadConfig(configPath)
		err := log.Init(config.LogLevel, logPath, fmt.Sprintf("%s/%s/%s", cmd.Root().Name(), config.Network, cmd.Name()), config.SlackUrl)
		if err != nil {
			return
		}
		podIndex := cmd.Flag("podIndex").Value.String()
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
		sc := scanner.New(config, quit)

		if podIndex != "" {
			podIndexInt, err := strconv.ParseInt(podIndex, 0, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			fileName := fmt.Sprintf("pod_%d.json", podIndexInt)
			proofs, podIndexes := scanner.GetProofs([]string{fmt.Sprintf("%s/%s", scanner.RewardProofDir, fileName)})
			proofs, podIndexes = scanner.FilterProofs(proofs, podIndexes, sc)
			// force confirm
			fmt.Printf("claimReward pods:%v, press YES to continue\n", podIndexes)
			confirm := ""
			fmt.Scanln(&confirm)
			if confirm != "YES" {
				return
			}
			sc.ClaimRewardWithProof(proofs, podIndexes)
		} else {
			// force confirm
			fmt.Println("claimReward all pods with filter, press YES to continue")
			confirm := ""
			fmt.Scanln(&confirm)
			if confirm != "YES" {
				return
			}
			sc.ClaimReward()
		}
	},
}

func init() {
	rootCmd.AddCommand(claimRewardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delegateToCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delegateToCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	claimRewardCmd.PersistentFlags().StringP("podIndex", "p", "", "pod")
}
