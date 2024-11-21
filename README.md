# xerr
errors


## Examples

```go
func main() {
    e := xerr.New()
	fmt.Printf("%s\n", e)

	e2 := xerr.WrapError(e)
	fmt.Printf("%s\n", e2)

	e3 := xerr.NewCode(2)
	fmt.Printf("%s\n", e3)

	e4 := xerr.NewError("hello")
	fmt.Printf("%s\n", e4)
}

```
