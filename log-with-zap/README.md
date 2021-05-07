# Test Uber ZAP logger

Test changing [ZAP](https://pkg.go.dev/go.uber.org/zap?utm_source=godoc) log level on the fly without application restart.

```bash
curl -X PUT -d '{"level":"debug"}' localhost:1065/log_level
curl -X PUT -d '{"level":"info"}' localhost:1065/log_level
curl -X PUT -d '{"level":"info"}' localhost:1065/  # EXPECTED 404
curl -X GET localhost:1065/log_level
```