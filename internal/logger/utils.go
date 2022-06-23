package logger

import "encoding/json"

func displayJson(str any) string {
	r, err := json.MarshalIndent(str, "", "  ")
	if err != nil {
		return "ERROR DURING PARSING"
	}
	return string(r)
}
