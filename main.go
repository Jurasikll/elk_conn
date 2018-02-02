// elk_conn project main.go
package main

import (
	"elk_conn/elk"
	"elk_conn/rest_db_conn"
	"encoding/json"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type test_answer_get struct {
	Hits struct {
		Total int
	}
}

const (
	CONFIG_PATH string = `C:\goplace\src\elk_conn\conf.ini`
	TEST_ROAD   string = "test_parent/test_child/"
)

var set = new(struct {
	Elk_addres    string
	Db_host       string
	Db_port       int
	Db_user       string
	Db_pwd        string
	Db_name       string
	T_user        string
	T_pwd         string
	T_local_ip    string
	T_local_port  int
	T_server_ip   string
	T_server_port int
	T_remote_ip   string
	T_remote_port int
	Test_sql      string
})

func main() {
	toml.DecodeFile(CONFIG_PATH, &set)
	db_test()
}

func db_test() {
	var id int
	var crd time.Time
	var mod time.Time
	var list int
	var name string
	fmt.Println(set.Elk_addres)
	conn := rest_db_conn.Init(set.T_user, set.T_pwd, &rest_db_conn.Endpoint{Host: set.T_local_ip, Port: set.T_local_port}, &rest_db_conn.Endpoint{Host: set.T_server_ip, Port: set.T_server_port}, &rest_db_conn.Endpoint{Host: set.T_remote_ip, Port: set.T_remote_port}, set.Db_host, set.Db_port, set.Db_user, set.Db_pwd, set.Db_name)

	rows, err := conn.Query(set.Test_sql)
	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		rows.Scan(&id, &crd, &mod, &name, &list)

		fmt.Println(id, "-", crd, "-", mod, "-", name, "-", list)
	}
}

func elk_get_test() {
	elk_conn := elk.NewElk(set.Elk_addres)
	elk_conn.Get_data(TEST_ROAD, "_id:AWDqn0Lv62WqGwI2oVMS")
	ta := test_answer_get{}
	json.Unmarshal(elk_conn.Answer_bytes, &ta)
	fmt.Println(string(elk_conn.Answer_bytes))
	fmt.Printf("%d\r\n", ta.Hits.Total)
}

func elk_put_test() {
	test_data := struct {
		User    string    `json:"user"`
		Init_id int       `json:"init_id"`
		Is_set  bool      `json:"is_set"`
		Time    time.Time `json:"time"`
	}{User: "Test_user", Init_id: 125, Is_set: false, Time: time.Now()}
	fmt.Printf("%v123\r\n", test_data)

	elk_conn := elk.NewElk(set.Elk_addres)
	body, _ := json.Marshal(test_data)
	elk_conn.Put_data(body, TEST_ROAD, "AWEN2U5i62WqGwI2QS4d")
	fmt.Printf("%v\r\n", string(elk_conn.Answer_bytes))
}
