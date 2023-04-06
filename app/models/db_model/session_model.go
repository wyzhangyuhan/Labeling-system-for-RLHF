package db_model

import (
	"fmt"
	"label_system/utils/encode_array"
	"log"

	"gorm.io/gorm"
)

func CreateSessionFactory() *SessionModel {
	return &SessionModel{BaseModel: BaseModel{DB: UseDbConn()}}
}
func CreateItemFactory() *ItemModel {
	return &ItemModel{BaseModel: BaseModel{DB: UseDbConn()}}
}

type SessionModel struct {
	BaseModel

	SessionID int64  `gorm:"column:sessionId; primaryKey" json:"session_id"`
	DatasetId string `gorm:"column:datasetId" json:"dataset_id"`
	ModelID   string `gorm:"column:modelId" json:"model_id"`
	UserId    int64  `gorm:"column:userid" json:"user_id"`
}

func (sm *SessionModel) TableName() string {
	return "sessionitem"
}

func (sm *SessionModel) GetItem(sessionId int64) (*SessionModel, bool) {
	item := &SessionModel{SessionID: sessionId}
	res := sm.First(&item)
	if res.Error != nil {
		return nil, false
	} else {
		return item, true
	}
}

func (sm *SessionModel) GetDatasetItem(sessionId int64) (*[]DatasetModel, bool) {
	var results []DatasetModel
	sm.Model(&SessionModel{}).Select("*").Joins("left join dataset on  dataset.id = sessionitem.datasetId").Scan(&results)
	return &results, true
}

func (sm *SessionModel) NewSession(sessionid, userid int64, datasetid, modelid string) bool {
	res := sm.Create(
		&SessionModel{
			SessionID: sessionid,
			DatasetId: datasetid,
			ModelID:   modelid,
			UserId:    userid,
		})
	if res.Error != nil {
		log.Printf("sessionitem数据库插入有误")
		fmt.Printf("%v", res.Error)
		return false

	} else {
		return true
	}
}

func (sm *SessionModel) GroupByUser(begin, last string, userid int64) int {

	// var smList []SessionModel
	// res := sm.Where("userid = ? AND updatetime BETWEEN ? AND ?", userid, begin, last).Find(&smList)
	// // res := subQuery.Select("userid, count(userid) as total").Group("userid").Find(&result)
	// if res.Error != nil {
	// 	log.Printf("获取标注员总数失败")
	// 	return nil

	// } else {
	// 	return smList
	// }

	// type results struct {
	// 	edited_answer string
	// }

	// res := sm.Model(&SessionModel{}).Select("dataitem.edited_answer").Joins("left join dataitem on dataitem.sessionId = sessionitem.sessionId")
	// fmt.Printf("%v", res)
	return 0

}

type ItemModel struct {
	BaseModel

	// MessageID string `gorm:"column:messageId; primaryKey" json:"message_id"`
	SessionId    int64  `gorm:"column:sessionId; primaryKey" json:"session_id"`
	Query        string `gorm:"column:query" json:"query"`
	Answer       string `gorm:"column:answer" json:"answer"`
	EditedQuery  string `gorm:"column edited_query" json:"edited_query"`
	EditedAnswer string `gorm:"column edited_answer" json:"edited_answer"`
	UserName     string `gorm:"column:username" json:"username"`
}

func (im *ItemModel) TableName() string {
	return "dataitem"
}

func (im *ItemModel) GetItem(sessionid int64) *ItemModel {
	var item ItemModel = ItemModel{SessionId: sessionid}
	im.First(&item)
	return &item
}

func (im *ItemModel) BrutalUpdateItem(query, answer, edit_q, edit_a, username string, sessionid int64) bool {
	var item ItemModel = ItemModel{SessionId: sessionid}
	res := im.Model(&item).Updates(ItemModel{Query: query, Answer: answer, EditedQuery: edit_q, EditedAnswer: edit_a})
	if res.Error != nil {
		log.Printf("dataitem数据库插入有误")
		return false
	} else {
		return true
	}
}

func (im *ItemModel) UpdateItem(query, answer, edit_q, edit_a, username string, sessionid int64) bool {
	var item ItemModel = ItemModel{SessionId: sessionid}
	res := im.First(&item)
	var up_res *gorm.DB
	if res.Error == nil {
		new_query, new_answer, new_eq, new_ea := CombineRes(item, query, answer, edit_q, edit_a)
		fmt.Printf("Query: %v;\n Answer: %v;\n Edited_query: %v;\n Edited_answer:%v\n", new_query, new_answer, new_eq, new_ea)
		up_res = im.Model(&item).Updates(ItemModel{Query: new_query, Answer: new_answer, EditedQuery: new_eq, EditedAnswer: new_ea})
	} else {
		up_res = im.Create(
			&ItemModel{
				SessionId:    sessionid,
				Query:        query,
				Answer:       answer,
				EditedQuery:  edit_q,
				EditedAnswer: edit_a,
				UserName:     username,
			})
	}

	if up_res.Error != nil {
		log.Printf("dataitem数据库插入有误")
		return false

	} else {
		return true
	}
}

func (im *ItemModel) AddMoreItem(query, answer, edit_q, edit_a, username string, sessionid int64) bool {
	var item ItemModel = ItemModel{SessionId: sessionid}
	res := im.First(&item)
	var up_res *gorm.DB
	if res.Error == nil {
		new_query, new_answer, new_eq, new_ea := AddMoreRes(item, query, answer, edit_q, edit_a)
		fmt.Printf("Query: %v;\n Answer: %v;\n Edited_query: %v;\n Edited_answer:%v\n", new_query, new_answer, new_eq, new_ea)
		up_res = im.Model(&item).Updates(ItemModel{Query: new_query, Answer: new_answer, EditedQuery: new_eq, EditedAnswer: new_ea})
	} else {
		return false
	}

	if up_res.Error != nil {
		log.Printf("dataitem数据库插入有误")
		return false

	} else {
		return true
	}
}

func (im *ItemModel) ChangeTheItem(query, answer, edit_q, edit_a, username string, sessionid int64) bool {
	var item ItemModel = ItemModel{SessionId: sessionid}
	res := im.First(&item)
	var up_res *gorm.DB
	if res.Error == nil {
		new_query, new_answer, new_eq, new_ea := ChangeRes(item, query, answer, edit_q, edit_a)
		fmt.Printf("Query: %v;\n Answer: %v;\n Edited_query: %v;\n Edited_answer:%v\n", new_query, new_answer, new_eq, new_ea)
		up_res = im.Model(&item).Updates(ItemModel{Query: new_query, Answer: new_answer, EditedQuery: new_eq, EditedAnswer: new_ea})
	} else {
		return false
	}

	if up_res.Error != nil {
		log.Printf("dataitem数据库插入有误")
		return false

	} else {
		return true
	}
}

func (im *ItemModel) QueryByUpdateTimeRange(begin, last string) []ItemModel {
	var imlist []ItemModel
	res := im.Where("updatetime BETWEEN ? AND ?", begin, last).Find(&imlist)
	if res.Error != nil {
		fmt.Printf("%v", res.Error)
	}
	return imlist
}

func (im *ItemModel) GroupByUser(username, begin, last string) []ItemModel {
	var imlist []ItemModel
	res := im.Where("username = ? AND updatetime BETWEEN ? AND ?", username, begin, last).Find(&imlist)
	if res.Error != nil {
		fmt.Printf("%v", res.Error)
	}
	return imlist
}

func CombineRes(item ItemModel, query, answer, edit_q, edit_a string) (string, string, string, string) {
	dj := encode_array.StringArrayDecoder([]byte(query), []byte(answer), []byte(edit_q), []byte(edit_a))
	item_dj := encode_array.StringArrayDecoder([]byte(item.Query), []byte(item.Answer), []byte(item.EditedQuery), []byte(item.EditedAnswer))
	//remove duplicate
	if len(dj.EQ) != 0 {
		dj.EQ = removeDuplicate(item_dj.Q, dj.EQ)
	}
	if len(dj.EA) != 0 {
		dj.EA = removeDuplicate(item_dj.A, dj.EA)
	}
	item_dj.Q = append(item_dj.Q, dj.Q...)
	item_dj.A = append(item_dj.A, dj.A...)
	item_dj.EQ = append(item_dj.EQ, dj.EQ...)
	item_dj.EA = append(item_dj.EA, dj.EA...)

	return encode_array.StringArrayEncoder("query", item_dj.Q),
		encode_array.StringArrayEncoder("answer", item_dj.A),
		encode_array.StringArrayEncoder("edited_query", item_dj.EQ),
		encode_array.StringArrayEncoder("edited_answer", item_dj.EA)
}

func AddMoreRes(item ItemModel, query, answer, edit_q, edit_a string) (string, string, string, string) {
	dj := encode_array.StringArrayDecoder([]byte(query), []byte(answer), []byte(edit_q), []byte(edit_a))
	item_dj := encode_array.StringArrayDecoder([]byte(item.Query), []byte(item.Answer), []byte(item.EditedQuery), []byte(item.EditedAnswer))

	if len(dj.A) != 0 {
		item_dj.A[len(item_dj.A)-1] = append(item_dj.A[len(item_dj.A)-1], dj.A[0]...)
	} else {
		fmt.Printf("Add more err")
	}
	return encode_array.StringArrayEncoder("query", item_dj.Q),
		encode_array.StringArrayEncoder("answer", item_dj.A),
		encode_array.StringArrayEncoder("edited_query", item_dj.EQ),
		encode_array.StringArrayEncoder("edited_answer", item_dj.EA)
}

func ChangeRes(item ItemModel, query, answer, edit_q, edit_a string) (string, string, string, string) {
	dj := encode_array.StringArrayDecoder([]byte(query), []byte(answer), []byte(edit_q), []byte(edit_a))
	item_dj := encode_array.StringArrayDecoder([]byte(item.Query), []byte(item.Answer), []byte(item.EditedQuery), []byte(item.EditedAnswer))

	if len(dj.A) != 0 {
		item_dj.A[len(item_dj.A)-1] = dj.A[0]
	} else {
		fmt.Printf("Add more err")
	}
	return encode_array.StringArrayEncoder("query", item_dj.Q),
		encode_array.StringArrayEncoder("answer", item_dj.A),
		encode_array.StringArrayEncoder("edited_query", item_dj.EQ),
		encode_array.StringArrayEncoder("edited_answer", item_dj.EA)
}

func removeDuplicate(origin, edited [][]string) [][]string {

	length := len(origin)
	if length == 0 {
		return edited
	}
	for idx, v := range origin[length-1] {
		if idx >= len(edited[0]) {
			break
		}
		if v == edited[0][idx] {
			edited[0][idx] = ""
		}
	}
	return edited
}
