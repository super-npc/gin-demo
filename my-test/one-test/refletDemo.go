package one_test

import (
	"fmt"
	"reflect"
)

type RefletDemo struct {
	Name string `my:"dfd"`
}

func NewRefletDemo() {
	demo := RefletDemo{Name: "demo111"}
	of := reflect.TypeOf(demo)
	fmt.Printf("类型: %s, 种类: %s", of.Name(), of.Kind())

	for i := range of.NumField() {
		field := of.Field(i)
		fmt.Printf("name: %v, type: %v, tag: %v", field.Name, field.Type, field.Tag.Get("my"))
	}
}

func ChangeToObj() {
	jsonStr := `{"Name":"demo11222","Age":18,"On":true}`
	registerByStruct(&RefletDemo{})
	obj, err := NewStructFromJSONAndName("RefletDemo", []byte(jsonStr))
	if err != nil {
		panic(err)
	}

	fmt.Printf("返回对象类型：%T\n", obj)
	fmt.Printf("返回对象值  ：%+v\n", obj)
}
