package main

import (
	"fmt"
	"github.com/scottbrumley/elasticsearch"
	"strings"
	"io/ioutil"
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
	typeStr := "sub"
	recStr := "251105"
	bodyStr := "{\"mappings\": {\"" + typeStr + "\": {\"_timestamp\": {\"enabled\": true}}}}"
	// Create Index
	if (elasticsearch.IndexExists(myParms,indexStr) == false) {
		resp, _ = elasticsearch.AddIndex(myParms, indexStr, bodyStr)
		fmt.Println("Add Index " + indexStr + " Status " + string(resp.Status))
	}

	// Test Index Exists
	fmt.Println("     Test Index Exists ...")
	if (elasticsearch.IndexExists(myParms,indexStr)){
		fmt.Println("     Index " + indexStr + " exists Status " + string(resp.Status))
	} else {
		fmt.Println("     Index " + indexStr + " does not exists Status " + string(resp.Status))
	}


	// Read test document file and add document to Elastic Search
	b, err := ioutil.ReadFile("./govendor/src/github.com/scottbrumley/test/test_rec.js") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	str = strings.Replace(str, " ","",-1)
	str = strings.Replace(str, "\n","",-1)
	str = strings.TrimSpace(str)
	fmt.Println(str)
	elasticsearch.AddDocument(myParms,indexStr,typeStr,recStr,str)

	// Delete Index
	//resp, _ = elasticsearch.DeleteIndex(myParms,indexStr)
	//fmt.Println("Delete Index " + indexStr + " Status " + string(resp.Status))
}
