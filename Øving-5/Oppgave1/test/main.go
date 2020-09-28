package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
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
	conn, err := net.Dial("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	msg, err := json.Marshal(&msgBody{2, 3, "add"})
	if err != nil {
		log.Fatal(err)
	}

	conn.Write(msg)

	buffer := make([]byte, 1024)
	conn.Read(buffer)

	res := &resBody{}

	if err = json.Unmarshal(bytes.Trim(buffer, "\x00"), res); err != nil {
		log.Fatal(err)
	}

	switch res.Operation {
	case "add":
		fmt.Println(2, "+", 3, "=", res.Output)
	case "subtract":
		fmt.Println(2, "-", 3, "=", res.Output)
	}
}
