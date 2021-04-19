# ctxgroup

A wrapper around [errgroup](golang.org/x/sync/errgroup), making context usage more obvious. Instead of relying on
shadowing errgroup's sub-context variables, it follows the "pass `context.Context` as the first argument of the
function" idiom, which simplifies the context scoping through variable shadowing.

Example usage:

```go
ctx := context.Background()
g := ctxgroup.WithContext(ctx)
g.GoCtx(func(ctx context.Context) error {
    // ctx gets cancelled when the other goroutine errors out or when Wait() returns
    <-ctx.Done()
    fmt.Println("first goroutine cancelled")
    return nil
})
g.GoCtx(func(ctx context.Context) error {
    return fmt.Errorf("second goroutine erroring out")
})
if err := g.Wait(); err != nil {
    fmt.Printf("ctxgroup.Wait returned: %v\n", err)
}
```
