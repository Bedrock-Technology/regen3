/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Bedrock-Technology/regen3/config"
	"github.com/Bedrock-Technology/regen3/log"
	"github.com/Bedrock-Technology/regen3/scanner"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var getPodsCmd = &cobra.Command{
	Use:   "getpods",
	Short: "getpods",
	Long:  `getpods`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getpods called")
		configPath := cmd.Flag("config").Value.String()
		logPath := cmd.Flag("log").Value.String()
		config := config.LoadConfig(configPath)
		err := log.Init(config.LogLevel, logPath, fmt.Sprintf("%s/%s/%s", cmd.Root().Name(), config.Network, cmd.Name()), config.SlackUrl)
		if err != nil {
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
		scanner.GetAllPod()
	},
}

func init() {
	rootCmd.AddCommand(getPodsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
