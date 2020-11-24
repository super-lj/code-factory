package main

import (
	"ci-backend/actor"
	"ci-backend/config"
	"ci-backend/dao"
	"ci-backend/handler"
	"ci-backend/image"
	"log"
	"os"
	"time"
)

func main() {
	// read the yml file
	log.Println("Reading config file...")
	if len(os.Args) < 2 {
		log.Fatalln("Please specify config file path.")
	}
	err := config.InitCIConfig(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Reading config file complete!")

	// init watched git repos
	log.Println("Initing watched repos...")
	err = config.InitWatchedRepos()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Watched repos init complete!")

	// init DB
	log.Println("Initing DB...")
	err = dao.RunDB.AutoMigrate(&dao.Run{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB init complete!")

	// build the CFCI Docker image
	log.Println("Initing Docker image...")
	err = image.InitDockerImages()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Docker image init complete!")

	// start repo observers
	log.Println("Starting Repo Observers...")
	repoObChs := []chan string{}
	for repoName := range config.WatchedRepos {
		obCh, err := actor.StartRepoObserver(repoName)
		if err != nil {
			log.Fatalf("Start Repo OB failed: %v", err)
		}
		repoObChs = append(repoObChs, obCh)
	}
	go func() {
		for {
			for _, ch := range repoObChs {
				select {
				case msg := <-ch:
					log.Printf("Repo ob failed: %s", msg)
				default:
				}
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
	log.Println("Repo Observers started!")

	// start thrift server
	log.Println("Starting Thrift Server")
	if err := handler.InitThriftServer("localhost:9090").Serve(); err != nil {
		log.Fatalf("Cannot start thrift server: %v", err)
	}
}
