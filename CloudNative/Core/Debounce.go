package Core

import (
	"context"
	"sync"
	"time"
)

// DebounceFirst
// time.Duration 은 1 nanosecond 를 의미한다.
// Thread-Safety 를 보장하기 위해 함수 전체를 뮤텍스로 감싼다.
// 정해진 시간 d 이내에 함수가 다시 호출될 경우 마지막으로 호출된 시간을 추적해 캐시된 결과를 반환한다.
func DebounceFirst(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time
	var result string
	var err error
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()

		defer func() {
			// 함수 호출될 때마다 초기화되도록 한다.
			threshold = time.Now().Add(d)
			m.Unlock()
		}()

		// 마지막으로 호출된 시간 확인
		if time.Now().Before(threshold) {
			return result, err
		}

		result, err = circuit(ctx)

		return result, err
	}
}

// DebounceLast
// 함수가 호출된 이후 충분한 시간이 흘렀는지를 확인하기 위해 time.Ticker 를 사용한다.
func DebounceLast(circuit Circuit, d time.Duration) Circuit {
	threshold := time.Now()
	var ticker *time.Ticker
	var result string
	var err error
	var once sync.Once
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()

		defer m.Unlock()

		// 호출될 때마다 시간 갱신
		threshold = time.Now().Add(d)

		// circuit 호출 로직은 단 한번만 발생한다.
		once.Do(func() {
			ticker = time.NewTicker(time.Millisecond * 100)

			go func() {
				defer func() {
					m.Lock()
					ticker.Stop()
					once = sync.Once{}
					m.Unlock()
				}()

				for {
					select {
					case <-ticker.C:
						m.Lock()
						// tick 만큼 시간이 흐르는 동안 debounce 호출이 없었다면 circuit 서비스를 호출한다.
						if time.Now().After(threshold) {
							result, err = circuit(ctx)
							m.Unlock()
							return
						}
					case <-ctx.Done():
						m.Lock()
						result, err = "", ctx.Err()
						m.Unlock()
						return
					}

				}
			}()
		})

		return result, err
	}
}
