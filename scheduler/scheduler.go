package scheduler

import (
	"encoding/json"
	"fmt"
	"label_system/app/curd/infer_curd"
	"label_system/app/curd/labeled"
	"label_system/app/oss_curd"
	"label_system/global/consts"
	"label_system/utils/encode_array"
	"log"
	"time"
)

type saveForm struct {
	encode_array.DecodeJson
	Model string `json:"model"`
}

func DailyLabelExport() {
	for {
		now := time.Now()                                                                    //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24)                                                      //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location()) //获取下一个凌晨的日期
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		t := time.NewTimer(next.Sub(now)) //计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C

		begin, end := start.Format(consts.TimeFormat), next.Format(consts.TimeFormat)
		ExportLabel(begin, end)
	}
}

func getModelVersion(sessionid int64) string {
	sessionContent, _ := labeled.CreateSessionCurdFactory().GetItem(sessionid)
	_, modelItem := infer_curd.CreateModelCurdFactory().GetModelById(sessionContent.ModelID)
	return modelItem.ModelName + "-" + modelItem.Version
}

func ExportLabel(begin, end string) {
	fmt.Printf("Data Exporting~ Begin: %v; End: %v\n", begin, end)
	userfile := make(map[string]string)
	items := labeled.CreateItemCurdFactory().QueryTimeRange(begin, end)
	for _, v := range items {
		dj := encode_array.StringArrayDecoder([]byte(v.Query), []byte(v.Answer), []byte(v.EditedQuery), []byte(v.EditedAnswer))
		dj.User = v.UserName
		modelName := getModelVersion(v.SessionId)
		sf := &saveForm{*dj, modelName}
		label, err := json.Marshal(sf)
		if err != nil {
			log.Printf("err when converting %v json string", v.SessionId)
		}
		if _, ok := userfile[dj.User]; ok {
			userfile[dj.User] = userfile[dj.User] + string(label) + "\n"
		} else {
			userfile[dj.User] = string(label) + "\n"
		}
	}
	for username, jsonfile := range userfile {
		oss_curd.Oc.SaveLabeled(end, username, jsonfile)
	}

}
