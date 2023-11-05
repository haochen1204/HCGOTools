package Group

import (
	"sync"
	"sync/atomic"
	"time"
)

type WaitGroupCount struct {
	sync.WaitGroup
	NowCount int64
	MaxCount int64
}

func NewWaitGroupCount(count *int64) WaitGroupCount {
	var NowCount = 1
	var MaxCount = atomic.LoadInt64(count)
	return WaitGroupCount{NowCount: int64(NowCount), MaxCount: MaxCount}
}

func (wg *WaitGroupCount) Add(delta int) {
	atomic.AddInt64(&wg.NowCount, int64(delta))
	wg.WaitGroup.Add(delta)
}

func (wg *WaitGroupCount) WaitThread() bool {
	for {
		if int64(wg.GetCount()) >= wg.MaxCount {
			//fmt.Println("等待中...", wg.NowCount, wg.MaxCount)
			time.Sleep(1)
		} else {
			return true
		}
	}
}

func (wg *WaitGroupCount) Done() {
	atomic.AddInt64(&wg.NowCount, -1)
	wg.WaitGroup.Done()
}

func (wg *WaitGroupCount) GetCount() int {
	return int(atomic.LoadInt64(&wg.NowCount))
}
