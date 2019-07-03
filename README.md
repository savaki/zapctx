# zapx

context wrapper for uber's zap library

#### Add logger to Context

```go
parent := context.Background()
ctx := zapctx.NewContext(parent, logger)
```

#### Retrieve logger from Context

```go
logger := zapctx.FromContext(ctx)
logger.Info("blah")
```

or used directly

```go
zapctx.FromContext(ctx).Info("blah")
```
