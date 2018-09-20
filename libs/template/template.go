package template

import (
	"html/template"
	"runtime"
	"strings"
)

func NewFuncMap() []template.FuncMap {
	return []template.FuncMap{map[string]interface{}{
		"GoVer": func() string {
			return strings.Title(runtime.Version())
		},
		"ToTF": func(value interface{}) string {
			switch value.(type) {
			case int:
				if value.(int) > 0 {
					return "是"
				}
			case bool:
				if value.(bool) {
					return "是"
				}
			}
			return ""
		},
	}}
}
