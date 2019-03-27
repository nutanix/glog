package glog

import (
	"context"
	"fmt"
)

// Logger provides logging functionality with additional prefixing.
type Logger struct {
	*loggingT
	Verbose
	prefix string
}

// NewLogger creates a Logger instance with empty prefix.
func NewLogger() *Logger {
	return &Logger{
		loggingT: &logging,
		Verbose:  V(logging.verbosity),
	}
}

// NewLoggerWithPrefix creates a Logger with a given prefix.
func NewLoggerWithPrefix(format string, a ...interface{}) *Logger {
	return &Logger{
		loggingT: &logging,
		Verbose:  V(logging.verbosity),
		prefix:   fmt.Sprintf(format, a...),
	}
}

// AddPrefix appends existing logger with a specified prefix.
func (l *Logger) AddPrefix(format string, a ...interface{}) *Logger {
	l.prefix += fmt.Sprintf(format, a...)
	return l
}

// Info is equivalent to the global Info function, with the addition of prefix from this Logger.
func (l *Logger) Info(args ...interface{}) {
	if l.Verbose {
		l.print(infoLog, l.extendWithPrefix(args)...)
	}
}

// Infoln is equivalent to the global Infoln function, with the addition of prefix from this Logger.
func (l *Logger) Infoln(args ...interface{}) {
	if l.Verbose {
		l.println(infoLog, l.extendWithPrefix(args)...)
	}
}

// Infof is equivalent to the global Infof function, with the addition of prefix from this Logger.
func (l *Logger) Infof(format string, args ...interface{}) {
	if l.Verbose {
		l.printf(infoLog, l.pfx(format), args...)
	}
}

// Warning is equivalent to the global Warning function, with the addition of prefix from this Logger.
func (l *Logger) Warning(args ...interface{}) {
	if l.Verbose {
		l.print(warningLog, l.extendWithPrefix(args)...)
	}
}

// Warningln is equivalent to the global Warningln function, with the addition of prefix from this Logger.
func (l *Logger) Warningln(args ...interface{}) {
	if l.Verbose {
		l.println(warningLog, l.extendWithPrefix(args)...)
	}
}

// Warningf is equivalent to the global Warningf function, with the addition of prefix from this Logger.
func (l *Logger) Warningf(format string, args ...interface{}) {
	if l.Verbose {
		l.printf(warningLog, l.pfx(format), args...)
	}
}

// Error is equivalent to the global Error function, with the addition of prefix from this Logger.
func (l *Logger) Error(args ...interface{}) {
	if l.Verbose {
		l.print(errorLog, l.extendWithPrefix(args)...)
	}
}

// Errorln is equivalent to the global Errorln function, with the addition of prefix from this Logger.
func (l *Logger) Errorln(args ...interface{}) {
	if l.Verbose {
		l.println(errorLog, l.extendWithPrefix(args)...)
	}
}

// Errorf is equivalent to the global Errorf function, with the addition of prefix from this Logger.
func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.Verbose {
		l.printf(errorLog, l.pfx(format), args...)
	}
}

// Fatal is equivalent to the global Fatal function, with the addition of prefix from this Logger.
func (l *Logger) Fatal(args ...interface{}) {
	if l.Verbose {
		l.print(fatalLog, l.extendWithPrefix(args)...)
	}
}

// Fatalln is equivalent to the global Fatalln function, with the addition of prefix from this Logger.
func (l *Logger) Fatalln(args ...interface{}) {
	if l.Verbose {
		l.println(fatalLog, l.extendWithPrefix(args)...)
	}
}

// Fatalf is equivalent to the global Fatalf function, with the addition of prefix from this Logger.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.printf(fatalLog, l.pfx(format), args...)
}

func (l *Logger) extendWithPrefix(args []interface{}) []interface{} {
	if l.prefix != "" {
		args = append([]interface{}{
			l.prefix,
		}, args...)
	}
	return args
}

func (l *Logger) pfx(log string) string {
	return fmt.Sprintf("%v %v", l.prefix, log)
}

// V reports whether verbosity at the call site is at least the requested level.
// The returned value is a boolean of type Verbose, which implements Info, Infoln
// and Infof. These methods will write to the Info log if called.
// Thus, one may write either
//	if glog.V(2) { glog.Info("log this") }
// or
//	glog.V(2).Info("log this")
// The second form is shorter but the first is cheaper if logging is off because it does
// not evaluate its arguments.
//
// Whether an individual call to V generates a log record depends on the setting of
// the -v and --vmodule flags; both are off by default. If the level in the call to
// V is at least the value of -v, or of -vmodule for the source file containing the
// call, the V call will log.
func (l Logger) V(level Level) *Logger {
	l.Verbose = V(level)
	return &l
}

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// glogKey is the key for glog.Logger values in Contexts. It is
// unexported; clients use glog.NewContext and glog.FromContext
// instead of using this key directly.
const glogKey key = iota + 1

// NewContext returns a new Context that carries value u.
func (l *Logger) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, glogKey, l)
}

// FromContext returns the Logger value stored in ctx, if any.
// Creates a new logger if not present
func FromContext(ctx context.Context) (l *Logger) {
	var ok bool
	if l, ok = ctx.Value(glogKey).(*Logger); !ok {
		l = NewLogger()
	}

	return l
}
