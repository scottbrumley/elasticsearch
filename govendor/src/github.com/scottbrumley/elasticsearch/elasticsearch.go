package elasticsearch

import (
	"fmt"
	"os"
	"flag"
	"net/http"
	"io/ioutil"
	"crypto/tls"
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

func showSyntax(){
	fmt.Println("For Help " + os.Args[0] + " -h")
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

// Authenticate to Wink API and pull back tokens
func getURL(myParms ParamStruct)(res *http.Response, retStr string){
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: myParms.SslIgnore},
	}

	client := &http.Client{
		Transport: tr,
	}

	//fmt.Println(myParms.Url)
	req, err := http.NewRequest(myParms.Method, myParms.Url, nil)
	if ( (myParms.UserName != "") && (myParms.UserPass != "") ){
		req.SetBasicAuth(myParms.UserName, myParms.UserPass)
	}

	resp, err := client.Do(req)

	if err != nil {
		return res, ""
	} else {
		defer resp.Body.Close()

		//fmt.Println("response Status:", resp.Status)
		//fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println("response Body:", string(body))
		retStr = string(body)
		res = resp

		//data := decodeDevices(string(body))
		//fmt.Println(data)
		return res, retStr
	}
}

func ConnectES(myParms ParamStruct)(resp *http.Response, respStr string){
	resp, respStr = getURL(myParms)
	return resp, respStr
}

func IndexExists(myParms ParamStruct, indexParm string)(bool){
	myParms.Url = myParms.Url + "/" + indexParm
	myParms.Method = "HEAD"
	resp, _ := getURL(myParms)
	if resp.Status == "404 Not Found"{
		return false
	} else {
		return true
	}
}

func DeleteIndex(myParms ParamStruct, indexParm string)(resp *http.Response, respStr string){
	myParms.Url = myParms.Url + "/" + indexParm
	myParms.Method = "DELETE"
	resp, respStr = getURL(myParms)
	return resp, respStr
}

