package logger

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	DebugF(format string, args ...interface{})
	InfoF(format string, args ...interface{})
	WarnF(format string, args ...interface{})
	ErrorF(format string, args ...interface{})
	FatalF(format string, args ...interface{})
}
