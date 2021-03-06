package log

import (
	"context"

	"github.com/danielkrainas/gobag/context"
)

func Info(ctx context.Context, format string, args ...interface{}) {
	log := bagcontext.GetLogger(ctx)
	if len(args) > 0 {
		log.Infof(format, args...)
	} else {
		log.Info(format)
	}
}

func Debug(ctx context.Context, format string, args ...interface{}) {
	log := bagcontext.GetLogger(ctx)
	if len(args) > 0 {
		log.Infof(format, args...)
	} else {
		log.Info(format)
	}
}

func Warn(ctx context.Context, format string, args ...interface{}) {
	log := bagcontext.GetLogger(ctx)
	if len(args) > 0 {
		log.Warnf(format, args...)
	} else {
		log.Warn(format)
	}
}

func Error(ctx context.Context, format string, args ...interface{}) {
	log := bagcontext.GetLogger(ctx)
	if len(args) > 0 {
		log.Errorf(format, args...)
	} else {
		log.Error(format)
	}
}
