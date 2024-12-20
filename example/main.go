package main

import (
	"log"

	ads "github.com/0x7A77/go-ADS"
)

func main() {
	conn, err := ads.NewConnection("10.2.30.26", "192.168.202.200.1.1", "192.168.202.213.1.1", 48898,801, 888)
	conn.Connect()
	if err != nil {
		log.Println("ads connection fail!")
	}
	log.Println("ads conn successed!")
	defer conn.Close()
	device, err := conn.ReadDeviceInfo()
	if err != nil {
		log.Println("deviceinfo cannot print!")
	}
	log.Println(device)

	// 温度：6
	data, err := conn.Read(0x4020, 0x870, 6)
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("温度:", string(data.Data))
}
