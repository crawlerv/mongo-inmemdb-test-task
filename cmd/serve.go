package cmd

import (
	"context"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/controllers"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/db/mongo"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/resources"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/service"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
	"net/http"
)

type ServeCommand struct {
	App resources.App
	C   controllers.Controller
}

func (mc *ServeCommand) RegisterCommand(parent *cobra.Command) {
	parent.AddCommand(&cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			mc.App.Log().Infof("Running serve command")

			memCacheService := service.NewMemoryCache()
			ticker := service.NewTicker(mc.App, mongo.NewRepo(context.Background(), mc.App), memCacheService)
			go ticker.LoadData(context.Background())

			router := httprouter.New()
			router.GET("/list", mc.C.List)
			router.POST("/create", mc.C.Create)
			router.POST("/update/:id", mc.C.Update)
			router.POST("/delete/:id", mc.C.Delete)
			mc.App.Log().Fatal(http.ListenAndServe(":8888", router))

		},
	})
}
