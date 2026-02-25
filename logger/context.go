package logger

import (
	"context"

	"go.uber.org/zap"
)

// Объявляем кастомный тип для ключа контекста во избежание коллизий
type contextKey int

const (
	// Устанавливаем единый ключ для логгера,через который и будем его доставать
	loggerContextKey contextKey = 0
)

// Передает в родительский контекст логгер.
func ToContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

// достает логгер из контекста. Если в контексте логгер не
// обнаруживается - возвращает глобальный логгер.
func FromContext(ctx context.Context) *zap.SugaredLogger {
	l := getLogger(ctx)

	return l
}

// создает логгер из уже имеющегося в контексте и устанавливает метаданные.
// Принимает ключ и значение, которые будут наследоваться дочерними логгерами.
func WithKV(ctx context.Context, key string, value any) context.Context {
	log := FromContext(ctx).With(key, value)
	return ToContext(ctx, log)
}

// Получение логгера либо из глобального инстанса, либо из контекста, вне зависимости от того,
// подходящий ли сейчас вызван уровень логгирования
func getLogger(ctx context.Context) *zap.SugaredLogger {
	logger := globalLogger
	if contextLogger, ok := ctx.Value(loggerContextKey).(*zap.SugaredLogger); ok {
		logger = contextLogger
	}
	return logger
}
