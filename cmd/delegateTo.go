/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Bedrock-Technology/regen3/config"
	"github.com/Bedrock-Technology/regen3/log"
	"github.com/Bedrock-Technology/regen3/scanner"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
)

// delegateToCmd represents the delegateTo command
var delegateToCmd = &cobra.Command{
	Use:   "delegateTo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delegateTo called")
		configPath := cmd.Flag("config").Value.String()
		logPath := cmd.Flag("log").Value.String()
		config := config.LoadConfig(configPath)
		err := log.Init(config.LogLevel, logPath, fmt.Sprintf("%s/%s/%s", cmd.Root().Name(), config.Network, cmd.Name()), config.SlackUrl)
		if err != nil {
			return
		}
		podIndex := cmd.Flag("podIndex").Value.String()
		operator := cmd.Flag("operator").Value.String()
		if podIndex == "" || operator == "" {
			fmt.Println("podIndex or operator nil")
			return
		}
		podIndexInt, err := strconv.ParseInt(podIndex, 0, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		//force confirm
		fmt.Printf("podIndex: %v, operator: %s press YES to continue\n", podIndexInt, operator)
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
		scanner := scanner.New(config, quit)
		scanner.DelegateTo(podIndexInt, operator)
	},
}

func init() {
	rootCmd.AddCommand(delegateToCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delegateToCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delegateToCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setupCmd.PersistentFlags().StringP("podIndex", "p", "", "pod")
	setupCmd.PersistentFlags().StringP("operator", "o", "", "operator")
}
