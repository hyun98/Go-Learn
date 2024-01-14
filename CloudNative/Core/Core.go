package Core

import "context"

// Circuit
// DB 또는 다른 서비스와 상호 작용하는 함수의 시그니처
// Context 를 받아서 (string, error) 를 반환하는 Circuit 이라는 람다 함수 형식 정의
type Circuit func(context.Context) (string, error)

type Effector func(context.Context) (string, error)
