package boot

import (
	"context"
	"fmt"
	"time"
)

type Logger interface {
	Info(ctx context.Context, msg string)
	InfoWithFormat(ctx context.Context, msg string, param ...interface{})
	Warn(ctx context.Context, msg string)
	WarnWithFormat(ctx context.Context, msg string, param ...interface{})
	Error(ctx context.Context, msg string)
	ErrorWithFormat(ctx context.Context, msg string, param ...interface{})
}

var allLog Logger

func RefLogger() Logger {
	return allLog
}
func injectLogger(all Logger) {
	allLog = all
}

type testLog struct {
}

func getCurrentTime() string {
	return time.Now().Format("2006-04-02 15:04:05")
}

func printWithMsg(level, msg string) {
	fmt.Printf("[%s]%s:%s\n", getCurrentTime(), level, msg)
}

func printWithMsgAndParam(level, msg string, param ...interface{}) {
	fmt.Printf("[%s]%s:%s\n", getCurrentTime(), level, fmt.Sprintf(msg, param))
}
func (t *testLog) Info(_ context.Context, msg string) {
	printWithMsg("Info", msg)
}

func (t *testLog) InfoWithFormat(_ context.Context, msg string, param ...interface{}) {
	printWithMsgAndParam("Info", msg, param)
}

func (t *testLog) Warn(_ context.Context, msg string) {
	printWithMsg("Warn", msg)
}

func (t *testLog) WarnWithFormat(_ context.Context, msg string, param ...interface{}) {
	printWithMsgAndParam("Warn", msg, param)
}

func (t *testLog) Error(_ context.Context, msg string) {
	printWithMsg("Error", msg)
}

func (t *testLog) ErrorWithFormat(_ context.Context, msg string, param ...interface{}) {
	printWithMsgAndParam("Error", msg, param)
}

func init() {
	injectLogger(new(testLog))
}
