package utils

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

const (
	machineIDBits      uint8 = 10                              //机器id占用的位数
	sequenceNumberBits uint8 = 12                              //当前毫秒内的顺序数
	epochTimestamp           = 1668009600                      //纪元，开始于北京时间 2022-11-10 00:00:00 这个值不要随便修改不然id可能重复
	machineMax         int64 = -1 ^ (-1 << machineIDBits)      //支持的最大机器id数量
	sequenceNumberMask int64 = -1 ^ (-1 << sequenceNumberBits) //支持的最大序列id数量
)

type snowflake struct {
	machineID      int64
	mu             sync.Mutex
	epoch          time.Time //纪元
	old            int64     //上次生成的雪花id毫秒时间
	sequenceNumber int64     //序列号
}

func GenSnowflakeId() (int64, error) {
	//1. New雪花id的实例
	snowflake, err := newSnowflake(1)
	if err != nil {
		return 0, err
	}
	//2.获取应该雪花id
	return snowflake.gen(), nil
}

// NewSnowflake 传入服务器编号产生一个雪花id实例
func newSnowflake(machineID int64) (*snowflake, error) {
	if machineID < 0 || machineID > machineMax {
		return nil, errors.New("machine id must be between 0 and " + strconv.FormatInt(machineMax, 10))
	}
	s := snowflake{machineID: machineID}
	s.epoch = time.Unix(epochTimestamp, 0)
	return &s, nil
}

// Gen 生成一个int64类型的雪花ID
func (s *snowflake) gen() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Since(s.epoch).Milliseconds()
	if now == s.old { //毫秒时间等于上一个毫秒时间时
		s.sequenceNumber = (s.sequenceNumber + 1) & sequenceNumberMask //序列化+1 但不能超出其最大值
		if s.sequenceNumber == 0 {                                     //但序列号溢出时，此时序列号的bit位全为0，即sequenceNumber == 0
			for now <= s.old { //保险起见在判断一次，当前毫秒时间是否真的是小于等于上一次生雪花id时的的毫秒数。
				now = time.Since(s.epoch).Milliseconds() //再重新获取一遍当前毫秒时间。
			}
		}
	} else {
		s.sequenceNumber = 0
	}
	s.old = now
	//当前毫秒数左移22位，即雪花id的那41位时间戳，或上当前机器好左移12位，即将机器好放入它本该在的那个位置，后边同理
	return now<<(machineIDBits+sequenceNumberBits) | (s.machineID << sequenceNumberBits) | s.sequenceNumber
}

//使用的算法来源于https://github.com/langwan/chihuo/blob/main/go%E8%AF%AD%E8%A8%80/snowflake,我只不过在其上添加了注释，及修改了纪元时间而已
