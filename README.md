# Models

Shared domain models and small utility types generated from OpenAPI-style YAML definitions.

## Packages

* `utility/golang/dmutil`
* `utility/golang/dmtime`
* `utility/definitions`

## Quick Start

```golang
package quick_start

import (
    "fmt"
    "time"

    "github.com/Deimvis/models/utility/golang/dmtime"
)

func main() {
    data := []byte(`{"d":0,"h":0,"m":1,"s":0,"ms":0,"mcs":0,"ns":0}`)
    var dur dmtime.ConstDuration
    if err := dur.UnmarshalJSON(data); err != nil {
        panic(err)
    }
    fmt.Println(dur.Duration() == time.Minute)
}
```
