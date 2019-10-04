package main

import (
	"errors"
	"fmt"
)

func Divide(num int, divider int) (int, error) {
	if divider == 0 {
		return num, errors.New(fmt.Sprintf("%s %d %s", "can't divide", num, "by 0"))
	}
	return num / divider, nil
}

type customError struct {
	arg int
	err string
}

func (custErr *customError) Error() string {
	return fmt.Sprintf("%d - %s", custErr.arg, custErr.err)
}

func DivideWithCustomError(num int, divider int) (int, error) {
	if divider == 0 {
		return num, &customError{num, fmt.Sprintf("%s %d %s", "cant' divide", num, "by 0")}
	}
	return num / divider, nil
}

func testErrors() {
	num := 30
	for _, data := range []int{2, 3, 5, 0} {
		if ret, err := Divide(num, data); err != nil {
			fmt.Println("Divide function failed:", err)
		} else {
			fmt.Printf("Divide function worked for %d by %d = %d\n", num, data, ret)
		}
	}
	num = 60
	for _, data := range []int{2, 3, 5, 0} {
		if ret, err := DivideWithCustomError(num, data); err != nil {
			fmt.Println("DivideWithCustomError function failed:", err)
		} else {
			fmt.Printf("DivideWithCustomError function worked for %d by %d = %d\n", num, data, ret)
		}
	}

	// get error data back from function DivideWithCustomError to use programmatically
	_, err := DivideWithCustomError(60, 0)
	if custErr, boolVal := err.(*customError); boolVal {
		fmt.Print(custErr.arg, " ")
		fmt.Println(custErr.err)
	}

	_, err1 := DivideWithCustomError(60, 2)
	custErr, boolVal := err1.(*customError)
	fmt.Print(custErr, " ")
	fmt.Println(boolVal)
}
