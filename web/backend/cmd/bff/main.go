package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"web-backend/loader"
	"web-backend/resolver"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"
)

func main() {
	// setup and start GraphQL server
	log.Println("Starting GraphQL web server")
	bstr, err := ioutil.ReadFile("schema/schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	s := graphql.MustParseSchema(string(bstr), &resolver.RootResolver{})
	http.Handle("/query",
		cors.Default().Handler(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					// create dataloader with short-lived cache valid within a http request
					ctx := context.WithValue(context.Background(), "loaders", loader.GetDataLoaders())
					handler := &relay.Handler{Schema: s}
					handler.ServeHTTP(w, r.WithContext(ctx))
				},
			),
		),
	)
	http.Handle("/", http.FileServer(http.Dir("./playground")))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
