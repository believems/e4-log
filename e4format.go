package e4_log

import (
	"bytes"
	"encoding/json"
	"time"
)

type E4Log struct {
	Timestamp   time.Time `json:"timestamp"`   //时间戳
	Host        string    `json:"host"`        //主机地址
	Application string    `json:"application"` //组件
	Component   string    `json:"component"`   //模块
	LogLevel    string    `json:"log_level"`   //日志级别
	ThreadName  string    `json:"thread_name"` //进程名称
	Extend      E4Extend  `json:"extend"`      //扩展属性
	Msg         string    `json:"msg"`         //消息体
}

var exampleE4log = E4Log{
	Timestamp:   time.Now(),
	Host:        "8.8.9.8",
	Application: "APPLICATION",
	Component:   "COMPONENT",
	LogLevel:    "DEBUG",
	ThreadName:  "THREAD_NAME",
	Extend:      E4Extend{"key": "value"},
	Msg:         "This is Message",
}

var E4LogFields = []string{"timestamp", "host", "application", "component", "log_level", "thread_name", "extend", "msg"}

func CheckFields(fields ...string) bool {
	for _, field := range fields {
		if !Contains(E4LogFields, field) {
			return false
		}
	}
	return true
}

func CheckMapKey(checkedMap map[string]interface{}) bool {
	keys := make([]string, 0, len(checkedMap))
	for k := range checkedMap {
		keys = append(keys, k)
	}
	return CheckFields(keys...)
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func NewE4Log() *E4Log {
	return &E4Log{}
}

func ParseJSON(jsonStr string) (*E4Log, error) {
	data := &E4Log{}
	err := json.Unmarshal([]byte(jsonStr), data)
	return data, err
}

func (e4log *E4Log) ExampleLine() string {
	return exampleE4log.String()
}

func (e4log *E4Log) StringMap() (map[string]interface{}, error) {
	var stringMap map[string]interface{}
	data, err := json.Marshal(e4log)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &stringMap)
	if err != nil {
		return nil, err
	}
	return stringMap, nil
}

func (e4log *E4Log) MarshalJSON() ([]byte, error) {
	type Alias E4Log
	return json.Marshal(&struct {
		Timestamp string `json:"timestamp"`
		*Alias
	}{
		Timestamp: e4log.Timestamp.Format(time.RFC3339),
		Alias:     (*Alias)(e4log),
	})
}

func (e4log *E4Log) JSon() string {
	data, err := json.Marshal(e4log)
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		return ""
	} else {
		return out.String()
	}
}

func (e4log *E4Log) String() string {
	data, err := json.Marshal(e4log)
	if err != nil {
		return ""
	} else {
		var compact bytes.Buffer
		err = json.Compact(&compact, data)
		if err != nil {
			return ""
		}
		return compact.String()
	}
}
