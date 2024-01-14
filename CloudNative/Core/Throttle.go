package Core

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Throttle
// 에러 반환 전략을 사용하는 기본적인 토큰 버킷 알고리즘을 사용했다.
func Throttle(effector Effector, max uint, refill uint, d time.Duration) Effector {
	var tokens = max
	var once sync.Once
	return func(ctx context.Context) (string, error) {
		if ctx.Err() != nil {
			return "", ctx.Err()
		}

		// 일정 시간마다 토큰을 리필하는 로직은 스로틀 별로 한 번만 실행한다.
		once.Do(func() {
			ticker := time.NewTicker(d)

			go func() {
				defer ticker.Stop()

				for {
					select {
					case <-ctx.Done():
						return
					case <-ticker.C: // 주기 d 마다 토큰을 리필한다.
						t := tokens + refill
						if t > max {
							t = max
						}
						tokens = t
					}
				}
			}()
		})

		if tokens <= 0 {
			return "", fmt.Errorf("too many calls")
		}

		tokens--

		return effector(ctx)
	}
}
