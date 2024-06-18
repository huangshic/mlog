package main

import (
	"github/huangshic/mlog/mlog"
	"github/huangshic/mlog/mlog/conf"
	"github/huangshic/mlog/mlog/zaplog"
	"time"

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
