#Elastic Search Go Program

## Setup GOPATH
```
export GOPATH=<project base directory>
```

## RUN 
```
go run main.go -url http://localhost:9200
```

## Build Package
```
./scripts/cibuild.sh
```

## Example Code
```
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


	// Add Document: Read test document file and add document to Elastic Search
	b, err := ioutil.ReadFile("./test/test_rec") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	str = strings.Replace(str, " ","",-1)
	str = strings.Replace(str, "\n","",-1)
	str = strings.TrimSpace(str)
	elasticsearch.AddDocument(myParms,indexStr,typeStr,recStr,str)

	// Get Document
	respStatus,respStr := elasticsearch.GetDocument(myParms,indexStr,typeStr,recStr)
	if (respStatus.StatusCode == 200) {
		fmt.Println(respStr)
	} else {
		fmt.Println("Document Not found status code: " + respStatus.Status)
	}

	//Update Document
	// Read test document file and add document to Elastic Search
	b, err = ioutil.ReadFile("./test/updatetest") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str = string(b) // convert content to a 'string'
	str = strings.Replace(str, " ","",-1)
	str = strings.Replace(str, "\n","",-1)
	str = strings.TrimSpace(str)
	elasticsearch.UpdateDocument(myParms,indexStr,typeStr,recStr,str)

	// Get Document
	respStatus,respStr = elasticsearch.GetDocument(myParms,indexStr,typeStr,recStr)
	if (respStatus.StatusCode == 200) {
		fmt.Println(respStr)
	} else {
		fmt.Println("Document Not found status code: " + respStatus.Status)
	}

	// Delete Document
	respStatus,respStr = elasticsearch.DeleteDocument(myParms,indexStr,typeStr,recStr)
	if (respStatus.StatusCode == 200) {
		fmt.Println(respStr)
	} else {
		fmt.Println("Document Not found status code: " + respStatus.Status)
	}

	// Get Document
	respStatus,respStr = elasticsearch.GetDocument(myParms,indexStr,typeStr,recStr)
	if (respStatus.StatusCode == 200) {
		fmt.Println(respStr)
	} else {
		fmt.Println("Document Not found status code: " + respStatus.Status)
	}


	// Delete Index
	respStatus,respStr = elasticsearch.DeleteIndex(myParms,indexStr)
	if (respStatus.StatusCode == 200) {
		fmt.Println(respStr)
	} else {
		fmt.Println("Delete Index Failed: " + respStatus.Status)
	}

}

```