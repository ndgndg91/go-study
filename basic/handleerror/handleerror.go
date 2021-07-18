package handleerror

import "fmt"

type CanNotBeZero struct {
	a int
	b int
}

func (e *CanNotBeZero) Error() string {
	return fmt.Sprintf("나눗셈 시도 중 %d / %d 제수가 0 값이 들어왔습니다.", e.a, e.b)
}

func newCanNotBeZeroErr(a, b int) *CanNotBeZero {
	return &CanNotBeZero{a, b}
}

func Drive() {
	f()
	fmt.Println("계속 실행됩니다.")
}

func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		if r, ok := recover().(*CanNotBeZero); ok {
			fmt.Println("panic 복구 - ", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
	fmt.Printf("9 / 3 = %d\n", h(10, 2))
}

func h(a, b int) int {
	if b == 0 {
		panic(newCanNotBeZeroErr(a, b))
	}

	return a / b
}
