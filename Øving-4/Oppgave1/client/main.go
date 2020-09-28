package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type msgBody struct {
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
}

type resBody struct {
	Output string `json:"output"`
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

			message := &msgBody{value1, value2}
			data, err := json.Marshal(message)
			if err != nil {
				log.Fatal(err)
			}

			switch method {
			case "add":
				resp, err := http.Post("http://localhost:8080/add", "application/json", bytes.NewBuffer(data))
				if err != nil {
					log.Fatal(err)
				}

				defer resp.Body.Close()

				res := &resBody{}
				json.NewDecoder(resp.Body).Decode(res)

				fmt.Println(value1, "+", value2, "=", res.Output)

			case "subtract":
				resp, err := http.Post("http://localhost:8080/subtract", "application/json", bytes.NewBuffer(data))
				if err != nil {
					log.Fatal(err)
				}

				defer resp.Body.Close()

				res := &resBody{}
				json.NewDecoder(resp.Body).Decode(res)

				fmt.Println(value1, " - ", value2, " = ", res.Output)
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
