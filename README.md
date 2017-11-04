**Floop**

**When you need to execute function repeatedly**

# Usage

```go
package main

import (
	"fmt"
	"github.com/wuriyanto48/floop"
)

func pr() {
	res := mul(5, 5)
	fmt.Println(res)
}

func mul(a int, b int) int {
	return a * b
}

func main() {
	f := floop.New(1)
	f.Start(pr)
}
```