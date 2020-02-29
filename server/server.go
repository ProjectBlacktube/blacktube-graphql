package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/ProjectBlacktube/blacktube-graphql/graphql"
	"github.com/ProjectBlacktube/blacktube-graphql/manager"
	"github.com/gobuffalo/pop"
)

func New() graphql.Config {
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

	return graphql.Config{
		Resolvers: &graphql.Resolver{
			UserManager:  userManager,
			VideoManager: videoManager,
		},
	}
}

func main() {
	http.Handle("/", playground.Handler("Blacktube", "/version"))
	http.Handle("/graphql", handler.NewDefaultServer(graphql.NewExecutableSchema(New())))
	//handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
	//	// send this panic somewhere
	//	log.Print(err)
	//	debug.PrintStack()
	//	return errors.New("user message on panic")
	//}),
	log.Fatal(http.ListenAndServe(":8080", nil))
}
