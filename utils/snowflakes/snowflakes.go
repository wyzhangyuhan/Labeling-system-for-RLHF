package snowflakes

import (
	"label_system/global/consts"
	"sync"
	"time"
)

type SnowFlaker interface {
	GenId() int64
}

// 创建一个雪花算法生成器(生成工厂)
func CreateSnowflakeFactory() SnowFlaker {
	return &SnowFlakes{
		timestamp: 0,
		machineId: int64(123),
		sequence:  0,
	}
}

type SnowFlakes struct {
	sync.Mutex
	timestamp int64
	machineId int64
	sequence  int64
}

func (s *SnowFlakes) GenId() int64 {
	s.Lock()
	defer func() {
		s.Unlock()
	}()
	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & consts.SequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = now
	r := (now-consts.StartTimeStamp)<<consts.TimestampShift | (s.machineId << consts.MachineIdShift) | (s.sequence)
	return r

}
