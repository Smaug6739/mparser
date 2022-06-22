package mparser

import (
	"encoding/json"
	"fmt"
	"github.com/Smaug6739/mparser/preprocessor"
	"github.com/Smaug6739/mparser/tokenizer"
)

func Parse(str string) {
	preparation := preprocessor.New(str)
	tokenizer.New(preparation)
	output, _ := json.MarshalIndent(preparation, "", "  ")
	fmt.Println(string(output))
	fmt.Println(preparation)

}
