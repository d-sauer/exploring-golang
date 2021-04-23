package main

// import "fmt"

// var V int

// func F() { fmt.Printf("Hello, number %d\n", V) }

import (
	"fmt"

	ext "github.com/d-sauer/exploring-go/cli-sandbox/sandbox/extension"
	"github.com/spf13/cobra"
)

// hiCmd represents the hi command
var hiCmd = &cobra.Command{
	Use:   "hi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hi called")
	},
}

func init() {
	fmt.Println("Cobra `hello` init")
	ext.RegisterPlugin(hiCmd)
}
