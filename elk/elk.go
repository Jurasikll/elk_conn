package elk

//import (
//	"net"
//)

type Elk struct {
	Answer_bytes []byte
	body         string
	elk_addres   string
}

func NewElk(addres string) *Elk {
	return &Elk{elk_addres: addres}
}

func (e Elk) new_rec(path string, body string) {

}
