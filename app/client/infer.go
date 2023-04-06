package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"label_system/app/curd/infer_curd"
	"label_system/app/models/req_model"
	"label_system/config"
	"label_system/utils/build_modelstring"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type InferGateway struct {
	ModelVs    string
	HttpClient *http.Client
}
type InferReq struct {
	ModelId string   `json:"model_id" binding:"required"`
	Query   []string `json:"query" binding:"required"`
	Answer  []string `json:"answer" binding:"required"`
}

func (infer *InferGateway) ChatModel(req *InferReq) (*req_model.SubmitQuestionRsp, error) {
	// url := config.Conf.AckdgeDomain + "/v1/model/infer/chatgpt"
	_, modelItem := infer_curd.CreateModelCurdFactory().GetModelById(req.ModelId)
	url := config.Conf.AckdgeDomain + "/v1/model/infer/" + modelItem.ModelName

	if modelItem.ModelName == "multi_qa_gen_server_v0" {
		modelstring := build_modelstring.BuildString(req.Query, req.Answer)
		inferReq := req_model.InferReq{

			Inputs: []map[string]interface{}{
				{
					"name":  "input_text",
					"type":  "string",
					"shape": []int{1, 1},
					"data":  []string{modelstring},
				},
			},
		}
		log.Printf("ChatGPT req:%+v\n", inferReq)
		jsonData, _ := json.Marshal(inferReq)
		request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		request.Header.Add("Content-type", "application/json;charset=utf-8")
		request.Header.Add("x-infer-request-token", "lQ/YkBb3y36HofA2jRyxJHqRhEqIjOMLE4luiUiU7/ehkOQxaBJ17NL2NYXh6+2BbBZ284I2BkjJ4I/uy/Pplg==")
		// rst, err := infer.HttpClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
		rst, err := infer.HttpClient.Do(request)
		if err != nil {
			fmt.Printf("infer err %v\n", err)
			return nil, err
		}
		var inferRsp req_model.InferRsp
		json.NewDecoder(rst.Body).Decode(&inferRsp)
		log.Printf("ChatGPT rsp:%+v\n", inferRsp)
		if inferRsp.Code != 0 {
			err = fmt.Errorf("code:%v message:%v", inferRsp.Code, inferRsp.Message)
			log.Printf("%v", err)
			return nil, err
		}
		rsp := req_model.InferRsp2SubmitQuestionRsp(&inferRsp)
		return rsp, nil
	} else {
		modelstring := build_modelstring.BuildStringNeox(req.Query, req.Answer)
		inferReq := req_model.InferReq{

			Inputs: []map[string]interface{}{
				{
					"name":  "INPUT_0",
					"type":  "string",
					"shape": []int{3, 1},
					"data":  []string{modelstring, modelstring, modelstring},
				},
				{
					"name":  "INPUT_1",
					"type":  "uint32",
					"shape": []int{3, 1},
					"data":  []int{1024, 1024, 1024},
				},
				{
					"name":  "temperature",
					"type":  "float",
					"shape": []int{3, 1},
					"data":  []float32{1.0, 1.0, 1.0},
				},
				{
					"name":  "runtime_top_p",
					"type":  "float",
					"shape": []int{3, 1},
					"data":  []float32{0.7, 0.7, 0.7},
				},
				{
					"name":  "runtime_top_k",
					"type":  "uint32",
					"shape": []int{3, 1},
					"data":  []int{0, 0, 0},
				},
				{
					"name":  "beam_width",
					"type":  "uint32",
					"shape": []int{3, 1},
					"data":  []int{1, 1, 1},
				},
				{
					"name":  "random_seed",
					"type":  "uint64",
					"shape": []int{3, 1},
					"data":  []int64{GenRandSeed(), GenRandSeed(), GenRandSeed()},
				},
			},
		}
		log.Printf("ChatGPT req:%+v\n", inferReq)
		jsonData, _ := json.Marshal(inferReq)
		request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		request.Header.Add("Content-type", "application/json;charset=utf-8")
		request.Header.Add("x-infer-request-token", "lQ/YkBb3y36HofA2jRyxJHqRhEqIjOMLE4luiUiU7/ehkOQxaBJ17NL2NYXh6+2BbBZ284I2BkjJ4I/uy/Pplg==")

		rst, err := infer.HttpClient.Do(request)
		if err != nil {
			fmt.Printf("infer err %v\n", err)
			return nil, err
		}
		var inferRsp req_model.InferRsp
		json.NewDecoder(rst.Body).Decode(&inferRsp)
		log.Printf("ChatGPT rsp:%+v\n", inferRsp)
		if inferRsp.Code != 0 {
			err = fmt.Errorf("code:%v message:%v", inferRsp.Code, inferRsp.Message)
			log.Printf("%v", err)
			return nil, err
		}
		rsp := req_model.InferRsp2SubmitQuestionRspNeoX(&inferRsp)
		// fmt.Printf("%v", rsp)
		return rsp, nil
	}
}

func GenRandSeed() int64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	num := rand.Int63n(time.Now().Unix())
	return num
}
