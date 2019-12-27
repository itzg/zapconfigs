package zapconfigs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func NewDebugEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func NewDefaultEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		MessageKey:     "M",
		// leave out stack trace and caller
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// NewDebugLogger provides a logger suitable for logging in production code, but at debug level.
// This and NewDefaultLogger both use "console" style encoding to ensure that they are interchangeable.
// If an error occurs during logger creation, then log.Fatal is used to report the error and exit.
func NewDebugLogger(options ...zap.Option) *zap.Logger {
	zapConfig := zap.Config{
		Encoding: "console",
		EncoderConfig: NewDebugEncoderConfig(),
		Level: zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths: []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := zapConfig.Build(options...)
	if err != nil {
		log.Fatal(err)
	}
	return logger
}

// NewDefaultLogger provides a human-readable logger suitable for production code. It logs at
// info level and excludes the caller and stack trace fields.
// This and NewDebugLogger both use "console" style encoding to ensure that they are interchangeable.
// If an error occurs during logger creation, then log.Fatal is used to report the error and exit.
func NewDefaultLogger(options ...zap.Option) (*zap.Logger) {
	return NewLeveledLogger(zapcore.InfoLevel, options...)
}

// NewLeveledLogger is the same as NewDefaultLogger but allows for specifying the logging level.
func NewLeveledLogger(level zapcore.Level, options ...zap.Option) (*zap.Logger) {
	zapConfig := zap.Config{
		Encoding: "console",
		EncoderConfig: NewDefaultEncoderConfig(),
		Level: zap.NewAtomicLevelAt(level),
		OutputPaths: []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := zapConfig.Build(options...)
	if err != nil {
		log.Fatal(err)
	}
	return logger
}
