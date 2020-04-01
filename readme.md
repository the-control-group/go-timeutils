Go Timeutils
===

Usage
---

```go

package main

import (
	"fmt"
	"encoding/json"
	"github.com/the-control-group/go-timeutils"
)

type Config struct {
	Duration timeutils.ApproxBigDuration `json:"duration"`
}

func main() {
	c := Config{}
	json.Unmarshal([]byte(`{"duration":"23 days"}`), &c)
    fmt.Println(c.Duration.String())
    fmt.Println(c.Duration.ApproxString())
    fmt.Println(c.Duration.Pretty())
    fmt.Println(c.Duration.ApproxPretty())
    fmt.Println(time.Duration(c.Duration))
}

// Output
// 23d0h 0s
// ~23d0h
// 23 days, 0 hours 0s
// ~ 23 days, 0 hours
// 552h0m0s

```

Caveats
---

 - ApproxBigDuration: There's a reason it's called "Approx[imate]BigDuration" and there's a reason go's time package doesn't have a unit for inexact durations like days, months and years.