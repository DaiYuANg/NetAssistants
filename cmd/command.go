package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var app = fx.New(
	fx.Provide(zap.NewExample),
	schedulerModule,
	databaseModule,
	contextModule,
	errorModule,
	uiModule,
)

var RootCmd = cobra.Command{Use: "ProtocolAssistant",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Start(context.Background())
	}}

func Execute() error {
	return RootCmd.Execute()
}
