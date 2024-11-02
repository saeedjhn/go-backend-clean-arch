package prettyprint

import (
	"encoding/json"
	"fmt"
)

func PrettyPrintData(data interface{}) {
	// Convert data to pretty-printed JSON.
	if prettyOutput, err := json.MarshalIndent(data, "", "  "); err == nil {
		fmt.Println(string(prettyOutput))
	} else {
		// Handle error
		fmt.Println("*****")
	}
}
