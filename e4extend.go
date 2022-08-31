package e4_log

import (
	"bytes"
	"encoding/json"
)

type E4Extend map[string]interface{}

func (extend *E4Extend) JsonString() (string, error) {
	data, err := json.Marshal(extend)
	if err != nil {
		return "", err
	} else {
		var compact bytes.Buffer
		err := json.Compact(&compact, data)
		if err != nil {
			return "", err
		}
		return compact.String(), nil
	}
}

func (extend *E4Extend) String() (string, error) {
	return extend.JsonString()
}
