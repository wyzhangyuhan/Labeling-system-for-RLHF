package build_modelstring

import (
	"fmt"
	"label_system/global/consts"
)

func BuildString(query, answer []string) string {
	var modelstring string = ""
	for idx, v := range query {
		if idx == 0 {
			modelstring = modelstring + consts.UserChar
		} else {
			modelstring = modelstring + consts.SepChar + consts.UserChar
		}
		modelstring = modelstring + v
		if idx < len(answer) {
			modelstring = modelstring + consts.SepChar + consts.BotChar
			modelstring = modelstring + answer[idx]
		}
	}
	fmt.Printf("modelString: %v\n", modelstring)
	return modelstring
}

func BuildStringNeox(query, answer []string) string {
	var modelstring string = ""
	for idx, v := range query {
		modelstring = modelstring + "<human>：" + v
		if idx < len(answer) {
			modelstring = modelstring + "<bot>："
			modelstring = modelstring + answer[idx]
		}
	}
	modelstring += "<bot>："
	fmt.Printf("modelString: %v\n", modelstring)
	return modelstring
}
