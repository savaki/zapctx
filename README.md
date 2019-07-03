# zapx

context wrapper for uber's zap library

#### Add logger to Context

```go
parent := context.Background()
ctx := zapx.NewContext(parent, logger)
```

#### Retrieve logger from Context

```go
logger := zapx.FromContext(ctx)
logger.Info("blah")
```

or used directly

```go
zapx.FromContext(ctx).Info("blah")
```
