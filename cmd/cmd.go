package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"test_websocket/client"
	"test_websocket/server"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "start and stop",
	Long:  "input cmd help for werewolf cmd list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start server or client",
	Long:  `start websocket server or client`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			startServerOrClient(args[0])
		} else {
			fmt.Println("lack of args")
		}
	},
}

func startServerOrClient(name string) {
	switch name {
	case "server":
		server.StartWebsocketServe()
	case "client":
		client.StartWebsocketClient()
	}
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
