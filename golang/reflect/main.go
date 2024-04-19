package main

import (
	"fmt"
	"reflect"
)

type testStruct struct {
	name       string `tag_1:"name_tag1" tag_2:"name_tag2"`
	age        int    `tag_1:"age_tag1"`
	notagField int
}

func BasicUse() {
	testObject := testStruct{"bocabbage", 233, 1}
	testType := reflect.TypeOf(testObject)
	testValue := reflect.ValueOf(testObject)
	fmt.Println(testType.String())

	fmt.Println(testValue.String())
	fmt.Printf("%v\n", testValue)

	fmt.Printf("%v\n", reflect.TypeOf(3))
	fmt.Printf("%v\n", reflect.ValueOf(3).Int())
}

func checkCanCall(v reflect.Value) {
	fmt.Printf("check %v\n", v)
}

func StructTagCheck() {
	testObject := testStruct{"bocabbage", 233, 1}
	typeOfObj := reflect.TypeOf(testObject)
	for i := 0; i < typeOfObj.NumField(); i++ {
		// get field
		field := typeOfObj.Field(i)
		// check value
		checkCanCall(reflect.ValueOf(testObject).Field(i))
		// check specific tag
		fmt.Printf("Tag value of tag_1: %s\n", field.Tag.Get("tag_1"))
		fmt.Printf("Tag value of tag_2: %s\n", field.Tag.Get("tag_2"))
	}
}

func main() {
	BasicUse()
	StructTagCheck()
}
