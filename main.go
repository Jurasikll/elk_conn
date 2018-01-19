// elk_conn project main.go
package main

import (
	"elk_conn/elk"
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
	Elk_addres string
})

func main() {

	toml.DecodeFile(CONFIG_PATH, &set)
	elk_get_test()
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
