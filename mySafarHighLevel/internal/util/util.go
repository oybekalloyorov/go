package util

import (
	"encoding/json"
	"fmt"
)

func PrintPrettyJSON(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}


