package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	"github.com/icalF/blacktube-graphql/graphql"
	"github.com/joho/godotenv"
	gqlhandler "github.com/sogko/graphql-go-handler"
)

func main() {
	// Load from file. If not exist, use environment variables only
	godotenv.Load()

	// simplest graphql server HTTP handler
	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &graphql.Schema,
		Pretty: true,
	})

	// create graphql endpoint
	http.Handle("/graphql", h)

	// serve!
	port := fmt.Sprintf(":%v", os.Getenv("APP_PORT"))
	addr := fmt.Sprintf(`%v%v`, os.Getenv("APP_URL"), port)
	log.Printf(`GraphQL server starting up on %v`, addr)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed, %v", addr)
	}
}
