package log_format

import (
	"testing"
)

var e4f = NewE4Log()

func TestE4Log_ExampleLine(t *testing.T) {
	println(e4f.ExampleLine())
}

func TestE4Log_GrokParse(t *testing.T) {
	line := "{\"timestamp\":\"2022-08-31T14:36:57+08:00\",\"host\":\"8.8.9.8\",\"path\":\"/var/log\",\"application\":\"APPLICATION\",\"component\":\"COMPONENT\",\"log_level\":\"DEBUG\",\"thread_name\":\"THREAD_NAME\",\"extend\":{\"key\":\"value\"},\"msg\":\"This is Message\"}"
	e4log, err := ParseJSON(line)
	if err != nil {
		panic(err)
	}
	println(e4log.JSon())
}
