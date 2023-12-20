package main

import (
	"flag"
	"fmt"
)

func main() {
	//var str = "[{\"title\":\"wg_award_id\",\"image\":\"zero_winback_a_2\",\"text\":\"\",\"element1\":\"\"}]"
	//var attributeMap []map[string]string
	//_ = json.Unmarshal([]byte(str), &attributeMap)
	//fmt.Println(attributeMap)
	//fmt.Println(attributeMap[0]["image"])
	ts()
}

// go run main.go -f dev
func ts() {
	conf := flag.String("f", "", "the config override file")

	flag.Parse()
	fmt.Println(*conf)
}
