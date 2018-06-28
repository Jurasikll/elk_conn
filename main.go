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
	//	TEST_ROAD   string = "test_parent/test_child/"
	TEST_ROAD string = "cards/light/"
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
	//	elk_conn := elk.NewElk(set.Elk_addres)
	//	elk_conn.Delete("http://t2ru-elk-01:9200/cards")
	db_test()
	//	var nil_time time.Time = time.Time{}
	//	rec := rest_db_conn.Rest_elk_card{}
	//	//	rec.Assign_date = time.Now()
	//	test_rec := rest_db_conn.Rest_elk_card{}
	//	body, _ := json.Marshal(rec)
	//	if rec.Assign_date == test_rec.Assign_date {
	//		fmt.Printf("%v\r\n", string(body))
	//	} else {
	//		fmt.Printf("%v\r\n", rec.Assign_date)
	//	}
	//	rec.Assign_date = time.Now()
	//	fmt.Printf("%v\r\n", rec.Assign_date)
	//	rec.Assign_date = nil_time
	//	fmt.Printf("%v\r\n", rec.Assign_date)

}

func db_test() {
	var id int
	//	var crd time.Time

	var mod time.Time
	var event_id int
	var name string
	var events_type string
	var tmp_id int = 0
	var cards map[int]rest_db_conn.Temp_rest_card
	//	var nil_time time.Time
	//	var tmp_create time.Time
	//	var tmp_Assign time.Time
	//	var tmp_in_progress time.Time
	//	var tmp_to_vendor_date time.Time
	//	var tmp_close_date time.Time
	//	var tmp_is_done bool
	//	var tmp_is_copy bool = false
	var tmp_name string
	var tmp_crd time.Time
	var tmp_closed time.Time
	var tmp_closed_list int
	fmt.Println(set.Elk_addres)
	conn := rest_db_conn.Init(set.T_user, set.T_pwd, &rest_db_conn.Endpoint{Host: set.T_local_ip, Port: set.T_local_port}, &rest_db_conn.Endpoint{Host: set.T_server_ip, Port: set.T_server_port}, &rest_db_conn.Endpoint{Host: set.T_remote_ip, Port: set.T_remote_port}, set.Db_host, set.Db_port, set.Db_user, set.Db_pwd, set.Db_name)
	defer conn.Close_rc()
	rows, err := conn.Query(set.Test_sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	cards = map[int]rest_db_conn.Temp_rest_card{}
	for rows.Next() {
		rows.Scan(&id, &mod, &name, &event_id, &events_type)
		if tmp_id != id {
			//start
			//					cards[tmp_id] = rest_db_conn.Create_card()
			//					tmp_id = id
			//					tmp_create = mod
			//					tmp_Assign = nil_time
			//					tmp_in_progress = nil_time
			//					tmp_to_vendor_date = nil_time
			//					tmp_close_date = nil_time
			//					tmp_is_done = nil_time

			if tmp_id != 0 && tmp_closed_list != 0 {
				//				cards[tmp_id].Interval_to_close = int(cards[tmp_id].Close_date.Sub(cards[id].Create_date).Hours() / 24)
				cards[tmp_id] = rest_db_conn.Temp_rest_card{Card_id: tmp_id, Name: tmp_name, Create_date: tmp_crd, Close_date: tmp_closed, Close_list: tmp_closed_list, Interval_to_close: int(tmp_closed.Sub(tmp_crd).Hours() / 24)}
			}
			tmp_closed_list = 0
			tmp_name = name
			tmp_id = id
		}

		switch events_type {
		case "copy_card":
			//			tmp_is_copy = true
		case "add_card_user", "add_card":
			//			if tmp_Assign == nil_time {
			//				tmp_Assign = mod
			//			}
			tmp_crd = mod
		case "move_card":
			switch event_id {
			case 419:
				//						if tmp_in_progress == nil_time {
				//							tmp_in_progress = mod
				//						}
			case 420:
				//						if tmp_to_vendor_date == nil_time {
				//							tmp_to_vendor_date = mod
				//						}
			case 421:
				//						if tmp_close_date == nil_time {
				//							tmp_close_date = mod
				//							tmp_is_done = true
				//						}
				tmp_closed = mod
				tmp_closed_list = 421
			case 422:
				//						if tmp_close_date == nil_time {
				//							tmp_close_date = mod
				//							tmp_is_done = false
				//						}
				tmp_closed = mod
				tmp_closed_list = 422
			}
		}

	}
	elk_conn := elk.NewElk(set.Elk_addres)
	var body []byte

	for _, card := range cards {
		//		fmt.Printf("%d - %v\r\n", id, card)

		body, _ = json.Marshal(card)
		//		elk_conn.Put_data(body, "cards/cardsState/", elk.ADD_NEW_REC_MARK)
		elk_conn.Put_data(body, TEST_ROAD, elk.ADD_NEW_REC_MARK)
		//		fmt.Printf("%v\r\n", card.Interval_to_close)

	}
	rows.Close()
	fmt.Println("--------------------------")

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
