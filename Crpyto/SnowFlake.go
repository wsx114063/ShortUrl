package Crypto

import (
	"sync"
	"time"
)

// Snowflake 结构定义
type Snowflake struct {
	mu            sync.Mutex
	lastTimestamp int64
	sequence      int64
	nodeID        int64
}

// NewSnowflake 创建新的 Snowflake 实例
func NewSnowflake(nodeID int64) *Snowflake {
	return &Snowflake{
		lastTimestamp: 0,
		sequence:      0,
		nodeID:        nodeID,
	}
}

// NextID 生成下一个唯一ID
func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	timestamp := time.Now().UnixNano() / 1000000 // 将时间戳调整到毫秒级别

	if timestamp < s.lastTimestamp {
		panic("Clock moved backwards")
	}

	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & 4095
		if s.sequence == 0 {
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	id := (timestamp << 22) | (s.nodeID << 12) | s.sequence
	return id
}
