package elasticsearch

import (
	"flag"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"bytes"
)

// Parameters from command line
type ParamStruct struct{
	UserName string
	UserPass string
	SslIgnore bool
	Method string
	Test bool
	Url string
}

// Collect parameters from the command line
func GetParams()(retParams ParamStruct){

	var userFlag = flag.String("user","","ES Username")
	var passFlag = flag.String("password","","ES Password")
	var methodFlag = flag.String("method","GET","HTTP Method")
	var testFlag = flag.Bool("test",false,"Testing Mode")
	var urlFlag = flag.String("url","http://localhost:9200","URL for API")
	flag.Parse()

	retParams.UserName = *userFlag
	retParams.UserPass = *passFlag
	retParams.Method = *methodFlag
	retParams.Url = *urlFlag
	retParams.Test = *testFlag

	return retParams
}

// URL Fetching function
func getURL(myParms ParamStruct, bodyStr string)(res *http.Response, retStr string){
	var req *http.Request
	var err error
	bodyByte := []byte(bodyStr)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: myParms.SslIgnore},
	}

	client := &http.Client{
		Transport: tr,
	}

	if (bodyStr != ""){
		req, err = http.NewRequest(myParms.Method, myParms.Url, bytes.NewBuffer(bodyByte))
	} else {
		req, err = http.NewRequest(myParms.Method, myParms.Url, nil)
	}

	if ( (myParms.UserName != "") && (myParms.UserPass != "") ){
		req.SetBasicAuth(myParms.UserName, myParms.UserPass)
	}

	resp, err := client.Do(req)

	if err != nil {
		return res, ""
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		retStr = string(body)
		res = resp

		//data := decodeDevices(string(body))
		return res, retStr
	}
}

// Connect to elastic search and return results plus HTTP status
func ConnectES(myParms ParamStruct)(resp *http.Response, respStr string){
	resp, respStr = getURL(myParms,"")
	return resp, respStr
}

// Check that Index given exists and return true or false
func IndexExists(myParms ParamStruct, indexParm string)(bool){
	myParms.Url = myParms.Url + "/" + indexParm
	myParms.Method = "HEAD"
	resp, _ := getURL(myParms,"")
	if resp.Status == "404 Not Found"{
		return false
	} else {
		return true
	}
}

// Delete the Index given and return response Status and any response body as a string
func DeleteIndex(myParms ParamStruct, indexParm string)(resp *http.Response, respStr string){
	myParms.Url = myParms.Url + "/" + indexParm
	myParms.Method = "DELETE"
	resp, respStr = getURL(myParms,"")
	return resp, respStr
}

// Add the Index given and return response Status and any response body as a string
func AddIndex(myParms ParamStruct, indexParm string, jsonStr string)(resp *http.Response, respStr string){
	myParms.Url = myParms.Url + "/" + indexParm
	myParms.Method = "PUT"
	resp, respStr = getURL(myParms,jsonStr)
	return resp, respStr
}

func AddDocument(myParms ParamStruct, indexParm string, jsonStr string)(resp *http.Response, respStr string){
	//body := []byte("{\n  \"client_id\": \"" + myParms.ClientID + "\",\n  \"client_secret\": \"" + myParms.ClientSecret + "\",\n  \"username\": \"" + myParms.UserName + "\",\n  \"password\": \"" + myParms.UserPass + "\",\n  \"grant_type\": \"password\"\n}")
	myParms.Url = myParms.Url + "/" + indexParm
	myParms.Method = "PUT"
	resp, respStr = getURL(myParms,"")
	return resp, respStr
}
