package utils

import (
	"encoding/json"
	"strconv"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/mumushuiding/util"
	"github.com/pkg/errors"
)

func ParseAnyToData(data *any.Any) interface{} {
	var result interface{}
	err := json.Unmarshal(data.Value, &result)
	if err != nil {
		panic(errors.Wrap(err, "解析Any失败"))
	}
	return result
}

func ParseAnyToDataList(anyDataList []*any.Any) []interface{} {
	var dataList []interface{}
	for _, anyData := range anyDataList {
		data := ParseAnyToData(anyData)
		dataList = append(dataList, data)
	}
	return dataList
}

func ParseAnyToMap(data *any.Any) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal(data.Value, &result)
	if err != nil {
		panic(errors.Wrap(err, "解析Any失败"))
	}
	return result
}

func ParseAnyToMapBool(data *any.Any) map[string]bool {
	var result map[string]bool
	err := json.Unmarshal(data.Value, &result)
	if err != nil {
		panic(errors.Wrap(err, "解析Any失败"))
	}
	return result
}

func ParseAnyToMapIntFloatList(data *any.Any) map[int][]float64 {
	var result map[int][]float64
	err := json.Unmarshal(data.Value, &result)
	if err != nil {
		panic(errors.Wrap(err, "解析Any失败"))
	}
	return result
}

func ParseAnyToMapInt(data *any.Any) map[int]map[string]interface{} {
	var result map[int]map[string]interface{}
	err := json.Unmarshal(data.Value, &result)
	if err != nil {
		panic(errors.Wrap(err, "解析Any失败"))
	}
	return result
}

func ParseAnyToMapIntList(data *any.Any) map[int][]map[string]interface{} {
	var result map[int][]map[string]interface{}
	err := json.Unmarshal(data.Value, &result)
	if err != nil {
		panic(errors.Wrap(err, "解析Any失败"))
	}
	return result
}

func ParseAnyToMapList(anyDataList []*any.Any) []map[string]interface{} {
	var mapList []map[string]interface{}
	for _, anyData := range anyDataList {
		map1 := ParseAnyToMap(anyData)
		mapList = append(mapList, map1)
	}
	return mapList
}

func ParseMapToAny(data map[string]interface{}) (*any.Any, error) {
	jsonStr, err := util.ToJSONStr(data)
	if err != nil {
		return nil, err
	}
	return &any.Any{Value: []byte(jsonStr)}, nil
}

func ParseMapIntToAny(data map[int]map[string]interface{}) (*any.Any, error) {
	jsonStr, err := util.ToJSONStr(data)
	if err != nil {
		return nil, err
	}
	return &any.Any{Value: []byte(jsonStr)}, nil
}

func ParseMapListToAnyList(dataList []map[string]interface{}) ([]*any.Any, error) {
	var anyList []*any.Any
	for _, data := range dataList {
		anyData, err := ParseMapToAny(data)
		if err != nil {
			return nil, err
		}
		anyList = append(anyList, anyData)
	}
	return anyList, nil
}

func ParseDataToAny(data interface{}) (*any.Any, error) {
	jsonStr, err := util.ToJSONStr(data)
	if err != nil {
		return nil, err
	}
	return &any.Any{Value: []byte(jsonStr)}, nil
}

func ParseDataListToAnyList(dataList []interface{}) ([]*any.Any, error) {
	var anyList []*any.Any
	for _, data := range dataList {
		anyData, err := ParseDataToAny(data)
		if err != nil {
			return nil, err
		}
		anyList = append(anyList, anyData)
	}
	return anyList, nil
}

func ParseJsonStrToAnyList(jsonStr string) ([]*any.Any, error) {
	var anyList []*any.Any
	var dataList []interface{}
	err := json.Unmarshal([]byte(jsonStr), &dataList)
	if err != nil {
		return nil, err
	}
	for _, data := range dataList {
		anyData, err := ParseDataToAny(data)
		if err != nil {
			return nil, err
		}
		anyList = append(anyList, anyData)
	}
	return anyList, nil
}

func ReverseMapArray(arr []map[string]interface{}) []map[string]interface{} {
	var result []map[string]interface{}
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}
