package main

import (
	"ci-backend/dao"
	"ci-backend/handler"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"gopkg.in/yaml.v2"
)

func main() {
	// read the yml file
	log.Println("Reading config file...")
	if len(os.Args) < 2 {
		log.Fatalln("Please specify config file path.")
	}
	bstr, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Cannot open config file: %v", err)
	}
	ymlMap := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(bstr), &ymlMap)
	if err != nil {
		log.Fatalf("Cannot unmarshal config file: %v", err)
	}
	log.Println("Reading config file complete!")

	// init DB
	log.Println("Initing DB...")
	dao.RunDB.AutoMigrate(&dao.Run{})
	log.Println("DB init complete!")

	// TODO: start repo observers
	log.Println("Starting Repo Observers...")
	for k, v := range ymlMap["repos"].(map[interface{}]interface{}) {
		repoName := k.(string)
		path := v.(map[interface{}]interface{})["path"].(string)
		repo, err := git.PlainOpen(path)
		if repo == nil || err != nil {
			log.Fatalf("Cannot open repo %v", err)
		}
		dao.WatchedRepos[repoName] = *repo
	}
	log.Println("Repo Observers started!")

	// start thrift server
	log.Println("Starting Thrift Server")
	if err := handler.InitThriftServer("localhost:9090").Serve(); err != nil {
		log.Fatalf("Cannot start thrift server: %v", err)
	}
}
