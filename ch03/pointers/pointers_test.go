package pointers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

/*
 1.POINTERS
 & 를 사용해서 변수의 주소를 얻어올 수 있다.
*/
func TestPointers(t *testing.T) {
	amount := 6
	fmt.Println(amount)  //6
	fmt.Println(&amount) //주소 값
}

func TestPointers2(t *testing.T) {
	var myInt int
	var myFloat float64
	var myBool bool

	fmt.Println(&myInt)
	fmt.Println(&myFloat)
	fmt.Println(&myBool)
}

func TestFoo(t *testing.T) {
	// todo test code
	expected := 1
	actual := 0
	assert.Equal(t, expected, actual, "기대값과 결과값이 다릅니다.")
}

/*
2. Pointer types
ex. *int (pointer to int로 읽고 이건 포인터 타입을 선언할 때 사용함)
*/
func Example_pointerTypes() {
	var myInt int
	var myFloat float64
	var myBool bool

	fmt.Println(reflect.TypeOf(&myInt))
	fmt.Println(reflect.TypeOf(&myFloat))
	fmt.Println(reflect.TypeOf(&myBool))
	// Output:
	//*int
	//*float64
	//*bool
}

func TestPointerTypes2(t *testing.T) {
	var myInt int
	var myIntPointer *int //myIntPointer 이 변수는 포인터 타입을 담을 수 있다
	myIntPointer = &myInt
	fmt.Println(myIntPointer)

	var myFloat float64
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)

	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
}

/*
3.Getting or changing the value at a pointer
*/

func TestPointerTypes3(t *testing.T) {
	myInt := 4
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer) //4

	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)
}

func TestPointerTypes4(t *testing.T) {
	//* operator를 통해서 pointer의 값을 변경할 수도 있다
	myInt := 4
	fmt.Println(myInt)
	myIntPointer := &myInt
	*myIntPointer = 8 //*로 값을 변경할수 있다.
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}
