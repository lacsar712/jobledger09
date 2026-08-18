package jsonx

import (
	"encoding/json"

	"example.com/jobledger/internal/model"
)

func DecodeExport(b []byte) (model.ExportFile, error) {
	var out model.ExportFile
	if err := json.Unmarshal(b, &out); err != nil {
		return model.ExportFile{}, err
	}
	return out, nil
}

func EncodeExport(f model.ExportFile) ([]byte, error) {
	return json.Marshal(f)
}

func DecodeMap(b []byte) (map[string]int, error) {
	var m map[string]int
	_ = json.Unmarshal(b, &m)
	return m, nil
}
