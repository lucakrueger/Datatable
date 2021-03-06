package core

import (
	"encoding/json"
	"fmt"
)

func (this *Listener) GenerateJSON() []byte {
	var types Types
	types.Types = this.Types

	file, err := json.MarshalIndent(types, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func ReadJSON(bytes []byte) Types {
	var types Types
	err := json.Unmarshal(bytes, types)
	if err != nil {
		fmt.Println(err)
	}
	return types
}
