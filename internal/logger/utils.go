package logger

import (
	"bytes"
	"encoding/json"
)

func displayJson(str any) string {

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(str)
	r := buffer.Bytes()
	if err != nil {
		return "ERROR DURING PARSING"
	}
	return string(r)
}
