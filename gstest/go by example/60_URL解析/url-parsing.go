package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)          //postgres
	fmt.Println(u.User)            //user:pass
	fmt.Println(u.Host)            //host.com:5432
	fmt.Println(u.User.Username()) //user
	p, _ := u.User.Password()
	fmt.Println(p)

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println("============")
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
