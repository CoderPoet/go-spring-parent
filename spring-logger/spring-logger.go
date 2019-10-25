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

package SpringLogger

import (
	"fmt"
	"os"
)

//
// 标准的 Logger 接口
//
type StdLogger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

//
// 带前缀名的 Logger 接口
//
type PrefixLogger interface {
	LogDebug(args ...interface{})
	LogDebugf(format string, args ...interface{})

	LogInfo(args ...interface{})
	LogInfof(format string, args ...interface{})

	LogWarn(args ...interface{})
	LogWarnf(format string, args ...interface{})

	LogError(args ...interface{})
	LogErrorf(format string, args ...interface{})

	LogFatal(args ...interface{})
	LogFatalf(format string, args ...interface{})
}

var Console = &console{}

//
// 控制台打印
//
type console struct {
}

func (c *console) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func (c *console) Debug(args ...interface{}) {
	fmt.Println(args...)
}

func (c *console) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func (c *console) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (c *console) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func (c *console) Warn(args ...interface{}) {
	fmt.Println(args...)
}

func (c *console) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func (c *console) Error(args ...interface{}) {
	fmt.Println(args...)
}

func (c *console) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
	os.Exit(0)
}

func (c *console) Fatal(args ...interface{}) {
	fmt.Println(args...)
	os.Exit(0)
}

//
// 为了平衡调用栈的深度，增加一个 StdLogger 包装类
//
type StdLoggerWrapper struct {
	l StdLogger
}

func (w *StdLoggerWrapper) Debugf(format string, args ...interface{}) {
	w.l.Debugf(format, args...)
}

func (w *StdLoggerWrapper) Debug(args ...interface{}) {
	w.l.Debug(args...)
}

func (w *StdLoggerWrapper) Infof(format string, args ...interface{}) {
	w.l.Infof(format, args...)
}

func (w *StdLoggerWrapper) Info(args ...interface{}) {
	w.l.Info(args...)
}

func (w *StdLoggerWrapper) Warnf(format string, args ...interface{}) {
	w.l.Warnf(format, args...)
}

func (w *StdLoggerWrapper) Warn(args ...interface{}) {
	w.l.Warn(args...)
}

func (w *StdLoggerWrapper) Errorf(format string, args ...interface{}) {
	w.l.Errorf(format, args...)
}

func (w *StdLoggerWrapper) Error(args ...interface{}) {
	w.l.Error(args...)
}

func (w *StdLoggerWrapper) Fatalf(format string, args ...interface{}) {
	w.l.Fatalf(format, args...)
}

func (w *StdLoggerWrapper) Fatal(args ...interface{}) {
	w.l.Fatal(args...)
}
