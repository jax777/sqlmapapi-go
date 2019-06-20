package Sqli

import (
	"testing"
)

func Test_Task(t *testing.T) {
	var Sqltask Sqltasks
	if Sqltask.Task_new() == true {
		if Sqltask.Task_delete() {
		} else {
			t.Error("task_delete wrong \n")
		}
	} else {
		t.Error("task_new wrong\n")
	}
}

func Test_Scan_start(t *testing.T) {
	var Sqltask Sqltasks
	Sqltask.URL = "http://127.0.0.1"
	if Sqltask.Task_new() == true {
		if Sqltask.Scan_start() == false {
			t.Error("Scan_start wrong\n")
		} else {
			Sqltask.Task_delete()
		}
	}
}

func Test_Scan_status(t *testing.T) {
	var Sqltask Sqltasks
	Sqltask.URL = "http://127.0.0.1"
	Sqltask.Task_new()
	Sqltask.Scan_start()
	if Sqltask.Scan_status() != "running" {
		t.Error("Scan_status wrong\n")
	}
	Sqltask.Task_delete()
}

func Test_sql(t *testing.T) {
	var Sqltask Sqltasks
	Sqltask.URL = "http://127.0.0.1/sql.php"
	Sqltask.Method = "post"
	Sqltask.Body = "sql=111"
	Sqltask.Run()
	if Sqltask.Vul == false {
		t.Error("something wrong\n")
	}
}
