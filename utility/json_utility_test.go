package utility_test

import (
	"log"
	"testing"

	. "goweb-scaffold/utility"
)

func TestJsonStringToMap(t *testing.T) {
	testedJsonString := "{\"hash\":\"ww9xPA7Abvwf8CTcih\",\"name\":\"joe\"}"

	result := JsonStringToMap(testedJsonString)
	log.Printf("result: %+v", result)
}

func TestJsonMapToString(t *testing.T) {

	testedJsonMap := map[string]interface{}{
		"name": "leo",
		"age":  10,
	}

	result := JsonMapToString(testedJsonMap)

	log.Printf("result: %+v", result)
}

func TestJsonStringToList(t *testing.T) {
	testedString := "[[\"10.240.0.80\",\"10.240.0.12\"],[\"10.240.0.113\"]]"
	result := [][]string{}

	array := JsonStringToList(testedString)
	for _, subArr := range array {
		group := []string{}
		for _, ele := range subArr.([]interface{}) {
			group = append(group, ele.(string))
		}
		result = append(result, group)
	}
	log.Printf("RtmpGroup: %+v", result)
}

func TestPrettyPrintJsonMap(t *testing.T) {
	testedJsonMap := map[string]interface{}{"apple": 5, "lettuce": 7}

	PrettyPrintJsonMap(testedJsonMap)
}
