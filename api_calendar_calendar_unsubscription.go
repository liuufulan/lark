// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// UnsubscribeCalendarChangeEvent 该接口用于以用户身份取消订阅当前身份下日历列表中的日历变更事件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/unsubscription
func (r *CalendarService) UnsubscribeCalendarChangeEvent(ctx context.Context, request *UnsubscribeCalendarChangeEventReq, options ...MethodOptionFunc) (*UnsubscribeCalendarChangeEventResp, *Response, error) {
	if r.cli.mock.mockCalendarUnsubscribeCalendarChangeEvent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#UnsubscribeCalendarChangeEvent mock enable")
		return r.cli.mock.mockCalendarUnsubscribeCalendarChangeEvent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Calendar",
		API:                 "UnsubscribeCalendarChangeEvent",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/unsubscription",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(unsubscribeCalendarChangeEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarUnsubscribeCalendarChangeEvent mock CalendarUnsubscribeCalendarChangeEvent method
func (r *Mock) MockCalendarUnsubscribeCalendarChangeEvent(f func(ctx context.Context, request *UnsubscribeCalendarChangeEventReq, options ...MethodOptionFunc) (*UnsubscribeCalendarChangeEventResp, *Response, error)) {
	r.mockCalendarUnsubscribeCalendarChangeEvent = f
}

// UnMockCalendarUnsubscribeCalendarChangeEvent un-mock CalendarUnsubscribeCalendarChangeEvent method
func (r *Mock) UnMockCalendarUnsubscribeCalendarChangeEvent() {
	r.mockCalendarUnsubscribeCalendarChangeEvent = nil
}

// UnsubscribeCalendarChangeEventReq ...
type UnsubscribeCalendarChangeEventReq struct {
}

// UnsubscribeCalendarChangeEventResp ...
type UnsubscribeCalendarChangeEventResp struct {
}

// unsubscribeCalendarChangeEventResp ...
type unsubscribeCalendarChangeEventResp struct {
	Code int64                               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 错误描述
	Data *UnsubscribeCalendarChangeEventResp `json:"data,omitempty"`
}
