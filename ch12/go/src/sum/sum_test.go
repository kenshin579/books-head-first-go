package sum

import (
	"fmt"
)

func Example_GetSumFromFile() {
	//path, _ := os.Getwd()
	//fmt.Println("path", path)
	sum := GetSumFromFile("../../bin/data.txt")
	fmt.Println(sum)

	//Output:
	//Opening ../../bin/data.txt
	//Closing file
	//Sum: 50.75
	//50.75
}

func Example_GetSumFromFile_bad_data() {
	//path, _ := os.Getwd()
	//fmt.Println("path", path)
	sum := GetSumFromFile("../../bin/bad-data.txt")
	fmt.Println(sum)

	//Output:
	//Opening ../../bin/data.txt
	//Closing file
	//Sum: 50.75
	//50.75
}
