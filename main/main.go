package main

// import (
// 	"api_crud/api"
// 	"api_crud/config"
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	cnf := config.Get()
// 	source := fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=disable TimeZone=%s", cnf.Database.User, cnf.Database.Password, cnf.Database.Host, cnf.Database.Port, cnf.Database.Database, cnf.TZ)
// 	pg, err := sql.Open("postgres", source)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	store := api.NewStore(pg)
// 	test, _ := store.GetBooks(context.Background())
// 	fmt.Println(test)
// }
