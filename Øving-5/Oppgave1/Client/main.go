package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
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
	for {
		fmt.Println("Add or subtract?")
		var output string
		fmt.Scanln(&output)
		output = strings.ToLower(output)
		if output == "add" || output == "subtract" {
			method := output
			fmt.Println("Value 1: ")
			fmt.Scanln(&output)
			value1, err := strconv.Atoi(output)
			if err != nil {
				fmt.Println("You must give a number!")
				continue
			}
			fmt.Println("Value 2: ")
			fmt.Scanln(&output)
			value2, err := strconv.Atoi(output)
			if err != nil {
				fmt.Println("You must give a number!")
				continue
			}

			msg, err := json.Marshal(&msgBody{value1, value2, method})
			if err != nil {
				log.Fatal(err)
			}

			conn, err := net.Dial("udp", ":8080")
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			conn.Write(msg)

			buffer := make([]byte, 1024)
			conn.Read(buffer)

			res := &resBody{}

			if err = json.Unmarshal(bytes.Trim(buffer, "\x00"), res); err != nil {
				log.Fatal(err)
			}

			switch res.Operation {
			case "add":
				fmt.Println(value1, "+", value2, "=", res.Output)
			case "subtract":
				fmt.Println(value1, "-", value2, "=", res.Output)
			}
		} else {
			fmt.Println("You must choose to subtract or add")
		}

		fmt.Println("Continue? [Yes/No]")
		fmt.Scanln(&output)
		output = strings.ToLower(output)
		if output == "yes" {
		} else if output == "no" {
			break
		}
	}
}
