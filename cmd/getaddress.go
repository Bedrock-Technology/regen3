/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Bedrock-Technology/regen3/config"
	"github.com/Bedrock-Technology/regen3/keyAgentRpc"
	"github.com/spf13/cobra"
)

// getaddressCmd represents the getaddress command
var getaddressCmd = &cobra.Command{
	Use:   "getaddress",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getaddress called")
		configPath := cmd.Flag("config").Value.String()
		config := config.LoadConfig(configPath)
		client := keyAgentRpc.NewClient(config.KeyAgent)
		address, err := client.Eth1Address()
		if err != nil {
			fmt.Println("Eth1Address error:", err)
			return
		}
		fmt.Printf("service[%d], index[%d], address[%s]", config.KeyAgent.Service, config.KeyAgent.Index, address)
	},
}

func init() {
	rootCmd.AddCommand(getaddressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getaddressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getaddressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
