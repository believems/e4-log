package log_format

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
		json.Compact(&compact, data)
		return compact.String(), nil
	}
}

func (extend *E4Extend) String() (string, error) {
	return extend.JsonString()
}
