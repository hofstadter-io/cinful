# cinful

Golang library and CLI to detect if a program
is running in a CI system.

Pronounced "sinful" by mashing together CI, info, and the 'ful' suffix.

Originally ported from [watson/ci-info](https://github.com/watson/ci-info)
and [other links](https://adamj.eu/tech/2020/03/09/detect-if-your-tests-are-running-on-ci/)


### as a CLI

prints details if in CI, otherwise nothing

```sh
cinful

cinful list # show all
```

Binaries are available on the releases page.


### as a pkg

```go
package main

import (
	"fmt"

	"github.com/hofstadter-io/cinful"
)

func main() {
	vendor := cinful.Info()
	if vendor != nil {
		fmt.Println(vendor)
	}
}
```

where

```go
type Vendor struct {
	Name     string `json:"name,omitempty"`
	Constant string `json:"constant,omitempty"`
	Env      any    `json:"env,omitempty"`
	PR       any    `json:"pr,omitempty"`
	Val      string `json:"val,omitempty"`
}
```

Env will be a string after checking,
but due to vendor differences, is a
string, list, or map prior.

### releasing

git must be clean first

release: `goreleaser --rm-dist -p 1`

snapshot (test): `goreleaser --rm-dist -p 1 --snapshot`
