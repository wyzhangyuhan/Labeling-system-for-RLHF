package req_model

type SubmitQuestionReq struct {
	ModelId   string   `json:"model_id" binding:"required"`
	SessionId string   `json:"session_id" binding:"required"`
	UserName  string   `json:"user_name" binding:"required"`
	Query     []string `json:"query" binding:"required"`
	Answer    []string `json:"answer" binding:"required"`
	IsFirst   bool     `json:"isfirst" `
}
type SubmitQuestionRsp struct {
	MessageId string   `json:"MessageId"`
	Answers   []string `json:"answers"`
}
