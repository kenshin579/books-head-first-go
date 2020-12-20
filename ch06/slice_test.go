package slice

import "fmt"

func Example_slicing() {
	//todo : 330
	array3 := [5]string{"a", "b", "c", "d", "e"}
	slice3 := array3[0:3]
	slice3[2] = "X"
	fmt.Println(array3)
	fmt.Println(slice3)

	// Output:
	//[a b X d e]
	//[a b X]

}
