package helpers

import (
	"encoding/json"
	"log"
)

func JsonParseStruct(s interface{}) string {
	structConv, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(structConv)
}
