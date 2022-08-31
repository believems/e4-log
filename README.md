- 输出的格式: `"timestamp", "host", "path", "application", "component", "log_level", "thread_name", "extend", "msg"`
- JSON格式
```json
{
  "timestamp": "2022-08-31T14:36:57+08:00",
  "host": "8.8.9.8",
  "path": "/var/log",
  "application": "APPLICATION",
  "component": "COMPONENT",
  "log_level": "DEBUG",
  "thread_name": "THREAD_NAME",
  "extend": {
    "key": "value"
  },
  "msg": "This is Message"
}
```
```go
require (
	github.com/believems/e4-log v0.0.2
)
```