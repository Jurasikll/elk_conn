package elk

import (
	"bytes"

	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	METHOD_POST               string = "POST"
	METHOD_PUT                string = "PUT"
	URL_DELIMETR              string = "/"
	CREATE_RECORD_WITH_ID_PTR string = "%s/%s%s"
	GET_DATA_URL_PTRN         string = "%s/%s_search?q=%s"
	ADD_NEW_REC_MARK          string = "-1"
)

type Elk struct {
	Answer_bytes []byte
	body         string
	elk_addres   string
	http_client  http.Client
}

func NewElk(addres string) *Elk {
	return &Elk{elk_addres: addres, http_client: http.Client{}}
}

func (e *Elk) Get_main_info() {
	resp, _ := e.http_client.Get(e.elk_addres)
	e.update_resp(resp.Body)
}

//body - []byte convert to httpBody; road - string, path to put data; record_id - id of record, send elk.ADD_NEW_REC_MARK if need add next record
func (e *Elk) Put_data(body []byte, road string, record_id string) {
	var method string
	var url string
	if record_id != ADD_NEW_REC_MARK {
		url = fmt.Sprintf(CREATE_RECORD_WITH_ID_PTR, e.elk_addres, road, record_id)
		method = METHOD_PUT
	} else {
		url = e.elk_addres + `/` + road
		method = METHOD_POST
	}
	fmt.Println(url)
	req, _ := http.NewRequest(method, url, bytes.NewReader(body))
	resp, _ := e.http_client.Do(req)
	e.update_resp(resp.Body)
}

func (e *Elk) update_resp(body io.Reader) {

	e.Answer_bytes, _ = ioutil.ReadAll(body)
}

func (e *Elk) Delete(path string) {

	req, _ := http.NewRequest("DELETE", path, nil)
	e.http_client.Do(req)
}

func (e *Elk) Get_data(road string, query string) {
	if query != "" {
		url := fmt.Sprintf(GET_DATA_URL_PTRN, e.elk_addres, road, query)
		resp, _ := e.http_client.Get(url)
		e.update_resp(resp.Body)
	}

}
