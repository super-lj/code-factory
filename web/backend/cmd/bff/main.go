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
	// init thrift client
	// log.Println("Initing Thrift CIBackend Client")
	// err := loader.InitThriftCIClient("localhost:9090")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// setup and start GraphQL server
	log.Println("Starting GraphQL web server")
	bstr, err := ioutil.ReadFile("schema/schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	s := graphql.MustParseSchema(string(bstr), &resolver.RootResolver{})
	http.Handle("/query", cors.Default().Handler(&relay.Handler{Schema: s}))
	http.Handle("/", http.FileServer(http.Dir("./playground")))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
