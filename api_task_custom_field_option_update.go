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

// UpdateTaskCustomFieldOption 根据一个自定义字段的GUID和其选项的GUID, 更新该选项的数据。要更新的字段必须是单选或者多选类型, 且要更新的字段必须归属于该字段。
//
// 更新时, 将`update_fields`字段中填写所有要修改的任务字段名, 同时在`option`字段中填写要修改的字段的新值即可。`update_fields`支持的字段包括:
// * `name`: 选项名称
// * `color_index`: 选项的颜色索引值
// * `is_hidden`: 是否从界面上隐藏
// * `insert_before`: 将当前option放到同字段某个option之前的那个option_guid。
// * `insert_after`: 将当前option放到同字段某个option之后的那个option_guid。
// 更新选项需要自定义字段的编辑权限
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field-option/patch
func (r *TaskService) UpdateTaskCustomFieldOption(ctx context.Context, request *UpdateTaskCustomFieldOptionReq, options ...MethodOptionFunc) (*UpdateTaskCustomFieldOptionResp, *Response, error) {
	if r.cli.mock.mockTaskUpdateTaskCustomFieldOption != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#UpdateTaskCustomFieldOption mock enable")
		return r.cli.mock.mockTaskUpdateTaskCustomFieldOption(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "UpdateTaskCustomFieldOption",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/custom_fields/:custom_field_guid/options/:option_guid",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateTaskCustomFieldOptionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskUpdateTaskCustomFieldOption mock TaskUpdateTaskCustomFieldOption method
func (r *Mock) MockTaskUpdateTaskCustomFieldOption(f func(ctx context.Context, request *UpdateTaskCustomFieldOptionReq, options ...MethodOptionFunc) (*UpdateTaskCustomFieldOptionResp, *Response, error)) {
	r.mockTaskUpdateTaskCustomFieldOption = f
}

// UnMockTaskUpdateTaskCustomFieldOption un-mock TaskUpdateTaskCustomFieldOption method
func (r *Mock) UnMockTaskUpdateTaskCustomFieldOption() {
	r.mockTaskUpdateTaskCustomFieldOption = nil
}

// UpdateTaskCustomFieldOptionReq ...
type UpdateTaskCustomFieldOptionReq struct {
	CustomFieldGuid string                                `path:"custom_field_guid" json:"-"` // 要更新的选项的自定义字段GUID, 示例值: "ec5ed63d-a4a9-44de-a935-7ba243471c0a"
	OptionGuid      string                                `path:"option_guid" json:"-"`       // 要更新的选项的GUID, 示例值: "b13adf3c-cad6-4e02-8929-550c112b5633"
	Option          *UpdateTaskCustomFieldOptionReqOption `json:"option,omitempty"`           // 要更新的option数据
	UpdateFields    []string                              `json:"update_fields,omitempty"`    // 要更新的字段名, 支持, * `name`: 选项名称, * `color_index`: 选项的颜色索引值, * `is_hidden`: 是否从界面上隐藏, * `insert_before`: 将当前option放到同字段某个option之前, * `insert_after`: 将当前option放到同字段某个option之后, 示例值: ["name"], 长度范围: `1` ～ `20`
}

// UpdateTaskCustomFieldOptionReqOption ...
type UpdateTaskCustomFieldOptionReqOption struct {
	Name         *string `json:"name,omitempty"`          // 选项名称, 最大50个字符, 示例值: "高优"
	ColorIndex   *int64  `json:"color_index,omitempty"`   // 颜色索引值, 支持0～54中的一个数字, 示例值: 10, 取值范围: `0` ～ `54`
	InsertBefore *string `json:"insert_before,omitempty"` // 要放到某个option之前的option_guid, 示例值: "2bd905f8-ef38-408b-aa1f-2b2ad33b2913"
	InsertAfter  *string `json:"insert_after,omitempty"`  // 要放到某个option之后的option_guid, 示例值: "b13adf3c-cad6-4e02-8929-550c112b5633"
	IsHidden     *bool   `json:"is_hidden,omitempty"`     // 是否隐藏, 示例值: false, 默认值: `false`
}

// UpdateTaskCustomFieldOptionResp ...
type UpdateTaskCustomFieldOptionResp struct {
	Option *UpdateTaskCustomFieldOptionRespOption `json:"option,omitempty"` // 更新后的option数据
}

// UpdateTaskCustomFieldOptionRespOption ...
type UpdateTaskCustomFieldOptionRespOption struct {
	Guid       string `json:"guid,omitempty"`        // 选项的GUID
	Name       string `json:"name,omitempty"`        // 选项名称, 不能为空, 最大50个字符
	ColorIndex int64  `json:"color_index,omitempty"` // 选项的颜色索引值, 可以是0～54中的一个数字。
	IsHidden   bool   `json:"is_hidden,omitempty"`   // 选项是否隐藏。隐藏后的选项在界面不可见, 也不可以再通过openapi将字段值设为该选项。
}

// updateTaskCustomFieldOptionResp ...
type updateTaskCustomFieldOptionResp struct {
	Code  int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                           `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateTaskCustomFieldOptionResp `json:"data,omitempty"`
	Error *ErrorDetail                     `json:"error,omitempty"`
}