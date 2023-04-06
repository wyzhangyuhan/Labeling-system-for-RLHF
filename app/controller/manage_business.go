package controller

import (
	"label_system/app/curd/dataset_curd"
	"label_system/app/curd/infer_curd"
	"label_system/app/models/db_model"
	"label_system/app/models/req_model"
	"label_system/app/response"
	"label_system/global/consts"
	"label_system/scheduler"

	"github.com/gin-gonic/gin"
)

func UploadModel(c *gin.Context) {
	var model req_model.ModelReq
	if err := c.ShouldBindJSON(&model); err != nil {
		response.Fail(c, consts.BackendError, "", "")
		return
	}

	ok, modelId := infer_curd.CreateModelCurdFactory().UploadModel(model.ModelName, model.Version)
	if !ok {
		response.Fail(c, consts.BackendError, "", "")
		return
	}
	response.Success(c, "", map[string]string{"modelid": modelId})
}

func DeleteModel(c *gin.Context) {
	var model req_model.ModelReq
	if err := c.ShouldBindJSON(&model); err != nil {
		response.Fail(c, consts.BackendError, "", "")
		return
	}

	ok := infer_curd.CreateModelCurdFactory().DeleteModel(model.ModelId)
	if !ok {
		response.Fail(c, consts.BackendError, "", "")
		return
	}
	response.Success(c, "", "")
}

func GetValidModel(c *gin.Context) {
	ok, validModels := infer_curd.CreateModelCurdFactory().GetValidModel()
	if !ok {
		response.Fail(c, consts.BackendError, "", "")
		return
	}
	response.Success(c, "", validModels)
}

func CategoryAdded(c *gin.Context) {
	var cate req_model.CateReq
	if err := c.ShouldBindJSON(&cate); err != nil {
		response.Fail(c, consts.BackendError, "", "")
		return
	}
	ok, cateid := dataset_curd.CreateCateCurdFactory().AddCate(cate.CateName, cate.ParentId)
	if !ok {
		response.Fail(c, consts.BackendError, "", "")
		return
	}
	response.Success(c, "", map[string]string{"category_id": cateid})
}

func DatasetAdded(c *gin.Context) {
	var data req_model.DataSetReq
	if err := c.ShouldBindJSON(&data); err != nil {
		response.Fail(c, consts.BackendError, "", "")
		return
	}

	ok, dataid := dataset_curd.CreateDatasetCurdFactory().AddDataset(data.DataName, data.CategoryId, data.DataDes)
	if !ok {
		response.Fail(c, consts.BackendError, "", "")
		return
	}
	response.Success(c, "", map[string]string{"data_id": dataid})
}

func GetDataset(c *gin.Context) {

	type CateCascader struct {
		CateID   string                  `json:"cate_id"`
		ParentID string                  `json:"parent_id"`
		CateName string                  `json:"cate_name"`
		Children []CateCascader          `json:"children"`
		DataSets []db_model.DatasetModel `json:"datasets"`
	}

	ok, catelist := dataset_curd.CreateCateCurdFactory().GetAllCateWithStruct()
	if !ok {
		response.Fail(c, consts.BackendError, "", "")
		return
	}
	var cate map[string]*CateCascader = make(map[string]*CateCascader)
	for _, v := range catelist {
		if v.ParentID == "" {
			_, datalist := dataset_curd.CreateDatasetCurdFactory().GetDatasetFromCate(v.CateID)
			var tmp = &CateCascader{CateID: v.CateID, ParentID: v.ParentID, CateName: v.CateName, Children: nil, DataSets: datalist}
			cate[v.CateID] = tmp
		}
	}
	for _, v := range catelist {
		if v.ParentID != "" {
			_, datalist := dataset_curd.CreateDatasetCurdFactory().GetDatasetFromCate(v.CateID)
			tmp := cate[v.ParentID]
			var child *CateCascader = &CateCascader{CateID: v.CateID, ParentID: v.ParentID, CateName: v.CateName, Children: nil, DataSets: datalist}
			tmp.Children = append(tmp.Children, *child)
			cate[v.ParentID] = tmp
		}
	}

	response.Success(c, "", cate)
}

func Save2OSS(c *gin.Context) {
	var timePoint struct {
		Begin string `json:"begin"`
		End   string `json:"end"`
	}
	c.ShouldBindJSON(&timePoint)
	scheduler.ExportLabel(timePoint.Begin, timePoint.End)
	//20230315000000
}
