package api

import (
	"api_crud/config"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type Days struct {
	Day      string
	Friday   time.Time
	Saturday time.Time
	Sunday   time.Time
}

func (server *Server) runCron(c **cron.Cron, config *config.API) {
	// ctx := context.Background()
	(*c) = cron.New()

	// (*c).AddFunc("@every 5s", func() {
	// 	test, err := server.Store.GetBooks(ctx)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(test)
	// })

	(*c).Start()
	log.Printf("%+v\n", (*c).Entries())
}
