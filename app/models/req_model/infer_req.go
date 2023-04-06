package req_model

import (
	"fmt"
	"strings"
)

type InferReq struct {
	RequestId string                   `json:"request_id"`
	Inputs    []map[string]interface{} `json:"inputs"`
}

type InferRsp struct {
	Code      int                      `json:"code"`
	Message   string                   `json:"message"`
	RequestId string                   `json:"request_id"`
	Outputs   []map[string]interface{} `json:"outputs"`
}

func InferRsp2SubmitQuestionRsp(t *InferRsp) *SubmitQuestionRsp {
	d := t.Outputs[0]["data"].([]interface{})
	var a []string
	for _, v := range d {
		s := fmt.Sprint(v)
		if strings.TrimSpace(s) == "" {
			fmt.Println("generate empty")
		} else {
			a = append(a, s)
		}
	}
	return &SubmitQuestionRsp{
		Answers: a,
	}
}

func InferRsp2SubmitQuestionRspNeoX(t *InferRsp) *SubmitQuestionRsp {
	d := t.Outputs[1]["data"].([]interface{})
	var a []string
	for _, v := range d {
		s := fmt.Sprint(v)
		if strings.TrimSpace(s) == "" {
			fmt.Println("generate empty")
		} else {
			s = strings.Replace(s, "<|endoftext|>", "", -1)
			fmt.Printf("All Infer%v", s)
			s = GetBotLatestRes(s)
			// fmt.Printf("%v", s)
			a = append(a, s)
		}
	}
	return &SubmitQuestionRsp{
		Answers: a,
	}
}

func GetBotLatestRes(modelstring string) string {
	resList := strings.Split(modelstring, "<bot>：")
	return resList[len(resList)-1]
}
