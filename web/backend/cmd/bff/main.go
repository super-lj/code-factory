package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"web-backend/resolver"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"
)

func main() {
	// read and setup GraphQL schema and resolver
	bstr, err := ioutil.ReadFile("schema/schema.graphql")
	if err != nil {
		panic(err)
	}
	s := graphql.MustParseSchema(string(bstr), &resolver.RootResolver{})
	http.Handle("/query", cors.Default().Handler(&relay.Handler{Schema: s}))

	// register graphql playground handler
	http.Handle("/", http.FileServer(http.Dir("./playground")))

	// start web server
	log.Println("Starting web server")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
