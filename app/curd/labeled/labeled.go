package labeled

import (
	"fmt"
	"label_system/app/models/db_model"
	"label_system/global/consts"
	"label_system/utils/encode_array"
	"time"
)

func CreateSessionCurdFactory() *SessionCurd {
	return &SessionCurd{db_model.CreateSessionFactory()}
}

func CreateItemCurdFactory() *ItemCurd {
	return &ItemCurd{db_model.CreateItemFactory()}
}

type SessionCurd struct {
	sessionModel *db_model.SessionModel
}

func (s *SessionCurd) NewSession(sessionid, userid int64, datasetid, modelid string) bool {
	return s.sessionModel.NewSession(sessionid, userid, datasetid, modelid)
}

func (s *SessionCurd) GroupByUser(userid int64) int {
	//生成每日至现在的时间
	begin := time.Now()
	begin = time.Date(begin.Year(), begin.Month(), begin.Day(), 0, 0, 0, 0, begin.Location())
	dayBegin := begin.Format(consts.TimeFormat)
	end := time.Now().Format(consts.TimeFormat)
	fmt.Printf("begin:%v, end:%v \n", dayBegin, end)
	return s.sessionModel.GroupByUser(dayBegin, end, userid)

}

func (s *SessionCurd) GetItem(sessionid int64) (*db_model.SessionModel, bool) {
	return s.sessionModel.GetItem(sessionid)
}

type ItemCurd struct {
	itemModel *db_model.ItemModel
}

func (i *ItemCurd) UpdateItem(query, answer, edit_q, edit_a [][]string, sessionid int64, username string) bool {

	encodedQuery := updateItemEncoder("query", query)
	encodedAns := updateItemEncoder("answer", answer)
	encodedEQ := updateItemEncoder("edited_query", edit_q)
	encodedEA := updateItemEncoder("edited_answer", edit_a)

	return i.itemModel.UpdateItem(encodedQuery, encodedAns, encodedEQ, encodedEA, username, sessionid)
}

func (i *ItemCurd) AddMoreItem(query, answer, edit_q, edit_a [][]string, sessionid int64, username string) bool {
	encodedQuery := updateItemEncoder("query", query)
	encodedAns := updateItemEncoder("answer", answer)
	encodedEQ := updateItemEncoder("edited_query", edit_q)
	encodedEA := updateItemEncoder("edited_answer", edit_a)

	return i.itemModel.AddMoreItem(encodedQuery, encodedAns, encodedEQ, encodedEA, username, sessionid)
}

func (i *ItemCurd) ChangeTheItem(query, answer, edit_q, edit_a [][]string, sessionid int64, username string) bool {
	encodedQuery := updateItemEncoder("query", query)
	encodedAns := updateItemEncoder("answer", answer)
	encodedEQ := updateItemEncoder("edited_query", edit_q)
	encodedEA := updateItemEncoder("edited_answer", edit_a)

	return i.itemModel.ChangeTheItem(encodedQuery, encodedAns, encodedEQ, encodedEA, username, sessionid)
}

func (i *ItemCurd) BrutalUpdateItem(query, answer, edit_q, edit_a [][]string, sessionid int64, username string) bool {

	encodedQuery := updateItemEncoder("query", query)
	encodedAns := updateItemEncoder("answer", answer)
	encodedEQ := updateItemEncoder("edited_query", edit_q)
	encodedEA := updateItemEncoder("edited_answer", edit_a)

	return i.itemModel.UpdateItem(encodedQuery, encodedAns, encodedEQ, encodedEA, username, sessionid)
}

func (i *ItemCurd) GetItem(sessionid int64) *encode_array.DecodeJson {
	dataitem := i.itemModel.GetItem(sessionid)
	return encode_array.StringArrayDecoder([]byte(dataitem.Query), []byte(dataitem.Answer), []byte(dataitem.EditedQuery), []byte(dataitem.EditedAnswer))
}

func (i *ItemCurd) QueryTimeRange(begin, end string) []db_model.ItemModel {
	return i.itemModel.QueryByUpdateTimeRange(begin, end)
}

func (i *ItemCurd) GroupByUser(username string) (num int) {
	//生成每日至现在的时间
	begin := time.Now()
	begin = time.Date(begin.Year(), begin.Month(), begin.Day(), 0, 0, 0, 0, begin.Location())
	dayBegin := begin.Format(consts.TimeFormat)
	end := time.Now().Format(consts.TimeFormat)
	fmt.Printf("begin:%v, end:%v \n", dayBegin, end)
	res := i.itemModel.GroupByUser(username, dayBegin, end)
	num = len(res)
	for _, v := range res {
		dj := encode_array.StringArrayDecoder([]byte(v.EditedAnswer), []byte(v.Answer))
		if len(dj.EA) < 1 || len(dj.EA) < len(dj.A) {
			num--
		}
	}
	return
}

func updateItemEncoder(key string, strArray [][]string) (encodeStr string) {
	if len(strArray) != 0 {
		encodeStr = encode_array.StringArrayEncoder(key, strArray)
	} else {
		encodeStr = ""
	}
	return
}
