package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"os"
	"runtime"
	"strings"
	"time"
)
const (
	ACCESS_URL = "/family/qos/startQos.action"
	UP_QOS_URL = "http://api.cloud.189.cn/family/qos/startQos.action"
)
/*
readALL
 */
 func readAll(reader io.Reader)string{
	 p,err:=ioutil.ReadAll(reader)
	 if err!=nil{
	 	return ""
	 }
	 return string(p[:])
 }
/*
封装的GET方法
 */
//Post http get method
func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("new request is fail ")
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	return client.Do(req)
}
/*
封装的POST方法
 */
//Post http post method
func Post(url string, body map[string]string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//add post body
	var req *http.Request
	encoder:= url2.Values{}
	if body != nil {
		for k,v := range body{
			encoder.Add(k,v)
		}
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(encoder.Encode()))
	req.URL.Query()
	if err != nil {
		return nil, errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	return client.Do(req)
}
/*
hmac 字符串生成
 */
func mac_sha1(msg, key string) string{
	h := hmac.New(sha1.New,[]byte(key[:]))
	h.Write([]byte(msg[:]))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
/*
计算signature
 */
func get_signature(access_url, session_key, session_secret, request_method, date string) string{
	params := fmt.Sprintf("SessionKey=%s&Operate=%s&RequestURI=%s&Date=%s", session_key,
		request_method,
		access_url,
		date)
	println(params)
	return mac_sha1(params, session_secret)
}
/*
格式化日期字符串
:return: str 'Sun, 31 Mar 2019 05:35:33 GMT'
 */
func create_date()string{
	//var time_stop int64 = 16000
	//var time_start int64 = 12500
	// 默认误差： time_stop-time_start = 3500
	//此处实现不够优雅，需要UTC 时间 程序大概率是可能在中国使用，故减去8小时
	tsInt64 := (time.Now().UnixNano() / 1e6+3500)/1000-8*3600
	return time.Unix(tsInt64,0).Format("Mon, 02 Jan 2006 15:04:05 GMT")
}
func heartService(session_key, session_secret, method string, extra_header, send_data map[string] string) (status_code int,response string){
	date := create_date()
	signature := get_signature(ACCESS_URL, session_key, session_secret, method, date)
	fmt.Printf("heart_beat:<signature:%s>\n",signature)
	fmt.Printf("date:<%s>\n",date)

	header :=map[string] string{
		"SessionKey": session_key,
		"Signature":signature,
		"Date":date,
	}
	/*
		write extra_header
	 */
	for k, v := range extra_header {
		header[k]=v
	}
	var res *http.Response
	var err error
	if method=="GET" {
		res, err = Get(UP_QOS_URL,send_data,header)
	}else{
		res, err = Post(UP_QOS_URL,send_data,nil,header)
	}
	if err!=nil{
		return 0,""
	}
	return res.StatusCode,readAll(res.Body)
}

func main(){
	data :=[]byte{}
	if len(os.Args) !=2{
		tempData,_ :=ioutil.ReadFile("./config.json")
		data = tempData[:]
	}else{
		tempData,_ :=ioutil.ReadFile("./config.json")
		data = tempData[:]
	}
	config, _ := simplejson.NewJson(data)
	session_key,_ := config.Get("session_key").String()
	session_secret,_ := config.Get("session_secret").String()
	setting,_ := config.Get("setting").Map()
	method := setting["method"].(string)
	rate,_ := setting["rate"].(json.Number).Int64()
	extra_header,_ := config.Get("extra_header").Map()
	send_data,_ := config.Get("send_data").Map()
	count :=1

	header := make(map[string]string)
	for k,v :=range extra_header{
		header[k] = v.(string)
	}
	params := make(map[string]string)
	for k,v :=range send_data{
		params[k] = v.(string)
	}
	for ;true;{
		runtime.Gosched()
		fmt.Printf("Sending heart_beat package <%d>\n",count)
		code,res:=heartService(session_key, session_secret, method, header, params)
		fmt.Println("status_code:",code)
		fmt.Println("response:",res)
		fmt.Printf("Send heart_beat <%d> package Success\n",count)
		fmt.Println("*******************************************")
		time.Sleep(time.Duration(rate)*time.Second)
		count+=1
	}
}