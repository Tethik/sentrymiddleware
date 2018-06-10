# echo-sentry-middleware
Another Sentry middleware for the Echo framework

## Usage
```go

func main() {
    // Echo instance
	e := echo.New()
	dsn, found := os.LookupEnv("SENTRY_RAVEN_DSN")
	if found {
		raven.SetDSN(dsn)
		e.Use(sentrymiddleware.SentryErrorRecover())
    }    
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
```


## Todo
1. Set DSN via this package (and use own client)
1. Tests
2. Better example.
3. Versioning.
