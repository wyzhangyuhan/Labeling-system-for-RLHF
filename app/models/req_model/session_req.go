package req_model

import "label_system/app/oss_curd"

type SessionReq struct {
	UserId      string `json:"user_id" binding:"required"`
	UserName    string `json:"user_name"`
	ModelId     string `json:"model_id" binding:"required"`
	DatesetId   string `json:"dataset_id" binding:"required"`
	DataSetName string `json:"dataset_name"`
}
type SessionRsp struct {
	SessionReq
	SessionId string `json:"session_id"`
	oss_curd.QuestionRaw
}

type RetrieveQueReq struct {
	DataSetId   string `json:"dataset_id"`
	DataSetName string `json:"dataset_name"`
	UserId      string `json:"user_id"`
}

type SubmitLabelReq struct {
	SessionId    string     `json:"session_id" binding:"required"`
	Query        [][]string `json:"query" `
	Answer       [][]string `json:"answer" `
	EditedQuery  [][]string `json:"edited_query" binding:"required"`
	EditedAnswer [][]string `json:"edited_answer" binding:"required"`
	UserName     string     `json:"user_name"`
}
