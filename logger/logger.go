package logger

import (
	"context"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *zap.SugaredLogger
	defaultLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
)

func init() {
	SetLogger(NewLogger(defaultLevel))
}

func NewLogger(level zapcore.LevelEnabler, options ...zap.Option) *zap.SugaredLogger {
	return newLogger(level, os.Stdout, options...)
}

func newLogger(level zapcore.LevelEnabler, writer io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if writer == nil {
		writer = os.Stdout
	}
	loggerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.AddSync(writer),
		level,
	)

	return zap.New(loggerCore, options...).Sugar()
}

func SetLogger(logger *zap.SugaredLogger) {
	globalLogger = logger
}

func Debugf(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Debugf(format, args...)
}

func Debug(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Debugw(message, keyValues...)
}

func Infof(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Infof(format, args...)
}

func Info(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Infow(message, keyValues...)
}

func Warnf(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Warnf(format, args...)
}

func Warn(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Warnw(message, keyValues...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Errorf(format, args...)
}

func Error(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Errorw(message, keyValues...)
}

func Fatalf(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Fatalf(format, args...)
}

func Fatal(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Fatalw(message, keyValues...)
}

func Panicf(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Panicf(format, args...)
}

func Panic(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Panicw(message, keyValues...)
}
