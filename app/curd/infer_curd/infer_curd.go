package infer_curd

import (
	"label_system/app/models/db_model"
	"label_system/utils/genid"
)

func CreateModelCurdFactory() *ModelCurd {
	return &ModelCurd{db_model.CreateModelFactory()}
}

type ModelCurd struct {
	inferModel *db_model.InferModel
}

func (m *ModelCurd) UploadModel(modelname, version string) (bool, string) {
	modelId := genid.GenSpecificId(modelname)
	return m.inferModel.UploadModel(modelname, modelId, version), modelId
}

func (m *ModelCurd) DeleteModel(modelId string) bool {
	return m.inferModel.DeleteModel(modelId)
}

func (m *ModelCurd) GetValidModel() (bool, []db_model.InferModel) {
	return m.inferModel.GetModel()
}

func (m *ModelCurd) GetModelById(modelId string) (bool, *db_model.InferModel) {
	return m.inferModel.GetModelById(modelId)
}
