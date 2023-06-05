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

// GetCoreHrJobLevel 根据 ID 查询单个职级
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_level/get
func (r *CoreHrService) GetCoreHrJobLevel(ctx context.Context, request *GetCoreHrJobLevelReq, options ...MethodOptionFunc) (*GetCoreHrJobLevelResp, *Response, error) {
	if r.cli.mock.mockCoreHrGetCoreHrJobLevel != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#GetCoreHrJobLevel mock enable")
		return r.cli.mock.mockCoreHrGetCoreHrJobLevel(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "GetCoreHrJobLevel",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/job_levels/:job_level_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHrJobLevelResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrGetCoreHrJobLevel mock CoreHrGetCoreHrJobLevel method
func (r *Mock) MockCoreHrGetCoreHrJobLevel(f func(ctx context.Context, request *GetCoreHrJobLevelReq, options ...MethodOptionFunc) (*GetCoreHrJobLevelResp, *Response, error)) {
	r.mockCoreHrGetCoreHrJobLevel = f
}

// UnMockCoreHrGetCoreHrJobLevel un-mock CoreHrGetCoreHrJobLevel method
func (r *Mock) UnMockCoreHrGetCoreHrJobLevel() {
	r.mockCoreHrGetCoreHrJobLevel = nil
}

// GetCoreHrJobLevelReq ...
type GetCoreHrJobLevelReq struct {
	JobLevelID string `path:"job_level_id" json:"-"` // 职级 ID, 示例值: "1515"
}

// GetCoreHrJobLevelResp ...
type GetCoreHrJobLevelResp struct {
	JobLevel *GetCoreHrJobLevelRespJobLevel `json:"job_level,omitempty"` // 职级信息
}

// GetCoreHrJobLevelRespJobLevel ...
type GetCoreHrJobLevelRespJobLevel struct {
	ID           string                                      `json:"id,omitempty"`            // 职级 ID
	LevelOrder   int64                                       `json:"level_order,omitempty"`   // 职级数值
	Code         string                                      `json:"code,omitempty"`          // 编码
	Name         []*GetCoreHrJobLevelRespJobLevelName        `json:"name,omitempty"`          // 名称
	Description  []*GetCoreHrJobLevelRespJobLevelDescription `json:"description,omitempty"`   // 描述
	Active       bool                                        `json:"active,omitempty"`        // 是否启用
	CustomFields []*GetCoreHrJobLevelRespJobLevelCustomField `json:"custom_fields,omitempty"` // 自定义字段
}

// GetCoreHrJobLevelRespJobLevelCustomField ...
type GetCoreHrJobLevelRespJobLevelCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHrJobLevelRespJobLevelDescription ...
type GetCoreHrJobLevelRespJobLevelDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHrJobLevelRespJobLevelName ...
type GetCoreHrJobLevelRespJobLevelName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHrJobLevelResp ...
type getCoreHrJobLevelResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHrJobLevelResp `json:"data,omitempty"`
}
