package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/crawlerv/mongo-inmemdb-test-task/config"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/controllers"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/db/mongo"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/resources"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/service"
	"github.com/crawlerv/mongo-inmemdb-test-task/pkg/logger"
	"github.com/spf13/cobra"
	"os"
)

type RootCommand struct {
	Log     *logger.Logger
	Config  *config.Config
	command *cobra.Command
}

var rootCmd = &cobra.Command{
	Use: "test-task",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		flag.CommandLine.Parse([]string{})
	},
}

func (rc *RootCommand) Execute() {
	ctx := context.Background()
	app := resources.New(ctx)

	repo := mongo.NewRepo(ctx, app)
	controller := controllers.NewController(app, service.New(app, repo))
	sc := ServeCommand{
		App: app,
		C:   controller,
	}
	sc.RegisterCommand(rootCmd)
	fc := FakerCommand{
		App: app,
		R:   repo,
	}
	fc.RegisterCommand(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	resources.RegisterConfigFlag(rootCmd)
}
