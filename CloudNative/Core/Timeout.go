package Core

import "context"

// SlowFunction
// 일정 시간 내에 동작을 마칠 수도 있고 아닐 수 도 있는 가상의 함수 시그니처
type SlowFunction func(string) (string, error)

type WithContext func(context.Context, string) (string, error)

// Timeout
// SlowFunction 을 직접 호출하는 대신 SlowFunction 을 클로저로 감싸고 WithContext 함수를 반환한다.
func Timeout(function SlowFunction) WithContext {
	return func(ctx context.Context, arg string) (string, error) {
		chres := make(chan string)
		cherr := make(chan error)

		// 오래 걸리는 함수의 실행을 고루틴에서 실행시키고, 고채널을 통해서 결과를 받아온다.
		go func() {
			res, err := function(arg)
			chres <- res
			cherr <- err
		}()

		select {
		case res := <-chres: // 채널에 응답이 적재되면 값을 반환한다.
			return res, <-cherr
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}
