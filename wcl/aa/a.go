package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

var jsonf []byte = []byte(`{"data":{"reportData":{"report": {"f1":{"data":[{"timestamp":1089287,"type":"cast","sourceID":56,"targetID":11,"abilityGameID":339690}]},"f2":{"data":[{"timestamp":1089287,"type":"cast","sourceID":56,"targetID":11,"abilityGameID":339690}]}}}}}`)

type sut struct {
	Data struct {
		ReportData struct {
			Report interface{} `Json:"report"`
		} `Json:"reportData"`
	} `Json:"data"`
}

type inside struct {
	Data []struct {
		Timestamp     float64 `Json:"timestamp"`
		Type          string  `Json:"type"`
		SourceID      int     `Json:"sourceID"`
		TargetID      int     `Json:"targetID"`
		AbilityGameID int     `Json:"abilityGameID"`
		Fight         int     `Json:"fight"`
	} `Json:"data"`
}

func main() {
	var v sut

	json.Unmarshal(jsonf, &v)

	fmt.Println(v.Data.ReportData.Report)
	for _, elt := range v.Data.ReportData.Report.(map[string]interface{}) {
		var result inside
		mapstructure.Decode(elt, &result)
		fmt.Println(result.Data[0].AbilityGameID)
	}
}
