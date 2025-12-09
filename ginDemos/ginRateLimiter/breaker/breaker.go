package breaker

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	// 断路器关闭状态
	STATE_CLOSE = iota
	// 断路器开启状态
	STATE_OPEN
	// 断路器限制状态
	STATE_HALF_OPEN
)

type Breaker struct {
	mu               sync.Mutex
	state            int
	failureThreshold int
	successThreshold int
	// 限制周期里最多请求次数
	halfMaxRequest int
	// 限制周期里请求次数
	halfCycleReqCount int
	// 既是正常情况下的交替周期又是断开状态下的时间周期
	// 也可以定义两个变量
	timeout time.Duration
	// 连续失败次数
	failureCount int
	// 连续成功次数
	successCount int
	// 周期开始时间
	cycleStartTime time.Time
}

func NewBreaker(failureThreshold int,
	successThreshold int, halfMaxRequest int, timeout time.Duration) *Breaker {
	return &Breaker{
		state:            STATE_CLOSE,
		failureThreshold: failureThreshold,
		successThreshold: successThreshold,
		halfMaxRequest:   halfMaxRequest,
		timeout:          timeout,
	}
}

func (b *Breaker) Exec(f func() error) error {
	b.before()
	// 经过 before 后断路器的状态得到清算，可以判断断路器的状态
	if b.state == STATE_OPEN {
		return errors.New("断路器处于打开状态，无法访问服务")
	}

	fmt.Println("b.state ",b.state)
	// if b.state == STATE_CLOSE {
	// 	err := f()
	// 	b.after(err)
	// 	return err
	// }

	if b.state == STATE_HALF_OPEN {
		if b.halfCycleReqCount >= b.halfMaxRequest {
			return errors.New("断路器处于半开启状态，单位时间内请求超过次数")
		}

		// err := f()
		// b.after(err)
		// return err
	}

	err := f()
	b.after(err)
	return err

	// return nil
}

func (b *Breaker) before() {
	b.mu.Lock()
	defer b.mu.Unlock()
	switch b.state {

	case STATE_OPEN:
		// 开启一个新的周期
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			b.state = STATE_HALF_OPEN
			b.reset()
			return
		}
	case STATE_HALF_OPEN:
		// 成功数量超过阈值则关闭断路器
		if b.successCount >= b.successThreshold {
			b.state = STATE_CLOSE
			b.reset()
			return
		}
		// 超过一个周期
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			b.cycleStartTime = time.Now()
			b.halfCycleReqCount = 0
			return
		}
	case STATE_CLOSE:
		// 断路器关闭状态超过一个周期则进入半关闭状态
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			fmt.Println("aaaaaa")
			b.reset()
			return
		}
	}
}
func (b *Breaker) after(err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if err == nil {
		b.onSuccess()
	} else {
		b.onFailure()
	}
}

func (b *Breaker) onSuccess() {
	b.failureCount = 0
	if b.state == STATE_HALF_OPEN {
		b.successCount++
		b.halfCycleReqCount++
		if b.successCount >= b.successThreshold {
			b.state = STATE_CLOSE
			b.reset()
		}
	}
}

func (b *Breaker) onFailure() {
	b.successCount = 0
	b.failureCount++
	if b.state == STATE_HALF_OPEN || (b.state == STATE_CLOSE && b.failureCount >= b.failureThreshold) {
		b.state = STATE_OPEN
		b.reset()
		return
	}
}

func (b *Breaker) reset() {
	b.successCount = 0
	b.failureCount = 0
	b.halfCycleReqCount = 0
	b.cycleStartTime = time.Now()
}
