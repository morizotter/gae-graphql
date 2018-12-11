package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"cloud.google.com/go/profiler"
	"github.com/monmaru/gae-graphql/application/gql"
	"github.com/monmaru/gae-graphql/infrastructure/datastore"
	"github.com/monmaru/gae-graphql/interfaces/router"
)

func main() {
	if err := register(); err != nil {
		log.Fatal(err)
	}
}

func register() error {
	// Stackdriver Profiler
	if enabled, _ := strconv.ParseBool(os.Getenv("PROFILE_ENABLED")); enabled {
		if err := profiler.Start(profiler.Config{DebugLogging: true}); err != nil {
			return err
		}
	}

	projID := os.Getenv("PROJECT_ID")
	ud := &datastore.UserDatastore{ProjID: projID}
	bd := &datastore.BlogDatastore{ProjID: projID}
	schema, err := gql.NewSchema(ud, bd)
	if err != nil {
		return err
	}
	router := router.Route(&schema)
	http.Handle("/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
