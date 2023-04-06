package controller

import (
	"label_system/app/client"
	"label_system/app/curd/dataset_curd"
	"label_system/app/curd/labeled"
	"label_system/app/curd/users"
	"label_system/app/models/req_model"
	"label_system/app/oss_curd"
	"label_system/app/response"
	"label_system/global/consts"
	"label_system/utils/encode_array"
	"label_system/utils/genid"
	"label_system/utils/snowflakes"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AskModel(c *gin.Context) { //获得的模型推理结果
	// Input:
	// 	modelid
	// 	sessionid
	//  query & answer

	var Que req_model.SubmitQuestionReq
	if err := c.ShouldBindJSON(&Que); err != nil {
		response.Fail(c, consts.BackendError, "", err)
		return
	}
	var infer client.InferGateway
	infer.HttpClient = &http.Client{Timeout: 120 * time.Second}
	res, err := infer.ChatModel(&client.InferReq{Que.ModelId, Que.Query, Que.Answer})
	if err != nil {
		response.Fail(c, consts.BackendError, "推理任务失败", err)
		return
	}
	res.MessageId = genid.GenSnowflakeString()
	sessionId_int64, _ := strconv.ParseInt(Que.SessionId, 10, 64)

	var ok bool
	var tmp = [][]string{res.Answers}
	tmpQ := []string{Que.Query[len(Que.Query)-1]}
	edit_tmpQ := []string{""}

	if len(Que.Query) > len(Que.Answer) {
		var oka, okb bool = true, true
		var tmpA [][]string
		if Que.IsFirst {
			if len(Que.Answer)-1 > 0 {
				tmpAA := []string{Que.Answer[len(Que.Answer)-1]}
				tmpA = [][]string{tmpAA}
			} else {
				tmpA = [][]string{}
			}
			dataItem := labeled.CreateItemCurdFactory().GetItem(sessionId_int64)
			if len(Que.Query) > len(dataItem.Q) { //初始数据以回答结尾，要先塞进去前端传入倒数第二个问题，
				tmpQ_2 := []string{Que.Query[len(Que.Query)-2]}
				oka = labeled.CreateItemCurdFactory().UpdateItem([][]string{}, [][]string{}, [][]string{tmpQ_2}, tmpA, sessionId_int64, Que.UserName) //单独先更新修改后的ans
				okb = labeled.CreateItemCurdFactory().UpdateItem([][]string{tmpQ}, tmp, [][]string{edit_tmpQ}, [][]string{}, sessionId_int64, Que.UserName)
			} else { //初始数据以问题结尾
				oka = labeled.CreateItemCurdFactory().UpdateItem([][]string{}, [][]string{}, [][]string{}, tmpA, sessionId_int64, Que.UserName) //单独先更新修改后的ans
				okb = labeled.CreateItemCurdFactory().UpdateItem([][]string{}, tmp, [][]string{tmpQ}, [][]string{}, sessionId_int64, Que.UserName)
			}

		} else {
			okb = labeled.CreateItemCurdFactory().UpdateItem([][]string{tmpQ}, tmp, [][]string{edit_tmpQ}, [][]string{}, sessionId_int64, Que.UserName)
		}
		ok = oka && okb
	} else {
		response.Fail(c, consts.BackendError, "你不问问题咋回答", "")
		return
	}
	if !ok {
		response.Fail(c, consts.BackendError, "数据库更新失败", "")
		return
	}

	response.Success(c, "Success", res)
}

func AskForMore(c *gin.Context) {
	var Que req_model.SubmitQuestionReq
	if err := c.ShouldBindJSON(&Que); err != nil {
		response.Fail(c, consts.BackendError, "", err)
		return
	}
	var infer client.InferGateway
	infer.HttpClient = &http.Client{Timeout: 120 * time.Second}
	res, err := infer.ChatModel(&client.InferReq{Que.ModelId, Que.Query, Que.Answer})
	if err != nil {
		response.Fail(c, consts.BackendError, "推理任务失败", err)
		return
	}
	res.MessageId = genid.GenSnowflakeString()
	sessionId_int64, _ := strconv.ParseInt(Que.SessionId, 10, 64)
	var tmp = [][]string{res.Answers}

	ok := labeled.CreateItemCurdFactory().AddMoreItem([][]string{}, tmp, [][]string{}, [][]string{}, sessionId_int64, Que.UserName)

	if !ok {
		response.Fail(c, consts.BackendError, "数据库更新失败", "")
		return
	}

	response.Success(c, "Success", res)
}

func ChangeTheRound(c *gin.Context) {
	var Que req_model.SubmitQuestionReq
	if err := c.ShouldBindJSON(&Que); err != nil {
		response.Fail(c, consts.BackendError, "", err)
		return
	}
	var infer client.InferGateway
	infer.HttpClient = &http.Client{Timeout: 120 * time.Second}
	res, err := infer.ChatModel(&client.InferReq{Que.ModelId, Que.Query, Que.Answer})
	if err != nil {
		response.Fail(c, consts.BackendError, "推理任务失败", err)
		return
	}
	res.MessageId = genid.GenSnowflakeString()
	sessionId_int64, _ := strconv.ParseInt(Que.SessionId, 10, 64)
	var tmp = [][]string{res.Answers}

	ok := labeled.CreateItemCurdFactory().ChangeTheItem([][]string{}, tmp, [][]string{}, [][]string{}, sessionId_int64, Que.UserName)

	if !ok {
		response.Fail(c, consts.BackendError, "数据库更新失败", "")
		return
	}

	response.Success(c, "Success", res)
}

func PureAsk(c *gin.Context) {
	var req client.InferReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.BackendError, "", err)
		return
	}
	var infer client.InferGateway
	infer.HttpClient = &http.Client{Timeout: 120 * time.Second}
	res, err := infer.ChatModel(&req)
	if err != nil {
		response.Fail(c, consts.BackendError, "推理任务失败", err)
		return
	}

	response.Success(c, "Success", res)

}

func StartSession(c *gin.Context) {
	var session req_model.SessionReq

	if err := c.ShouldBindJSON(&session); err != nil {
		response.Fail(c, consts.BackendError, "", err)
		return
	}

	sessionid := snowflakes.CreateSnowflakeFactory().GenId()
	userid_int64, err := strconv.ParseInt(session.UserId, 10, 64)
	_, userInfo := users.CreateUserCurdFactory().GetUserInfo(userid_int64)
	session.UserName = userInfo.UserName
	if err != nil {
		response.Fail(c, consts.BackendError, "", err)
		return
	}
	ok := labeled.CreateSessionCurdFactory().NewSession(sessionid, userid_int64, session.DatesetId, session.ModelId)
	if !ok {
		response.Fail(c, consts.BackendError, "Cannot new session", "")
		return
	}

	// 获取问题
	_, datasetitem := dataset_curd.CreateDatasetCurdFactory().GetDatasetItemById(session.DatesetId)
	session.DataSetName = datasetitem.DataName
	quelist := oss_curd.Oc.LoadDataset(session.DataSetName)
	if len(quelist) != 1 {
		response.Fail(c, consts.BackendError, "Get Que wrong", "")
		return
	}
	ret_que := quelist[0]

	//对生成的多个问题和回答在edited字段中，填充空字符串占位
	q := encode_array.UnfoldStringArray(ret_que.Question)
	a := encode_array.UnfoldStringArray(ret_que.Answer)
	var edit_q, edit_a [][]string

	edit_q = make([][]string, len(q)-1)
	if len(a)-1 <= 0 {
		edit_a = make([][]string, 0)
	} else {
		edit_a = make([][]string, len(a)-1)
	}

	for i := range edit_q {
		edit_q[i] = []string{""}
	}
	for i := range edit_a {
		edit_a[i] = []string{""}
	}

	if ok := labeled.CreateItemCurdFactory().UpdateItem(q, a, edit_q, edit_a, sessionid, session.UserName); !ok {
		response.Fail(c, consts.BackendError, "", "")
		return
	}

	sessRet := &req_model.SessionRsp{session, strconv.FormatInt(sessionid, 10), ret_que}
	response.Success(c, "开始会话", sessRet)
}

func GetLabelNumPerP(c *gin.Context) {
	var req struct {
		Userid string `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.BackendError, "", err)
		return
	}
	userid_int64, err := strconv.ParseInt(req.Userid, 10, 64)
	_, userInfo := users.CreateUserCurdFactory().GetUserInfo(userid_int64)
	if err != nil {
		response.Fail(c, consts.BackendError, "", err)
		return
	}
	// num := labeled.CreateSessionCurdFactory().GroupByUser(userid_int64)
	num := labeled.CreateItemCurdFactory().GroupByUser(userInfo.UserName)

	response.Success(c, "Success", map[string]int{"num": num})
}

func SubmitRes(c *gin.Context) {
	//提交编辑后数据
	//前端提交格式为: {query:[[],[],[]...], answers:[[],[],[]...]},
	var content req_model.SubmitLabelReq
	if err := c.ShouldBindJSON(&content); err != nil {
		response.Fail(c, consts.BackendError, "", err)
		return
	}
	sessionid_int64, err := strconv.ParseInt(content.SessionId, 10, 64)
	if err != nil {
		response.Fail(c, consts.BackendError, "", err)
		return
	}

	//入库
	if ok := labeled.CreateItemCurdFactory().UpdateItem([][]string{}, [][]string{}, content.EditedQuery, content.EditedAnswer, sessionid_int64, content.UserName); !ok {
		response.Fail(c, consts.BackendError, "", "")
		return
	}

	response.Success(c, "Success", content)
}
