// elk_conn project main.go
package main

import (
	"elk_conn/elk"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

const (
	CONFIG_PATH string = `C:\goplace\src\elk_conn\conf.ini`
)

func main() {
	set := new(struct {
		Elk_addres string
	})
	test_data := struct {
		User    string    `json:"user"`
		Init_id int       `json:"init_id"`
		Is_set  bool      `json:"is_set"`
		Time    time.Time `json:"time"`
	}{User: "Test_user", Init_id: 125, Is_set: false, Time: time.Now()}
	fmt.Printf("%v123\r\n", test_data)
	test_road := "test_parent/test_child/"
	toml.DecodeFile(CONFIG_PATH, &set)
	elk_conn := elk.NewElk(set.Elk_addres)
	elk_conn.Put_data(test_data, test_road, -1)
	fmt.Printf("%v\r\n", string(elk_conn.Answer_bytes))
}
