package mparser

import (
	"encoding/json"
	"fmt"
	"github.com/Smaug6739/mparser/preprocessor"
)

func Parse(str string) {
	preparation := preprocessor.New(str)
	output, _ := json.MarshalIndent(preparation, "", "  ")
	fmt.Println(string(output))
}
