package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var (
	serverHost string
	configPath string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "")
}

var rootContext context.Context

var rootCmd = &cobra.Command{
	Use:   "kindermud",
	Short: "`kindermud`",
	Long:  "`kindermud`",
}

func Execute(ctx context.Context) {
	SetContext(ctx)
	rootCmd.Execute()
}

func SetContext(ctx context.Context) {
	rootContext = ctx
}
