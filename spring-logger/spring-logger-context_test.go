/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package SpringLogger_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/go-spring/go-spring-parent/spring-logger"
)

type ContextLogger struct {
	ctx  context.Context
	tags []string
}

func (l *ContextLogger) CtxString() string {
	if v := l.ctx.Value("trace_id"); v != nil {
		return "trace_id:" + v.(string)
	}
	return ""
}

func (l *ContextLogger) TagString() string {
	if len(l.tags) == 0 {
		return ""
	}
	return strings.Join(l.tags, ",") + " "
}

func (l *ContextLogger) Printf(level string, format string, args ...interface{}) string {
	_, file, line, _ := runtime.Caller(3)
	if len(file) > 30 {
		file = "..." + file[len(file)-30:]
	}
	str := fmt.Sprintf("%s %s:%d %s %s", level, file, line, l.CtxString(), l.TagString())
	str += fmt.Sprintf(format, args...)
	fmt.Println(str)
	return str
}

func (l *ContextLogger) Debugf(format string, args ...interface{}) {
	l.Printf("[DEBUG]", format, args...)
}

func (l *ContextLogger) Debug(args ...interface{}) {
	fmt.Println(args...)
}

func (l *ContextLogger) Infof(format string, args ...interface{}) {
	l.Printf("[INFO]", format, args...)
}

func (l *ContextLogger) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (l *ContextLogger) Warnf(format string, args ...interface{}) {
	l.Printf("[WARN]", format, args...)
}

func (l *ContextLogger) Warn(args ...interface{}) {
	fmt.Println(args...)
}

func (l *ContextLogger) Errorf(format string, args ...interface{}) {
	l.Printf("[ERROR]", format, args...)
}

func (l *ContextLogger) Error(args ...interface{}) {
	fmt.Println(args...)
}

func (l *ContextLogger) Panicf(format string, args ...interface{}) {
	str := l.Printf("[PANIC]", format, args...)
	panic(errors.New(str))
}

func (l *ContextLogger) Panic(args ...interface{}) {
	fmt.Println(args...)
	panic(errors.New(""))
}

func (l *ContextLogger) Fatalf(format string, args ...interface{}) {
	l.Printf("[FATAL]", format, args...)
	os.Exit(1)
}

func (l *ContextLogger) Fatal(args ...interface{}) {
	fmt.Println(args...)
	os.Exit(1)
}

func TestDefaultTraceContext(t *testing.T) {

	// 设置全局转换函数
	SpringLogger.Logger = func(ctx context.Context, tags ...string) SpringLogger.StdLogger {
		return &ContextLogger{
			ctx:  ctx,
			tags: tags,
		}
	}

	ctx := context.WithValue(nil, "trace_id", "0689")
	tracer := SpringLogger.NewDefaultLoggerContext(ctx)

	fmt.Println()

	tracer.LogDebugf("level:%s %d", "debug", 0)
	tracer.LogInfof("level:%s %d", "info", 1)
	tracer.LogWarnf("level:%s %d", "warn", 2)
	tracer.LogErrorf("level:%s %d", "error", 3)
	//tracer.LogPanicf("level:%s %d", "panic", 4)
	//tracer.LogFatalf("level:%s %d", "fatal", 5)

	fmt.Println()

	tracer.Logger("__in").Debugf("level:%s %d", "debug", 0)
	tracer.Logger("__in").Infof("level:%s %d", "info", 1)
	tracer.Logger("__in").Warnf("level:%s %d", "warn", 2)
	tracer.Logger("__in").Errorf("level:%s %d", "error", 3)
	//tracer.Logger("__in").Panicf("level:%s %d", "panic", 4)
	//tracer.Logger("__in").Fatalf("level:%s %d", "fatal", 5)

	fmt.Println()
}
