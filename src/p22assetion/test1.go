package p22assetion

type Logger interface {
	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})

	With(keyvals ...interface{}) Logger
}
type nopLogger struct{}

func (nopLogger) Info(string, ...interface{})  {}
func (nopLogger) Debug(string, ...interface{}) {}
func (nopLogger) Error(string, ...interface{}) {}

func (l *nopLogger) With(...interface{}) Logger {
	return l
}

// Interface assertions
var _ Logger = (*nopLogger)(nil)

func NewNopLogger() Logger {
	return &nopLogger{}
}
