package db_model

import "log"

func CreateDatasetFactory() *DatasetModel {
	return &DatasetModel{BaseModel: BaseModel{DB: UseDbConn()}}
}

type DatasetModel struct {
	BaseModel

	ID          string `gorm:"column:id;primaryKey" json:"dataset_id"`
	DataName    string `gorm:"column:dataname" json:"data_name"`
	CateID      string `gorm:"column:cateid"  json:"category_id"`
	Description string `gorm:"column:description" json:"description"`
}

func (dm *DatasetModel) TableName() string {
	return "dataset"
}

func (dm *DatasetModel) UploadDataset(datasetname, datasetId, cateid, description string) bool {
	res := dm.Create(
		&DatasetModel{
			ID:          datasetId,
			DataName:    datasetname,
			CateID:      cateid,
			Description: description,
		})
	if res.Error != nil {
		log.Printf("dataset数据库插入有误")
		return false

	} else {
		return true
	}
}

func (dm *DatasetModel) GetFromCate(cateid string) (bool, []DatasetModel) {
	var datasets []DatasetModel
	res := dm.Where("cateid = ?", cateid).Find(&datasets)
	if res.Error != nil {
		log.Printf("dataset数据库查找有误")
		return false, nil

	} else {
		return true, datasets
	}
}

func (dm *DatasetModel) GetDatasetItemById(datasetid string) (bool, *DatasetModel) {
	datasets := &DatasetModel{ID: datasetid}
	res := dm.First(&datasets)
	if res.Error != nil {
		log.Printf("dataset数据库查找有误")
		return false, nil

	} else {
		return true, datasets
	}
}
