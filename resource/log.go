// Package resource
// @author： Boice
// @createTime：2022/11/28 11:28
package resource

import (
	"context"
	"credit-platform/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type (
	Logger interface {
		Zap() *zap.Logger
		Info(ctx context.Context, msg string, fields ...zap.Field)
		Error(ctx context.Context, msg string, fields ...zap.Field)
		Debug(ctx context.Context, msg string, fields ...zap.Field)
		Warn(ctx context.Context, msg string, fields ...zap.Field)
		Panic(ctx context.Context, msg string, fields ...zap.Field)
	}
	logger struct {
		log *zap.Logger
	}
)

func (l *logger) Zap() *zap.Logger {
	return l.log
}

func (l *logger) Panic(ctx context.Context, msg string, fields ...zap.Field) {
	l.log.Panic(msg, l.withBaseFields(ctx, fields...)...)
}

func (l *logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.log.Info(msg, l.withBaseFields(ctx, fields...)...)
}

func (l *logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.log.Error(msg, l.withBaseFields(ctx, fields...)...)
}

func (l *logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	l.log.Debug(msg, l.withBaseFields(ctx, fields...)...)
}

func (l *logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.log.Warn(msg, l.withBaseFields(ctx, fields...)...)
}

func (l *logger) withBaseFields(ctx context.Context, fields ...zap.Field) []zap.Field {
	requestID, ok := ctx.Value(constant.CtxRequestID).(string)
	if ok {
		fields = append(fields, zap.String("request_id", requestID))
	}
	return fields
}

func newLogger(config Config) Logger {
	w := zapcore.AddSync(io.MultiWriter(&lumberjack.Logger{
		Filename:  config.Log.File,
		MaxSize:   1024, // MB
		LocalTime: true,
		Compress:  true,
	}, os.Stdout))
	level, err := zap.ParseAtomicLevel(config.Log.Level)
	if err != nil {
		panic(err)
	}
	zapConfig := zap.NewProductionEncoderConfig()
	zapConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapConfig),
		w,
		level,
	)

	return &logger{
		log: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)),
	}
}
