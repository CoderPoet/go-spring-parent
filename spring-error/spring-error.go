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

package SpringError

import (
	"github.com/go-spring/go-spring-parent/spring-utils"
)

var (
	ERROR   = NewRpcError(-1, "ERROR")
	SUCCESS = NewRpcSuccess(0, "SUCCESS")
)

// ErrorCode 错误码
type ErrorCode struct {
	Code int32  // 错误码
	Msg  string // 错误信息
}

// NewErrorCode ErrorCode 的构造函数
func NewErrorCode(code int32, msg string) ErrorCode {
	return ErrorCode{
		Code: code,
		Msg:  msg,
	}
}

// RpcResult 定义 RPC 返回值
type RpcResult struct {
	ErrorCode

	Err  string      // 错误源
	Data interface{} // 返回值
}

// RpcSuccess 定义一个 RPC 成功值
type RpcSuccess ErrorCode

// NewRpcSuccess RpcSuccess 的构造函数
func NewRpcSuccess(code int32, msg string) RpcSuccess {
	return RpcSuccess(NewErrorCode(code, msg))
}

// Data 绑定一个值
func (r RpcSuccess) Data(data interface{}) *RpcResult {
	return &RpcResult{
		ErrorCode: ErrorCode(r),
		Data:      data,
	}
}

// RpcError 定义一个 RPC 异常值
type RpcError ErrorCode

// NewRpcError RpcError 的构造函数
func NewRpcError(code int32, msg string) RpcError {
	return RpcError(NewErrorCode(code, msg))
}

// Error 绑定一个错误
func (r RpcError) Error(err string) *RpcResult {
	return &RpcResult{
		ErrorCode: ErrorCode(r),
		Err:       err,
	}
}

// Panic 抛出一个异常值
func (r RpcError) Panic(err error) *SpringUtils.PanicCond {
	return SpringUtils.NewPanicCond(r.Error(err.Error()))
}
