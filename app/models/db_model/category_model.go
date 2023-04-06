package db_model

import "log"

func CreateCateFactory() *CategoryModel {
	return &CategoryModel{BaseModel: BaseModel{DB: UseDbConn()}}
}

type CategoryModel struct {
	BaseModel
	CateID   string `gorm:"column:id;primaryKey"`
	ParentID string `gorm:"column:parentId"`
	CateName string `gorm:"column:catename"`
}

// 自定义表名
func (u *CategoryModel) TableName() string {
	return "datacate"
}

func (cm *CategoryModel) AddCategory(cateid, parentId, catename string) bool {
	res := cm.Create(
		&CategoryModel{
			CateID:   cateid,
			ParentID: parentId,
			CateName: catename,
		})
	if res.Error != nil {
		log.Printf("datacate数据库插入有误")
		return false

	} else {
		return true
	}
}

func (cm *CategoryModel) GetCategory() (bool, []CategoryModel) {
	var cates []CategoryModel
	res := cm.Find(&cates)
	if res.Error != nil {
		log.Printf("datacate数据库读取有误")
		return false, nil

	} else {
		return true, cates
	}
}
