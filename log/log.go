package log

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"github.com/Bedrock-Technology/regen3/log/hooks"
	"github.com/davecgh/go-spew/spew"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

func Init(level string, logPath, serverName string, slackUrl string) (err error) {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	time.Local = time.UTC
	//logrus.SetFormatter(&logrus.JSONFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//})
	// rotate writer
	rotateLog, err := rotatelogs.New(
		logPath+".%Y%m%d",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(30*24*time.Hour), // 30 days
	)
	if err != nil {
		return
	}

	mw := io.MultiWriter(os.Stdout, rotateLog)
	// Use stdout for current dev
	// mw := io.MultiWriter(os.Stdout)
	logrus.SetOutput(mw)
	logrus.SetLevel(l)
	logrus.AddHook(hooks.NewCallStackHook())
	logrus.AddHook(hooks.NewServerHook(serverName))

	if slackUrl != "" {
		logrus.AddHook(hooks.NewSlackHook(slackUrl))
	}

	return
}

func PrintPanicStack(extras ...interface{}) {
	if x := recover(); x != nil {
		logrus.Errorf("%v", x)
		i := 0
		funcName, file, line, ok := runtime.Caller(i)
		for ok {
			logrus.Errorf("frame %v:[func:%v,file:%v,line:%v]", i, runtime.FuncForPC(funcName).Name(), file, line)
			i++
			funcName, file, line, ok = runtime.Caller(i)
		}

		for k := range extras {
			logrus.Errorf("EXRAS#%v DATA:%v", k, spew.Sdump(extras[k]))
		}
		panic("PANIC")
	}
}

type GormLogger struct {
	logger                    *logrus.Logger
	SlowThreshold             time.Duration
	Colorful                  bool
	IgnoreRecordNotFoundError bool
	LogLevel                  gormlogger.LogLevel
}

func NewGormLogger(logger *logrus.Logger) *GormLogger {
	return &GormLogger{
		logger:                    logger,
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  gormlogger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l *GormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	l.logger.WithContext(ctx).Logf(logrus.InfoLevel, s, []interface{}{args}...)
}

func (l *GormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	l.logger.WithContext(ctx).Logf(logrus.WarnLevel, s, []interface{}{args}...)
}

func (l *GormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	l.logger.WithContext(ctx).Logf(logrus.ErrorLevel, s, []interface{}{args}...)
}

var (
	infoStr      = "%s\n[info] "
	warnStr      = "%s\n[warn] "
	errStr       = "%s\n[error] "
	traceStr     = "%s\n[%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
)

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!errors.Is(err, gormlogger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.logger.Printf(traceWarnStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.logger.Printf(traceWarnStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.logger.Printf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.logger.Printf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == gormlogger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.logger.Printf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.logger.Printf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
