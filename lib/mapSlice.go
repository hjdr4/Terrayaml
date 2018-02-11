package lib

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hjdr4/yaml"
)

//We alias the yaml.MapSlice type in order to implement MarshalJSON() interface
type MapSlice yaml.MapSlice

func (m MapSlice) MarshalJSON() ([]byte, error) {
	ret := ""
	ret += "{"
	var inner []string
	for _, v := range m {
		innerStr := ""
		ks := v.Key.(string)
		innerStr += "\"" + ks + "\":"

		switch v.Value.(type) {
		case MapSlice:
			val := v.Value.(MapSlice)
			res, err := json.Marshal(val)
			if err != nil {
				return nil, err
			}
			innerStr += string(res)

		case string:
			j, err := json.Marshal(v.Value)
			if err != nil {
				return nil, err
			}
			innerStr += string(j)

		case []interface{}:
			j, err := json.Marshal(v.Value)
			if err != nil {
				return nil, err
			}
			innerStr += string(j)

		case map[interface{}]interface{}:
			use := Convert(v.Value)
			j, err := json.Marshal(use)
			if err != nil {
				return nil, err
			}
			innerStr += string(j)

		default:
			innerStr += fmt.Sprintf("%v", v.Value)
		}
		inner = append(inner, innerStr)
	}
	ret += strings.Join(inner, ",")
	ret += "}"
	return []byte(ret), nil
}
