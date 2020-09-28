package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
	"strconv"
)

type msgBody struct {
	Value1    int    `json:"value1"`
	Value2    int    `json:"value2"`
	Operation string `json:"operation"`
}

type resBody struct {
	Output    string `json:"output"`
	Operation string `json:"operation"`
}

func main() {

	pc, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buffer := make([]byte, 1024)
		_, addr, err := pc.ReadFrom(buffer)
		if err != nil {
			log.Fatal(err)
		}
		msg := &msgBody{}
		if err = json.Unmarshal(bytes.Trim(buffer, "\x00"), msg); err != nil {
			log.Fatal(err)
		}

		switch msg.Operation {
		case "add":
			res := &resBody{strconv.Itoa(msg.Value1 + msg.Value2), "add"}
			resp, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			pc.WriteTo(resp, addr)
		case "subtract":
			res := &resBody{strconv.Itoa(msg.Value1 - msg.Value2), "subtract"}
			resp, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			pc.WriteTo(resp, addr)
		}
	}
}
