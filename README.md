# go-case
Package to convert names, titles and sequences to snake_case, CamelCase and some other.

```go
import (
    "fmt"

    "github.com/vetcher/go-case"
)

func main() {
    s := "some string to convert"
    fmt.Println(go_case.ToCamelCase(s))

    // Output:
    // SomeStringToConvert
}
```