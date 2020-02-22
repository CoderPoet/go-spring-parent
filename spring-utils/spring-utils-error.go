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

package SpringUtils

// ToString 获取 error 的字符串
func ToString(err error) string {
	if err != nil {
		return err.Error()
	} else {
		return ""
	}
}

// Panic 返回一个封装的 panic 条件
func Panic(err error) *PanicCond {
	return NewPanicCond(err)
}

// PanicCond 封装触发 panic 的条件
type PanicCond struct {
	e interface{}
}

// NewPanicCond PanicCond 的构造函数
func NewPanicCond(e interface{}) *PanicCond {
	return &PanicCond{e}
}

// When 满足给定条件时抛出一个 panic
func (p *PanicCond) When(isPanic bool) {
	if isPanic {
		panic(p.e)
	}
}
