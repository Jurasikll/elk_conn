package rest_db_conn

import (
	"time"
)

type Rest_elk_card struct {
	Name                  string    `json:"name"`
	Card_id               int       `json:"card_id"`
	Assign_date           time.Time `json:"assign_date"`
	Create_date           time.Time `json:"create_date"`
	In_progress_date      time.Time `json:"in_progress_date"`
	To_vendor_date        time.Time `json:"to_vendor_date"`
	Close_date            time.Time `json:"close_date"`
	Create_to_assign      int       `json:"create_to_assign"`
	Assign_to_in_progress int       `json:"assign_to_in_progress"`
	Create_to_in_progress int       `json:"create_to_in_progress"`
	In_progress_to_close  int       `json:"in_progress_to_close"`
	Create_to_close       int       `json:"create_to_close"`
	Is_done               bool      `json:"is_done"`
	Is_cancelled          bool      `json:"is_cancelled"`
	Is_copy               bool      `json:"is_copy"`
	Assign_list           []string  `json:"assign_list"`
	Mark_list             []string  `json:"mark_list"`
}

//func Create_card(name string,
//	card_id int,
//	assign_date time.Time,
//	create_date time.Time,
//	in_progress_date time.Time,
//	to_vendor_date time.Time,
//	close_date time.Time,
//	is_done bool,
//	is_cancelled bool,
//	is_copy bool) *Rest_elk_card {
//	rec := &Rest_elk_card{
//		Name:                  name,
//		Card_id:               card_id,
//		Assign_date:           assign_date,
//		Create_date:           create_date,
//		In_progress_date:      in_progress_date,
//		To_vendor_date:        to_vendor_date,
//		Close_date:            close_date,
//		Create_to_assign:      int(assign_date.Sub(create_date).Hours() / 24),
//		Assign_to_in_progress: int(in_progress_date.Sub(assign_date).Hours() / 24),
//		Create_to_in_progress: int(in_progress_date.Sub(create_date).Hours() / 24),
//		In_progress_to_close:  int(in_progress_date.Sub(close_date).Hours() / 24),
//		Create_to_close:       int(close_date.Sub(create_date).Hours() / 24),
//		Is_done:               is_done,
//		Is_copy:               is_copy,
//		Assign_list:           assign_list,
//		Mark_list:             mark_list,
//	}
//	rec.Assign_date = assign_date
//	rec.In_progress_date = in_progress_date
//	rec.To_vendor_date = to_vendor_date
//	rec.Close_date = close_date
//	rec.Is_done = is_done
//	return rec
//}

func Create_card() *Rest_elk_card {
	rec := &Rest_elk_card{}
	return rec
}

type Temp_rest_card struct {
	Card_id           int       `json:"card_id"`
	Create_date       time.Time `json:"create_date"`
	Close_date        time.Time `json:"close_date"`
	Name              string    `json:"name"`
	Close_list        int       `json:"close_list"`
	Interval_to_close int       `json:"interval_to_close"`
}
