package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var (
	serverHost  string
	labProvider string
	configPath  string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&serverHost, "server", "s", "http://localhost:4188", "")
	rootCmd.PersistentFlags().StringVarP(&labProvider, "lab", "l", "", "")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "")
}

var rootContext context.Context

var rootCmd = &cobra.Command{
	Use:   "tenderbytes-lab",
	Short: "`tenderbytes-lab`",
	Long:  "`tenderbytes-lab`",
}

func Execute(ctx context.Context) {
	SetContext(ctx)
	rootCmd.Execute()
}

func SetContext(ctx context.Context) {
	rootContext = ctx
}
