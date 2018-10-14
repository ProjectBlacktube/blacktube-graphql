package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/99designs/gqlgen/handler"
	"github.com/gobuffalo/pop"
	gql "github.com/koneko096/blacktube-graphql/graphql"
	"github.com/koneko096/blacktube-graphql/manager"
)

func New() gql.Config {
	env := os.Getenv("APP_ENV")
	if len(env) == 0 {
		env = "development"
	}

	db, err := pop.Connect(env)
	if err != nil {
		log.Panic(err)
	}

	userManager := &manager.UserQueryManager{
		Db: db,
	}
	videoManager := &manager.VideoQueryManager{
		Db:          db,
		UserManager: userManager,
	}

	return gql.Config{
		Resolvers: &gql.Resolver{
			UserManager:  userManager,
			VideoManager: videoManager,
		},
	}
}

func main() {
	http.Handle("/", handler.Playground("Blacktube", "/query"))
	http.Handle("/query", handler.GraphQL(
		gql.NewExecutableSchema(New()),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			// send this panic somewhere
			log.Print(err)
			debug.PrintStack()
			return errors.New("user message on panic")
		}),
	))
	log.Fatal(http.ListenAndServe(":8083", nil))
}
