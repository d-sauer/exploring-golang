package cmd

import (
	"fmt"
	"plugin"

	ex "github.com/d-sauer/exploring-go/cli-sandbox/sandbox/extension"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pluginsCmd)
	loadPlugins()
}

var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "List all available plugins",
	Long:  `List all available plugins for the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Plugins from .sandbox/plugins folder")
		for i, p := range ex.GetPlugins() {
			fmt.Println(i, p.Use)
		}
	},
}

func loadPlugins() {
	_, err := plugin.Open("/Users/davorsauer/development/repo/d-sauer/exploring-golang/cli-sandbox/sandbox-hello-plugin/bin/hello.so")
	if err != nil {
		panic(err)
	}

	for _, p := range ex.GetPlugins() {
		rootCmd.AddCommand(p)
	}
}
