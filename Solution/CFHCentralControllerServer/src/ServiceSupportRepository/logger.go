package ServiceSupportRepository


var trace bool = false
var traceLogger StdLogger

func init() {
	traceLogger = Logger // use the package logger by default
}

// TraceLogger enables detailed logging of Http request matching and filter invocation. Default no logger is set.
// You may call EnableTracing() directly to enable trace logging to the package-wide logger.
func TraceLogger(logger StdLogger) {
	traceLogger = logger
	EnableTracing(logger != nil)
}

// expose the setter for the global logger on the top-level package
func SetLogger(customLogger StdLogger) {
	SetLogger(customLogger)
}

// EnableTracing can be used to Trace logging on and off.
func EnableTracing(enabled bool) {
	trace = enabled
}
