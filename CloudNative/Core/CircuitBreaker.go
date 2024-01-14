package Core

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Breaker
// 클로저 개념이 적용되었다 - 필요한 기능을 제공하기 위해 circuit 을 감싼 또 다른 함수를 만들었다.
// circuit : Circuit 타입 정의를 따르는 어떤 함수라도 사용 가능하다.
// failureThreshold : 서킷을 자동으로 open 상태로 바꾸기 전까지 허용하는 실패 횟수
func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	consecutiveFailures := 0
	lastAttempt := time.Now()
	var mutex sync.RWMutex

	return func(ctx context.Context) (string, error) {
		mutex.RLock() // Read Lock

		d := consecutiveFailures - int(failureThreshold)

		if d >= 0 {
			// failureThreshold 만큼 연속 실패 후 재시도 시간 간격을 2배씩 늘린다.
			allowRetryTime := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(allowRetryTime) {
				mutex.RUnlock()
				return "", errors.New("service unreachable")
			}
		}

		mutex.RUnlock() // read unlock

		response, err := circuit(ctx)

		mutex.Lock() // lock
		defer mutex.Unlock()

		lastAttempt = time.Now() // 요청 시간 기록

		if err != nil { // circuit이 에러를 반환하는 경우
			consecutiveFailures++ // 연속 실패 횟수를 증가시키고
			return response, err  // 해당 응답과 에러를 반환
		}

		consecutiveFailures = 0 // 실패 횟수 초기화

		return response, nil
	}
}
