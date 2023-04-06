package consts

const (
	TimeFormat string = "20060102150405"
	//SnowFlake 雪花算法
	StartTimeStamp = int64(1672502400000) //开始时间截 (2017-01-01)
	MachineIdBits  = uint(10)             //机器id所占的位数
	SequenceBits   = uint(12)             //序列所占的位数
	//MachineIdMax   = int64(-1 ^ (-1 << MachineIdBits)) //支持的最大机器id数量
	SequenceMask   = int64(-1 ^ (-1 << SequenceBits)) //
	MachineIdShift = SequenceBits                     //机器id左移位数
	TimestampShift = SequenceBits + MachineIdBits     //时间戳左移位数

	//状态码
	BackendError   int = 500
	BackendSuccess int = 0
	CurdUserError  int = 10010

	//模型字符串
	UserChar string = "user："
	BotChar  string = "bot："
	SepChar  string = "<sep>"
)
