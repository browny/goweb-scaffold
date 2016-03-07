package utility

import (
	"encoding/json"
	log "github.com/cihub/seelog"
)

// JsonStringToMap converts json string to map
func JsonStringToMap(jsonString string) map[string]interface{} {
	var dat map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &dat)
	if err != nil {
		log.Debugf("Unmarshal error: jsonString[%s], error[%s]", jsonString, err.Error())
	}

	return dat
}

// JsonMapToString converts json map to string
func JsonMapToString(jsonMap map[string]interface{}) []byte {
	dat, err := json.Marshal(jsonMap)
	if err != nil {
		log.Warn(err.Error())
	}

	return dat
}

// PrettyPrintJsonMap print json map with indent
func PrettyPrintJsonMap(jsonMap map[string]interface{}) {
	result, _ := json.MarshalIndent(jsonMap, "", "   ")

	log.Trace(string(result))
}

// JsonStringToList converts json string to list
func JsonStringToList(jsonString string) []interface{} {
	var dat []interface{}

	err := json.Unmarshal([]byte(jsonString), &dat)
	if err != nil {
		log.Debugf("Unmarshal error: jsonString[%s], error[%s]", jsonString, err.Error())
	}

	return dat
}
