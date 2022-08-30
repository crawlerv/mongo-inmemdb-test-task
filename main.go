package main

import (
	"flag"
	"github.com/crawlerv/mongo-inmemdb-test-task/cmd"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/resources"
	"github.com/crawlerv/mongo-inmemdb-test-task/pkg/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "tg-notifier",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		flag.CommandLine.Parse([]string{})
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
	rc := &cmd.RootCommand{}
	conf := resources.GetConfig()
	log := logger.Init(conf.Log.Dir, conf.Log.File, conf.Log.Encoding)

	rc = &cmd.RootCommand{
		Log:    log,
		Config: conf,
	}
	rc.Execute()
}

func init() {
	resources.RegisterConfigFlag(rootCmd)
}
