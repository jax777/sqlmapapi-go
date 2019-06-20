// main_test
package main

import (
	"testing"
)

type Response struct {
	Taskid     string `json:"taskid" db:",omitempty,json"`
	URL        string `json:"url" db:",json"`
	Vul        bool   `json:"vul" db:",json"`
	Method     string `json:"method" db:",json"`
	Cookie     string `json:"cookie" db:",json"`
	Body       string `json:"data" db:",json"`
	User_agent string `json:"user-agent" db:",json"`
	Status     int    `json:"status" db:",json"`
}

func Test_init(t *testing.T) {
	var test Response
	test.URL = "http://127.0.0.1/sql.php"
	test.Body = "sql=111"
	test.Method = "POST"
	readcon := dbsetup("http_info")
	readcon.Insert(test)

}
