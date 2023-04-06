package dataset_curd

import (
	"label_system/app/models/db_model"
	"label_system/utils/genid"
)

func CreateCateCurdFactory() *CategoryCurd {
	return &CategoryCurd{db_model.CreateCateFactory()}
}

func CreateDatasetCurdFactory() *DatasetCurd {
	return &DatasetCurd{db_model.CreateDatasetFactory()}
}

type CategoryCurd struct {
	categoryModel *db_model.CategoryModel
}

func (c *CategoryCurd) AddCate(catename, parentid string) (bool, string) {
	cateId := genid.GenSpecificId(catename)
	return c.categoryModel.AddCategory(cateId, parentid, catename), cateId
}

func (c *CategoryCurd) GetAllCateWithStruct() (bool, []db_model.CategoryModel) {

	ok, catelist := c.categoryModel.GetCategory()
	if !ok {
		return false, nil
	}

	return true, catelist
}

type DatasetCurd struct {
	datasetModel *db_model.DatasetModel
}

func (d *DatasetCurd) AddDataset(dataname, cateid, description string) (bool, string) {
	datasetId := genid.GenSpecificId(dataname)
	return d.datasetModel.UploadDataset(dataname, datasetId, cateid, description), datasetId
}

func (d *DatasetCurd) GetDatasetFromCate(cateid string) (bool, []db_model.DatasetModel) {

	return d.datasetModel.GetFromCate(cateid)
}

func (d *DatasetCurd) GetDatasetItemById(datasetid string) (bool, *db_model.DatasetModel) {

	return d.datasetModel.GetDatasetItemById(datasetid)
}
