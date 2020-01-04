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
	"errors"
	"fmt"
	"log"
	"os"
)

// defaultLogger 默认的日志输出器
var defaultLogger StdLogger = &Console{}

// SetLogger 设置新的日志输出器
func SetLogger(logger StdLogger) {
	defaultLogger = logger
}

// Debug 打印 DEBUG 日志
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Debugf 打印 DEBUG 日志
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Info 打印 INFO 日志
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof 打印 INFO 日志
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warn 打印 WARN 日志
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf 打印 WARN 日志
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Error 打印 ERROR 日志
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Errorf 打印 ERROR 日志
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Fatal 打印 FATAL 日志
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf 打印 FATAL 日志
func Fatalf(format string, args ...interface{}) {
	log.Fatalln()
	defaultLogger.Fatalf(format, args...)
}

// Level 日志输出级别
type Level uint32

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	}
	panic(errors.New("error log level"))
}

// Console 将日志打印到控制台
type Console struct {
	level Level
}

// SetLevel 设置日志的输出级别
func (c *Console) SetLevel(level Level) {
	c.level = level
}

// Debug 打印 DEBUG 日志
func (c *Console) Debug(args ...interface{}) {
	if c.level <= DebugLevel {
		c.print(DebugLevel, args...)
	}
}

// Debugf 打印 DEBUG 日志
func (c *Console) Debugf(format string, args ...interface{}) {
	if c.level <= DebugLevel {
		c.printf(DebugLevel, format, args...)
	}
}

// Info 打印 INFO 日志
func (c *Console) Info(args ...interface{}) {
	if c.level <= InfoLevel {
		c.print(InfoLevel, args...)
	}
}

// Infof 打印 INFO 日志
func (c *Console) Infof(format string, args ...interface{}) {
	if c.level <= InfoLevel {
		c.printf(InfoLevel, format, args...)
	}
}

// Warn 打印 WARN 日志
func (c *Console) Warn(args ...interface{}) {
	if c.level <= WarnLevel {
		c.print(WarnLevel, args...)
	}
}

// Warnf 打印 WARN 日志
func (c *Console) Warnf(format string, args ...interface{}) {
	if c.level <= WarnLevel {
		c.printf(WarnLevel, format, args...)
	}
}

// Error 打印 ERROR 日志
func (c *Console) Error(args ...interface{}) {
	if c.level <= ErrorLevel {
		c.print(ErrorLevel, args...)
	}
}

// Errorf 打印 ERROR 日志
func (c *Console) Errorf(format string, args ...interface{}) {
	if c.level <= ErrorLevel {
		c.printf(ErrorLevel, format, args...)
	}
}

// Fatal 打印 FATAL 日志
func (c *Console) Fatal(args ...interface{}) {
	c.print(FatalLevel, args...)
	os.Exit(0)
}

// Fatalf 打印 FATAL 日志
func (c *Console) Fatalf(format string, args ...interface{}) {
	c.printf(FatalLevel, format, args...)
	os.Exit(0)
}

// print
func (c *Console) print(level Level, args ...interface{}) {
	str := fmt.Sprintf("[%s]", level)
	str += fmt.Sprint(args...)
	fmt.Println(str)
}

// printf
func (c *Console) printf(level Level, format string, args ...interface{}) {
	str := fmt.Sprintf("[%s] ", level)
	str += fmt.Sprintf(format, args...)
	fmt.Println(str)
}
