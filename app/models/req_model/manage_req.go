package req_model

type ModelReq struct {
	ModelName string `json:"model_name"`
	ModelId   string `json:"model_id"`
	Version   string `json:"version"`
}

type CateReq struct {
	CateName string `json:"category_name"`
	ParentId string `json:"parent_id"`
}

type DataSetReq struct {
	DataName   string `json:"data_name"`
	CategoryId string `json:"category_id"`
	DataDes    string `json:"description"`
}
