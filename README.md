[![](https://godoc.org/github.com/itzg/zapconfigs?status.svg)](https://godoc.org/github.com/itzg/zapconfigs)
[![](https://img.shields.io/badge/go.dev-module-007D9C)](https://pkg.go.dev/github.com/itzg/zapconfigs)

Provides a few more opinionated zap logger configurations beyond the ones provided by [zap](https://godoc.org/go.uber.org/zap)'s package.

## Example

```go

import (
	"github.com/itzg/zapconfigs"
)

func main() {
    // ...

    var logger *zap.Logger
    if debug {
        logger = zapconfigs.NewDebugLogger()
    } else {
        logger = zapconfigs.NewDefaultLogger()
    }
    defer logger.Sync()

}
```