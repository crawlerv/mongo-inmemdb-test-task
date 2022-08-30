package cmd

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/db/mongo"
	"github.com/crawlerv/mongo-inmemdb-test-task/internal/resources"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"strconv"
	"time"
)

type FakerCommand struct {
	App resources.App
	R   mongo.Repo
}

func (fc *FakerCommand) RegisterCommand(parent *cobra.Command) {
	parent.AddCommand(&cobra.Command{
		Use: "faker",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Printf("Need one integer argument. error\n")
				return
			}
			count, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Need one integer argument. error - %v\n", err)
				return
			}
			rand.Seed(time.Now().UnixNano())
			m := make([]interface{}, 0, count)
			for i := 0; i < count; i++ {
				mm := &mongo.Model{}
				err = faker.FakeData(mm)
				if err != nil {
					fc.App.Log().Error(err)
					return
				}
				mm.CardNumber = uint32(rand.Intn(4294967295))
				mm.Age = uint8(rand.Intn(150))
				v := rand.Intn(1)
				if v == 0 {
					mm.Verified = false
				} else {
					mm.Verified = true
				}
				mm.ID = primitive.NewObjectID()
				m = append(m, mm)
			}
			err = fc.R.InsertMany(m)
			if err != nil {
				fc.App.Log().Error(err)
				return
			}
			fc.App.Log().Infof("Inserted %d records", count)
		},
	})
}
