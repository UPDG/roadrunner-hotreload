# RoadRunner Hot Reload

Allow to automatically reload workers on file change.

## Installation

### QuickBuild

Add to `.build.json` package `github.com/UPDG/roadrunner-hotreload` and register it as `rr.Container.Register(hotreload.ID, &hotreload.Service{})`

After it build RR using QuickBuild.

Example of final file:
```json
{
  "packages": [
    "github.com/spiral/roadrunner/service/env",
    "github.com/spiral/roadrunner/service/http",
    "github.com/spiral/roadrunner/service/rpc",
    "github.com/spiral/roadrunner/service/static",
    "github.com/UPDG/roadrunner-hotreload"
  ],
  "commands": [
    "github.com/spiral/roadrunner/cmd/rr/http"
  ],
  "register": [
    "rr.Container.Register(env.ID, &env.Service{})",
    "rr.Container.Register(rpc.ID, &rpc.Service{})",
    "rr.Container.Register(http.ID, &http.Service{})",
    "rr.Container.Register(static.ID, &static.Service{})",
    "rr.Container.Register(hotreload.ID, &hotreload.Service{})"
  ]
}
```

### Manual

1. Add dependency by running `go get github.com/UPDG/roadrunner-hotreload`

2. Add to `cms/rr/main.go` import `github.com/UPDG/roadrunner-hotreload`

3. Add to `cms/rr/main.go` line `rr.Container.Register(hotreload.ID, &hotreload.Service{})` after `rr.Container.Register(http.ID, &http.Service{})`

Final file should look like this:
```go
package main

import (
	"github.com/sirupsen/logrus"
	rr "github.com/spiral/roadrunner/cmd/rr/cmd"

	// services (plugins)
	"github.com/spiral/roadrunner/service/env"
	"github.com/spiral/roadrunner/service/http"
	"github.com/spiral/roadrunner/service/rpc"
	"github.com/spiral/roadrunner/service/static"
	"github.com/UPDG/roadrunner-hotreload"

	// additional commands and debug handlers
	_ "github.com/spiral/roadrunner/cmd/rr/http"
)

func main() {
	rr.Container.Register(env.ID, &env.Service{})
	rr.Container.Register(rpc.ID, &rpc.Service{})
	rr.Container.Register(http.ID, &http.Service{})
	rr.Container.Register(static.ID, &static.Service{})
	rr.Container.Register(hotreload.ID, &hotreload.Service{})

	rr.Logger.Formatter = &logrus.TextFormatter{ForceColors: true}

	// you can register additional commands using cmd.CLI
	rr.Execute()
}
```

## Configuration

Add your RoadRunner config (`.rr.yaml` by default) this lines:

```yaml
hotreload:
  # Enable or disable plugin
  enable: true
  
  # Path where to scan changed. Current directory by default.
  path: .
  
  # File mask to filter changes. *.php by default. Ref. to https://golang.org/pkg/path/filepath/#Match for available masks.
  files: "*.php"
  
  # Times in milliseconds between file checks. 500 (0.5 sec) by default.
  tick: 500
```
