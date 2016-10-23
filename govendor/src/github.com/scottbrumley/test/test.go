package main

import (
	"fmt"
	"github.com/scottbrumley/elasticsearch"
)

func main() {
	fmt.Println("Beging Testing")
	myParms := elasticsearch.GetParams()
	elasticsearch.ConnectES(myParms)
	indexStr := "wink"
	if (elasticsearch.IndexExists(myParms,indexStr)){
		fmt.Println("Index " + indexStr + " exists")
	} else {
		fmt.Println("Index " + indexStr + " does not exists")
	}

	/*
	myParms.Url = myParms.Url + "wink"
	resp := indexExists(myParms)

	if resp.Status == "404 Not Found"{
		fmt.Println("Index Not Found")
	}
	*/
}
