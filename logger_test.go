package glog

import (
	"context"
	"strings"
	"testing"
)

// Test that using the prefix works.
func TestPrefix(t *testing.T) {

	for _, test := range loggerTests {
		t.Run(test.name, func(t *testing.T) {

			setFlags()
			defer logging.swap(logging.newBuffers())

			firstPrefix := "examplePrefix"
			prefixedLogger := NewLoggerWithPrefix(firstPrefix)
			test.loggingFunc(prefixedLogger)
			if !contains(test.severity, test.logCharacter, t) {
				t.Errorf("%s has wrong character: %q", severityName[test.severity], contents(test.severity))
			}
			if !contains(test.severity, "hello", t) {
				t.Errorf("%d failed", test.severity)
			}
			if !contains(test.severity, firstPrefix, t) {
				t.Errorf("%d failed", test.severity)
			}
		})
	}
}

func TestPrefixedLoggingWithContext(t *testing.T) {

	for _, test := range loggerTests {
		t.Run(test.name, func(t *testing.T) {

			setFlags()
			defer logging.swap(logging.newBuffers())

			firstPrefix := "examplePrefix"
			prefixedLogger := NewLoggerWithPrefix(firstPrefix)
			ctx := prefixedLogger.NewContext(context.Background())
			test.loggingFunc(FromContext(ctx))
			if !contains(test.severity, test.logCharacter, t) {
				t.Errorf("%s has wrong character: %q", severityName[test.severity], contents(test.severity))
			}
			if !contains(test.severity, firstPrefix, t) {
				t.Errorf("%d failed", test.severity)
			}
		})
	}
}

func TestMultiPrefix(t *testing.T) {

	for _, test := range loggerTests {
		t.Run(test.name, func(t *testing.T) {

			setFlags()
			defer logging.swap(logging.newBuffers())

			firstPrefix := "examplePrefix"
			secondPrefix := "secondPrefix"
			prefixedLogger := NewLoggerWithPrefix(firstPrefix)
			prefixedLogger.AddPrefix(secondPrefix)

			test.loggingFunc(prefixedLogger)
			if !contains(test.severity, test.logCharacter, t) {
				t.Errorf("%s has wrong character: %q", severityName[test.severity], contents(test.severity))
			}
			if !contains(test.severity, "hello", t) {
				t.Errorf("%s failed", severityName[test.severity])
			}
			if !contains(test.severity, firstPrefix, t) {
				t.Errorf("%d failed", test.severity)
			}
			if !contains(test.severity, secondPrefix, t) {
				t.Errorf("%d failed", test.severity)
			}
			if !contains(test.severity, firstPrefix+secondPrefix, t) {
				t.Errorf("%s failed", severityName[test.severity])
			}
		})
	}
}

func TestMultiPrefixWithPoppedPrefix(t *testing.T) {

	for _, test := range loggerTests {
		t.Run(test.name, func(t *testing.T) {

			setFlags()
			defer logging.swap(logging.newBuffers())

			firstPrefix := "examplePrefix"
			secondPrefix := "secondPrefix"
			prefixedLogger := NewLoggerWithPrefix(firstPrefix)
			prefixedLogger.AddPrefix(secondPrefix)

			test.loggingFunc(prefixedLogger.PopPrefix())
			if !contains(test.severity, "hello", t) {
				t.Errorf("%s failed", severityName[test.severity])
			}
			if !contains(test.severity, firstPrefix, t) {
				t.Errorf("%d failed", test.severity)
			}
			// second prefix should not be seen anymore since its been popped out
			if contains(test.severity, secondPrefix, t) {
				t.Errorf("%d failed", test.severity)
			}
			if contains(test.severity, firstPrefix+secondPrefix, t) {
				t.Errorf("%s failed", severityName[test.severity])
			}
		})
	}
}

func TestMultiPrefixWithLogLevel(t *testing.T) {

	for _, test := range verboseLoggerTests {
		t.Run(test.name, func(t *testing.T) {
			setFlags()
			defer logging.swap(logging.newBuffers())
			firstPrefix := "examplePrefix"
			secondPrefix := "secondPrefix"
			prefixedLogger := NewLoggerWithPrefix(firstPrefix)
			prefixedLogger.AddPrefix(secondPrefix)
			prefixedLogger.verbosity.set(3)
			test.loggingFunc(prefixedLogger)

			if contains(test.severity, test.logCharacter, t) {
				t.Errorf("%s has wrong character: %q", severityName[test.severity], contents(test.severity))
			}
			if contains(test.severity, "hello", t) {
				t.Errorf("%s failed", severityName[test.severity])
			}
			if contains(test.severity, strings.Join([]string{firstPrefix, secondPrefix}, ""), t) {
				t.Errorf("%s failed", severityName[test.severity])
			}

			prefixedLogger.verbosity.set(4)
			test.loggingFunc(prefixedLogger)

			if !contains(test.severity, test.logCharacter, t) {
				t.Errorf("%s has wrong character: %q", severityName[test.severity], contents(test.severity))
			}
			if !contains(test.severity, "hello", t) {
				t.Errorf("%s failed", severityName[test.severity])
			}
			if !contains(test.severity, strings.Join([]string{firstPrefix, secondPrefix}, ""), t) {
				t.Errorf("%s failed", severityName[test.severity])
			}
		})
	}
}

var loggerTests = []struct {
	name         string
	severity     severity
	logCharacter string
	loggingFunc  func(l *Logger)
}{
	{
		name:         "Info",
		severity:     infoLog,
		logCharacter: "I",
		loggingFunc: func(l *Logger) {
			l.Info("hello")
		},
	},
	{
		name:         "Infoln",
		severity:     infoLog,
		logCharacter: "I",
		loggingFunc: func(l *Logger) {
			l.Infoln("hello")
		},
	},
	{
		name:         "Infof",
		severity:     infoLog,
		logCharacter: "I",
		loggingFunc: func(l *Logger) {
			l.Infof("hello: %s", "<NAME>")
		},
	},
	{
		name:         "Warning",
		severity:     warningLog,
		logCharacter: "W",
		loggingFunc: func(l *Logger) {
			l.Warning("hello")
		},
	},
	{
		name:         "Warningln",
		severity:     warningLog,
		logCharacter: "W",
		loggingFunc: func(l *Logger) {
			l.Warningln("hello")
		},
	},
	{
		name:         "Warningf",
		severity:     warningLog,
		logCharacter: "W",
		loggingFunc: func(l *Logger) {
			l.Warningf("hello: %s", "<NAME>")
		},
	},
	{
		name:         "Error",
		severity:     errorLog,
		logCharacter: "E",
		loggingFunc: func(l *Logger) {
			l.Error("hello")
		},
	},
	{
		name:         "Errorln",
		severity:     errorLog,
		logCharacter: "E",
		loggingFunc: func(l *Logger) {
			l.Errorln("hello")
		},
	},
	{
		name:         "Errorf",
		severity:     errorLog,
		logCharacter: "E",
		loggingFunc: func(l *Logger) {
			l.Errorf("hello: %s", "<NAME>")
		},
	},
}

var verboseLoggerTests = []struct {
	name         string
	severity     severity
	logCharacter string
	loggingFunc  func(l *Logger)
}{
	{
		name:         "Info",
		severity:     infoLog,
		logCharacter: "I",
		loggingFunc: func(l *Logger) {
			l.V(4).Info("hello")
		},
	},
	{
		name:         "Infoln",
		severity:     infoLog,
		logCharacter: "I",
		loggingFunc: func(l *Logger) {
			l.V(4).Infoln("hello")
		},
	},
	{
		name:         "Infof",
		severity:     infoLog,
		logCharacter: "I",
		loggingFunc: func(l *Logger) {
			l.V(4).Infof("hello: %s", "<NAME>")
		},
	},
	{
		name:         "Warning",
		severity:     warningLog,
		logCharacter: "W",
		loggingFunc: func(l *Logger) {
			l.V(4).Warning("hello")
		},
	},
	{
		name:         "Warningln",
		severity:     warningLog,
		logCharacter: "W",
		loggingFunc: func(l *Logger) {
			l.V(4).Warningln("hello")
		},
	},
	{
		name:         "Warningf",
		severity:     warningLog,
		logCharacter: "W",
		loggingFunc: func(l *Logger) {
			l.V(4).Warningf("hello: %s", "<NAME>")
		},
	},
	{
		name:         "Error",
		severity:     errorLog,
		logCharacter: "E",
		loggingFunc: func(l *Logger) {
			l.V(4).Error("hello")
		},
	},
	{
		name:         "Errorln",
		severity:     errorLog,
		logCharacter: "E",
		loggingFunc: func(l *Logger) {
			l.V(4).Errorln("hello")
		},
	},
	{
		name:         "Errorf",
		severity:     errorLog,
		logCharacter: "E",
		loggingFunc: func(l *Logger) {
			l.V(4).Errorf("hello: %s", "<NAME>")
		},
	},
}
