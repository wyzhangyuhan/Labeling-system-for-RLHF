package encode_array

import (
	"encoding/json"
	"log"
)

//	func StringArrayEncoder(array [][]string) string {
//		var res_string string = ""
//		for i_out, qv := range array {
//			if len(array) > 1 && i_out != 0 {
//				res_string = res_string + "-"
//			}
//			for i_in, v := range qv {
//				encoded_v := base64.StdEncoding.EncodeToString([]byte(v))
//				if i_in == 0 {
//					res_string = res_string + encoded_v
//				} else {
//					res_string = res_string + "&" + encoded_v
//				}
//			}
//		}
//		return res_string
//	}
type DecodeJson struct {
	Q    [][]string `json:"query"`
	A    [][]string `json:"answer"`
	EQ   [][]string `json:"edited_query"`
	EA   [][]string `json:"edited_answer"`
	User string     `json:"user"`
}

func StringArrayEncoder(key string, array [][]string) string {
	jsonString := map[string][][]string{key: array}
	jsonbytes, err := json.Marshal(jsonString)
	if err != nil {
		log.Panicf("encoder wrong")
	}
	return string(jsonbytes)

}
func StringArrayDecoder(rawGroup ...[]byte) *DecodeJson {
	var dj DecodeJson
	for _, v := range rawGroup {
		json.Unmarshal(v, &dj)
	}
	return &dj
}

func UnfoldStringArray(stream []string) [][]string {
	var unfolded = make([][]string, 0)
	for _, v := range stream {
		unfolded = append(unfolded, []string{v})
	}
	return unfolded
}
