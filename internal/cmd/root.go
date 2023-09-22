package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile = ""

var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
