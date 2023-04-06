package db_model

import "log"

func CreateModelFactory() *InferModel {
	return &InferModel{BaseModel: BaseModel{DB: UseDbConn()}}
}

type InferModel struct {
	BaseModel
	ModelID   string `gorm:"column:id;primaryKey"`
	ModelName string `gorm:"column:modelname" json:"model_name"`
	Version   string `gorm:"column:version" json:"version"`
	IfValid   bool   `gorm:"column:ifvalid" json:"ifvalid"`
}

// 自定义表名
func (im *InferModel) TableName() string {
	return "model"
}

func (im *InferModel) UploadModel(modelName, modelId, version string) bool {
	res := im.Create(
		&InferModel{
			ModelID:   modelId,
			ModelName: modelName,
			Version:   version,
			IfValid:   true,
		})
	if res.Error != nil {
		log.Printf("model数据库插入有误")
		return false

	} else {
		return true
	}
}

func (im *InferModel) DeleteModel(modelId string) bool {

	res := im.Model(&InferModel{ModelID: modelId}).Update("ifvalid", false)
	if res.Error != nil {
		log.Printf("model数据库更新有误")
		return false

	} else {
		return true
	}
}

func (im *InferModel) GetModel() (bool, []InferModel) {
	var models []InferModel
	res := im.Where("ifvalid = ?", true).Find(&models)
	if res.Error != nil {
		log.Printf("model数据库读取有误")
		return false, nil

	} else {
		return true, models
	}
}

func (im *InferModel) GetModelById(modelId string) (bool, *InferModel) {
	models := &InferModel{ModelID: modelId}
	res := im.First(&models)
	if res.Error != nil {
		log.Printf("model数据库读取有误")
		return false, nil

	} else {
		return true, models
	}
}
