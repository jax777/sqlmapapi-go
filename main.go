// main
// Copyright (c) 2016 jax777
package main

import (
	"fmt"
	"os"
	"scantask/Sqli"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	//Host     = "127.0.0.1"
	Port     = "27017"
	User     = ""
	Password = ""
	Database = "passive_scan"
)

func dbsetup(collection string) *mgo.Collection {
	Host := os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	URL := Host + ":" + Port
	session, err := mgo.Dial(URL) //连接服务器
	if err != nil {
		panic(err)
	}
	sess := session.DB(Database).C(collection)
	return sess
}

func scan(sqltask Sqli.Sqltasks, c chan bool) {
	sqltask.Run()
	if sqltask.Vul == true {
		con := dbsetup("vul_info")
		con.Insert(sqltask)
		fmt.Print("write\n")
		<-c
	}
}
func main() {
	c := make(chan bool, 10) //limit 10 scan
	var err error
	for {
		readcon := dbsetup("http_info")
		c <- true
		var sqltask Sqli.Sqltasks
		err = readcon.Find(bson.M{"status": 0}).One(&sqltask)
		fmt.Print("read\n")
		if err == nil {
			sqltask.Status = 1
			readcon.Update(bson.M{"url": sqltask.URL}, sqltask)
			go scan(sqltask, c)
		} else {
			time.Sleep(2 * time.Minute)
		}
	}
}

/*var Readdb dbconinfo
err := Rdbsetup(&Readdb)
if err == nil {
	defer Readdb.sess.Close()
	result := Readdb.col.Find(db.Cond{"status": 0})
	var sqltask Sqli.Sqltasks
	err1 := result.One(&sqltask)
	if err1 == nil {
		sqltask.Status = 1
		result.Update(sqltask)
		go scan(sqltask, c)
	}
}
*/
