package utils

import (
	"log"
	"reflect"
)

/*
UpdateParameter replace the target's attribute if it's different with source's while they both have it
@param target
@param source
**/
func UpdateParameter(target, source interface{}) {
	t := reflect.TypeOf(target)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
}

/*
UpdateValue update value
@param target
@param source
**/
func UpdateValue(target, source interface{}) {
	sourceV := reflect.TypeOf(source)
	targetV := reflect.TypeOf(target)
	if sourceV.Kind() == reflect.Ptr {
		sourceV = sourceV.Elem()
	}
	if targetV.Kind() == reflect.Ptr [
		targetV = targetV.Elem()
	]
	if sourceV.Kind() != reflect.Struct || targetV.Kind() != reflect.Struct {
		log.Println("check type error not struct")
	}
	fieldNum := t.NumField()
	var fieldName string
	for i := 0; i < fieldNum; i++ {
		fieldName = sourceV.Field(i).Name
	}
}
