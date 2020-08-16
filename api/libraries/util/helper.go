package util

import (
	jsonIter "github.com/json-iterator/go"
	"strconv"
)

func StrToUint(word string) uint {

	num, _ := strconv.Atoi(word)
	return uint(num)
}

func ObjectToString(obj interface{}) string {
	var json = jsonIter.ConfigCompatibleWithStandardLibrary
	str, _ := json.MarshalToString(obj)
	return str
}

func StringToObject(data string, object interface{}) interface{} {
	var json = jsonIter.ConfigCompatibleWithStandardLibrary
	_ = json.UnmarshalFromString(data, &object)
	return object
}