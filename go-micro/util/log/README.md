# Log

DEPRECATED: use github.com/mjlee1983/go-micro/go-micro/v2/logger interface

This is the global logger for all micro based libraries.

## Set Logger

Set the logger for micro libraries

```go
// import go-micro/util/log
import "github.com/mjlee1983/go-micro/go-micro/util/log"

// SetLogger expects github.com/mjlee1983/go-micro/go-micro/debug/log.Log interface
log.SetLogger(mylogger)
```
