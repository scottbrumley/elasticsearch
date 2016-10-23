package main

import (
	"fmt"
	"github.com/scottbrumley/elasticsearch"
	"strings"
)

func main() {
	fmt.Println("Beging Testing")
	fmt.Println("")
	myParms := elasticsearch.GetParams()

	fmt.Println("     Test Connection ...")
	resp, _ := elasticsearch.ConnectES(myParms)
	if ( strings.Contains(resp.Status,"200") ){
		fmt.Println("     Connection Success")
		fmt.Println("")
		//fmt.Println(respStr)
	} else {
		fmt.Println("     Connection Failed " + resp.Status)
	}

	indexStr := "test"
	// Create Index
	if (elasticsearch.IndexExists(myParms,indexStr) == false) {
		resp, _ = elasticsearch.AddIndex(myParms, indexStr)
		fmt.Println("Add Index " + indexStr + " Status " + string(resp.Status))
	}

	// Test Index Exists
	fmt.Println("     Test Index Exists ...")
	if (elasticsearch.IndexExists(myParms,indexStr)){
		fmt.Println("     Index " + indexStr + " exists Status " + string(resp.Status))
	} else {
		fmt.Println("     Index " + indexStr + " does not exists Status " + string(resp.Status))
	}

	// Delete Index
	resp, _ = elasticsearch.DeleteIndex(myParms,indexStr)
	fmt.Println("Delete Index " + indexStr + " Status " + string(resp.Status))
}
