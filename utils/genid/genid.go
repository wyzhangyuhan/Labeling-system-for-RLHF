package genid

import (
	"fmt"
	"label_system/global/consts"
	"label_system/utils/snowflakes"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

var timeZone = time.FixedZone("CST", 8*3600)

func GenRequestId() string {
	return time.Now().In(timeZone).Format(consts.TimeFormat) + "-" + randstr.String(4)
}

func GenSnowflakeString() string {
	return fmt.Sprint(snowflakes.CreateSnowflakeFactory().GenId()) + "-Msg" + randstr.String(11)
}

func GenerateInt64() int64 {
	idStr := time.Now().In(timeZone).Format(consts.TimeFormat) + randstr.Hex(4)
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}

func GenSpecificId(specificName string) string {
	newid := uuid.NewMD5(uuid.New(), []byte(specificName))
	return newid.String()
}
