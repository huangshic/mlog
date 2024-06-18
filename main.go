package main

import (
	"time"

	"github.com/huangshic/mlog/mlog"
	"github.com/huangshic/mlog/mlog/conf"
	"github.com/huangshic/mlog/mlog/zaplog"
	"go.uber.org/zap"
)

func main() {
	mlog.SetLogger(zaplog.New(
		conf.WithProjectName("pro"),
		conf.WithLogPath("log"),
		conf.WithLogLevel("info"),
		conf.WithMaxSize(100),
		conf.WithMaxAge(7),
		conf.WithIsStdOut("yes"),
		conf.WithLogName("test"),
	))
	mlog.Info("TTTTest info", zap.String("k", "j"))
	mlog.Error("test error", zap.String("k", "j"))
	time.Sleep(10 * time.Second)
}
