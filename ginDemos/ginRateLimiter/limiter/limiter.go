package limiter

import (
	"sync"
	"time"
)

type Limiter struct {
	tb *ToeknBucket
}

type ToeknBucket struct {
	// 多协程访问，加锁保护
	mu sync.Mutex
	// 桶大小
	size int
	// 当前桶剩余 token 数量
	count int

	// 填充速率，即每隔多久补充 addNum 个 token
	rateLimit time.Duration

	// 每次补充桶的数量
	addNum int

	// 最后请求时间
	lastRequestTime time.Time
}

func (tb *ToeknBucket)fillToken(){
	tb.count+=tb.getFillTokenCount()
}

func (tb *ToeknBucket) getFillTokenCount() int{
	//  满桶不需要补充
	if tb.count >= tb.size{
		return 0
	}
	if tb.lastRequestTime.IsZero(){
		return 0
	}
	duration := time.Since(tb.lastRequestTime)
	count := int(duration/tb.rateLimit) * tb.addNum
	remainNum := tb.size - tb.count
	if remainNum>=count{
		return count
	}else{
		return  remainNum
	}
}
func (tb *ToeknBucket) allow() bool{
	// 填充
	tb.fillToken()
	if tb.count >0 {
		tb.count--
		tb.lastRequestTime = time.Now()
		return true
	}
	return false
}

func NewLimiter(r time.Duration, size int, addNum int) *Limiter {
	return &Limiter{
		tb: &ToeknBucket{
			rateLimit: r,
			size:      size,
			count:     size,
			addNum:    addNum,
		},
	}
}

func (l *Limiter) Allow() bool {
	l.tb.mu.Lock()
	defer l.tb.mu.Unlock()
	return l.tb.allow()
}
