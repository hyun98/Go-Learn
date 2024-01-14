package Core

import (
	"context"
	"log"
	"time"
)

// Retry
// Effector 시그니처를 가진 함수를 반환하지만 함수가 외부 상태를 갖지 않는다.
// 따라서 동시성 문제를 해결하기 위해 뮤텍스 같은 기능을 사용할 필요가 없다.
func Retry(effector Effector, retries int, delay time.Duration) Effector {
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			response, err := effector(ctx)

			// 서비스 호출이 성공했거나 재시도 제한 횟수에 도달했으면 응답을 반환한다.
			if err == nil || r >= retries {
				return response, err
			}

			log.Printf("Attempt %d failed; retrying in %v", r+1, delay)

			// case에 걸릴 때까지 blocking 되는 특징이 있다.
			select {
			// delay 시간이 흐른 후 for문을 다시 순회한다.
			case <-time.After(delay):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}
