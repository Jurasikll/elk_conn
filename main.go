// elk_conn project main.go
package main

import (
	"elk_conn/elk"
	"fmt"

	"github.com/BurntSushi/toml"
)

const (
	CONFIG_PATH string = `C:\goplace\src\elk_conn\conf.ini`
)

func main() {
	set := new(struct {
		Elk_addres string
	})
	toml.DecodeFile(CONFIG_PATH, &set)
	elk_conn := elk.NewElk(set.Elk_addres)
	fmt.Printf("%v\r\n", elk_conn)
}
