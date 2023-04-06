package oss_curd

import (
	"encoding/json"
	"fmt"
	"log"
)

type QuestionRaw struct {
	Question []string `json:"Q"`
	Answer   []string `json:"A"`
	Tips     string   `json:"TIPS"`
}

func Stream2Json(stream []string) []QuestionRaw {
	var qrlist []QuestionRaw
	for _, v := range stream {
		var qr QuestionRaw
		if err := json.Unmarshal([]byte(v), &qr); err != nil {
			log.Printf("err: %v", err)
			fmt.Printf("Unmarshall string : %v", v)
			return qrlist
		}
		qrlist = append(qrlist, qr)
	}

	return qrlist
}
