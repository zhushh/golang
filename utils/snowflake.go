package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	dataCenterBits  = 5
	workerBits      = 5
	sequenceBits    = 12
	workerIdShift   = sequenceBits
	dataCenterShift = workerIdShift + workerBits
	timeShift       = dataCenterShift + dataCenterBits
	sequenceMask    = 0xFFF
)

type Snowflake struct {
	TimeMsId     int64
	DataCenterId int64
	WorkerId     int64
	SequenceId   int64
	mutex        sync.Mutex
}

// Snowflake初始化
func (s *Snowflake) Initial(dataCenter, worker int64) {
	s.DataCenterId = dataCenter
	s.WorkerId = worker
	s.TimeMsId = time.Now().Unix()
	s.SequenceId = 0
}

// 生成下一个id
func (s *Snowflake) NextId() (int64, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var id int64
	timeMs := time.Now().Unix()
	if timeMs < s.TimeMsId {
		return 0, fmt.Errorf("timer invalid")
	}

	if timeMs == s.TimeMsId { // 同一时刻生成
		s.SequenceId = (s.SequenceId + 1) & sequenceMask
		if s.SequenceId == 0 { /// 毫秒内序列溢出，阻塞到下一毫秒
			for timeMs <= s.TimeMsId {
				timeMs = time.Now().Unix()
			}
		}
	} else {
		s.SequenceId = 0
	}

	s.TimeMsId = timeMs
	id = ((timeMs) << timeShift) | ((s.DataCenterId) << dataCenterShift) | ((s.WorkerId) << workerIdShift) | (s.SequenceId)

	return id, nil
}

func main() {
	s := Snowflake{}
	s.Initial(12, 13)

	var mu sync.Mutex
	m := make(map[int64]bool)

	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()

			for j := 0; j < 5; j++ {
				id, _ := s.NextId()
				mu.Lock()
				if _, ok := m[id]; !ok {
					m[id] = true
				}
				mu.Unlock()
				fmt.Printf("goroutin %d get id %d\n", i, id)
			}
		}(i)
	}

	w.Wait()

	fmt.Printf("m size = %d", len(m))
}
