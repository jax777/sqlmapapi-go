// Copyright (c) 2016 jax777
package Sqli

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/antonholmquist/jason"
	"github.com/mozillazg/request"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	byt, _ := json.Marshal(obj)
	var dat map[string]interface{}
	json.Unmarshal(byt, &dat)
	return dat
}

func setrequest() *request.Request {
	c := &http.Client{}
	req := request.NewRequest(c)
	//req.Proxy = "http://127.0.0.1:8080"
	req.Client.Timeout = time.Duration(1 * time.Second)
	return req
}

type Header struct {
	http.Header
}

/*
type AutoSql interface {
	Task_new() bool
	Task_delete() bool
	Scan_start() bool
	Scan_status() string
	Scan_data()
	Option_set()
	Scan_stop()
	Scan_kill()
	Run()
}
*/
const Server = "http://127.0.0.1:8775/"

type Sqltasks struct {
	Taskid     string `json:"taskid" db:",omitempty,json"`
	URL        string `json:"url" db:",json"`
	Vul        bool   `json:"vul" db:",json"`
	Method     string `json:"method" db:",json"`
	Cookie     string `json:"cookie" db:",json"`
	Body       string `json:"data" db:",json"`
	User_agent string `json:"user-agent" db:",json"`
	Status     int    `json:"status" db:",json"`
	config     string `json:"tech" db:",json"`
}

/*
type Sqltest struct {
	Taskid        string `json:"taskid" db:",omitempty,json"`
	Method        string `json:"method" db:",json"`
	ContentType   string `json:"content_type" db:",json"`
	ContentLength uint   `json:"content_length" db:",json"`
	Host          string `json:"host" db:",json"`
	URL           string `json:"url" db:",json"`
	Header        Header `json:"header,omitempty" db:",json"`
	Body          []byte `json:"body,omitempty" db:",json"`
	RequestHeader Header `json:"request_header,omitempty" db:",json"`
	RequestBody   []byte `json:"request_body,omitempty" db:",json"`
	Vul           bool   `json:"vul" db:",json"`
	Status        int    `json:"status" db:",json"`
}
*/
func (Sqltask *Sqltasks) Task_new() bool {
	Sqltask.config = "BT"
	result := false
	req := setrequest()
	resp, err := req.Get(Server + "task/new")
	if err == nil {
		defer resp.Body.Close() // **Don't forget close the response body**
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 == nil {
			j, err2 := jason.NewObjectFromBytes(body)
			if err2 == nil {
				var err3 error
				Sqltask.Taskid, err3 = j.GetString("taskid")
				if err3 == nil && len(Sqltask.Taskid) > 0 {
					result = true
				}
			}
		}
	}
	return result
}

func (Sqltask *Sqltasks) Task_delete() bool {
	result := false
	req := setrequest()
	resp, err := req.Get(Server + "task/" + Sqltask.Taskid + "/delete")
	if err == nil {
		defer resp.Body.Close() // **Don't forget close the response body**
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 == nil {
			j, err2 := jason.NewObjectFromBytes(body)
			if err2 == nil {
				result, _ = j.GetBoolean("success")
			}
		}
	}
	return result
}

func (Sqltask *Sqltasks) Scan_start() bool {
	result := false
	req := setrequest()
	req.Headers = map[string]string{
		"Content-Type": "application/json",
	}
	req.Json = Struct2Map(Sqltask)
	resp, err := req.Post(Server + "scan/" + Sqltask.Taskid + "/start")
	if err == nil {
		defer resp.Body.Close() // **Don't forget close the response body**
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 == nil {
			j, err2 := jason.NewObjectFromBytes(body)
			if err2 == nil {
				success, err3 := j.GetBoolean("success")
				if err3 == nil {
					if success == true {
						result = true
					}
				}
			}
		}
	}
	return result
}

func (Sqltask *Sqltasks) Scan_status() string {
	result := ""
	req := setrequest()
	resp, err := req.Get(Server + "scan/" + Sqltask.Taskid + "/status")
	if err == nil {
		defer resp.Body.Close() // **Don't forget close the response body**
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 == nil {
			j, err2 := jason.NewObjectFromBytes(body)
			if err2 == nil {
				var err3 error
				Status, err3 := j.GetString("status")
				if err3 == nil {
					result = Status
				}
			}
		}
	}
	return result
}

func (Sqltask *Sqltasks) Scan_data() {
	req := setrequest()
	//req.Proxy = "http://127.0.0.1:8080"
	resp, err := req.Get(Server + "scan/" + Sqltask.Taskid + "/data")
	if err == nil {
		defer resp.Body.Close() // **Don't forget close the response body**
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 == nil {
			j, err2 := jason.NewObjectFromBytes(body)
			if err2 == nil {
				data, _ := j.GetInterface("data")
				dat, _ := json.Marshal(data)
				if len(dat) > 5 {
					Sqltask.Vul = true
				}
			}
		}
	}
}

/*
func (Sqltask *Sqltasks) Option_set() {
	req := setrequest()
	req.Headers = map[string]string{
		"Content-Type": "application/json",
	}
	options := make(map[string]map[string]interface{})
	option := make(map[string]interface{})
	option["smart"] = true
	options["option"] = option
	req.Json = options
	req.Post(Server + "option/" + Sqltask.Taskid + "/set")
}
*/

func (Sqltask *Sqltasks) Scan_stop() {
	req := setrequest()
	req.Get(Server + "task/" + Sqltask.Taskid + "/stop")
}
func (Sqltask *Sqltasks) Scan_kill() {
	req := setrequest()
	req.Get(Server + "task/" + Sqltask.Taskid + "/kill")
}

func (Sqltask *Sqltasks) Run() {
	Sqltask.config = "BT"
	Sqltask.Vul = false
	if Sqltask.Task_new() {
		//Sqltask.Option_set()
		Sqltask.Scan_start()
		time.Sleep(time.Second)
		start_time := time.Now()
		flag := true
		for flag {
			status := Sqltask.Scan_status()
			switch status {
			case "terminated":
				flag = false
				break
			case "running":
				time.Sleep(10 * time.Second)
				break
			default:
				flag = false
				break
			}
			if time.Now().Sub(start_time) > 3000*time.Second {
				Sqltask.Scan_stop()
				time.Sleep(time.Second)
				Sqltask.Scan_kill()
				flag = false
				break
			}
		}
		Sqltask.Scan_data()
		time.Sleep(10 * time.Second)
		Sqltask.Task_delete()
	}
}
