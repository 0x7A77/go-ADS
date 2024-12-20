# ADS

基于golang实现的倍福ads协议组件库，已实现TCP连接。本库fork[stamp/goADS: Beckhoff ADS client package for the Go Programming Language (golang)](https://github.com/stamp/goADS)开源库，将原自动生成的sourceNetId改成手动sourceNetId，以实现通过模块连接倍福plc的情况。

## 例



```golang
package main

import (
	"log"
	ads "github.com/0x7A77/go-ADS"
)

func main() {
    // ip，targetNetid，sourceNetid，port（默认48898）,targetPort（801、811、821）,sourcePort
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

	// 地址，偏移量，长度
	data, err := conn.Read(0x4020, 0x870, 6)
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("data:", string(data.Data))
}

```


