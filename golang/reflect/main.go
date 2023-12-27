package main

import (
	"fmt"
	"reflect"
)

type testStruct struct {
	name string
	age  int
}

func main() {
	testObject := testStruct{"bocabbage", 233}
	testType := reflect.TypeOf(testObject)
	testValue := reflect.ValueOf(testObject)
	fmt.Println(testType.String())

	fmt.Println(testValue.String())
	fmt.Printf("%v\n", testValue)
}
