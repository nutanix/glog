package glog

import (
	"bytes"
	"testing"
)

var fakeStdout bytes.Buffer

// Test that using the prefix works.
func TestPrefix(t *testing.T) {

	for _, test := range loggerTests {
		t.Run(test.name, func(t *testing.T) {

			setFlags()
			defer logging.swap(logging.newBuffers())

			prefixedLogger := NewLoggerWithPrefix("examplePrefix")
			test.loggingFunc(prefixedLogger)
			if !contains(test.severity, test.logCharacter, t) {
				t.Errorf("%s has wrong character: %q", severityName[test.severity], contents(test.severity))
			}
			if !contains(test.severity, "hello", t) {
				t.Errorf("%s failed", test.severity)
			}
			if !contains(test.severity, "examplePrefix", t) {
				t.Errorf("%s failed", test.severity)
			}
		})
	}
}

func TestMultiPrefix(t *testing.T) {

	for _, test := range loggerTests {
		t.Run(test.name, func(t *testing.T) {

			setFlags()
			defer logging.swap(logging.newBuffers())

			prefixedLogger := NewLoggerWithPrefix("examplePrefix")
			prefixedLogger.AddPrefix("secondPrefix")
			test.loggingFunc(prefixedLogger)
			if !contains(test.severity, test.logCharacter, t) {
				t.Errorf("%s has wrong character: %q", severityName[test.severity], contents(test.severity))
			}
			if !contains(test.severity, "hello", t) {
				t.Errorf("%s failed", severityName[test.severity])
			}
			if !contains(test.severity, "examplePrefixsecondPrefix", t) {
				t.Errorf("%s failed", severityName[test.severity], )
			}
		})
	}
}

func TestMultiPrefixWithLogLevel(t *testing.T) {

	for _, test := range verboseLoggerTests {
		t.Run(test.name, func(t *testing.T) {
			setFlags()
			defer logging.swap(logging.newBuffers())

			prefixedLogger := NewLoggerWithPrefix("examplePrefix")
			prefixedLogger.AddPrefix("secondPrefix")
			prefixedLogger.verbosity.set(3)
			test.loggingFunc(prefixedLogger)

			if contains(test.severity, test.logCharacter, t) {
				t.Errorf("%s has wrong character: %q", severityName[test.severity], contents(test.severity))
			}
			if contains(test.severity, "hello", t) {
				t.Errorf("%s failed", severityName[test.severity])
			}
			if contains(test.severity, "examplePrefixsecondPrefix", t) {
				t.Errorf("%s failed", severityName[test.severity], )
			}

			prefixedLogger.verbosity.set(4)
			test.loggingFunc(prefixedLogger)

			if !contains(test.severity, test.logCharacter, t) {
				t.Errorf("%s has wrong character: %q", severityName[test.severity], contents(test.severity))
			}
			if !contains(test.severity, "hello", t) {
				t.Errorf("%s failed", severityName[test.severity])
			}
			if !contains(test.severity, "examplePrefixsecondPrefix", t) {
				t.Errorf("%s failed", severityName[test.severity], )
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
