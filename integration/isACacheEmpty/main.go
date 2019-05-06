package main

import (
	"fmt"
	"time"

	"github.com/michaeldistler/gomemcache/memcache"
)

func main() {
	client := memcache.New("def", "10.0.7.187:11211", "10.0.6.150:11211", "10.0.7.29:11211")
	// tcp := &net.TCPAddr{
	// 	IP:   net.ParseIP("10.0.7.125"),
	// 	Port: 11211,
	// 	Zone: "",
	// }

	i := 0
	for i < 10000 {
		boolthing, err := client.IsACacheEmpty()
		fmt.Println(boolthing)
		if err != nil {
			fmt.Println(err)
		}
		i++
		time.Sleep(1 * time.Second)
	}

}
