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

	indexStr := "wink"
	fmt.Println("     Test Index Exists ...")
	if (elasticsearch.IndexExists(myParms,indexStr)){
		fmt.Println("     Index " + indexStr + " exists")
	} else {
		fmt.Println("     Index " + indexStr + " does not exists")
	}
}
