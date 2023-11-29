package main

import (
	"github.com/andyzhou/tinysearch"
	"log"
	"os"
	"sync"
)

/*
 * service example
 */

const (
	RpcPort = 6060
	IndexPath = "./search_data"
	IndexTag = "test"
	SuggesterTag = "test"
)

//hook for add doc
func hookForAddDoc(jsonBytes []byte) error {
	log.Println("jsonBytes:", jsonBytes)
	return nil
}

func main() {
	var (
		wg sync.WaitGroup
	)

	//watch signal
	tinysearch.WatchSignal(&wg)

	//init service
	service := tinysearch.NewServiceWithPara(IndexPath, RpcPort)

	//set relate path
	service.SetDataPath(IndexPath)
	//service.SetDictFile("")

	//set hook for add doc
	service.SetHookForAddDoc(hookForAddDoc)

	//add index
	err := service.AddIndex(IndexTag)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	//register suggester tag
	service.GetSuggest().RegisterSuggest(SuggesterTag)

	//start wait group
	log.Printf("start server on port %v\n", RpcPort)
	wg.Add(1)
	wg.Wait()
	service.Quit()
	log.Printf("stop server on port %v\n", RpcPort)
}