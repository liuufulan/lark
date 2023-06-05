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

// GetCoreHrNationalIDType 根据 ID 查询单个国家证件类型。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/national_id_type/get
func (r *CoreHrService) GetCoreHrNationalIDType(ctx context.Context, request *GetCoreHrNationalIDTypeReq, options ...MethodOptionFunc) (*GetCoreHrNationalIDTypeResp, *Response, error) {
	if r.cli.mock.mockCoreHrGetCoreHrNationalIDType != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#GetCoreHrNationalIDType mock enable")
		return r.cli.mock.mockCoreHrGetCoreHrNationalIDType(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "GetCoreHrNationalIDType",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/national_id_types/:national_id_type_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHrNationalIDTypeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrGetCoreHrNationalIDType mock CoreHrGetCoreHrNationalIDType method
func (r *Mock) MockCoreHrGetCoreHrNationalIDType(f func(ctx context.Context, request *GetCoreHrNationalIDTypeReq, options ...MethodOptionFunc) (*GetCoreHrNationalIDTypeResp, *Response, error)) {
	r.mockCoreHrGetCoreHrNationalIDType = f
}

// UnMockCoreHrGetCoreHrNationalIDType un-mock CoreHrGetCoreHrNationalIDType method
func (r *Mock) UnMockCoreHrGetCoreHrNationalIDType() {
	r.mockCoreHrGetCoreHrNationalIDType = nil
}

// GetCoreHrNationalIDTypeReq ...
type GetCoreHrNationalIDTypeReq struct {
	NationalIDTypeID string `path:"national_id_type_id" json:"-"` // 证件类型 ID, 示例值: "121515"
}

// GetCoreHrNationalIDTypeResp ...
type GetCoreHrNationalIDTypeResp struct {
	NationalIDType *GetCoreHrNationalIDTypeRespNationalIDType `json:"national_id_type,omitempty"` // 国家证件类型信息
}

// GetCoreHrNationalIDTypeRespNationalIDType ...
type GetCoreHrNationalIDTypeRespNationalIDType struct {
	ID                        string                                                                `json:"id,omitempty"`                          // 证件类型 ID
	CountryRegionID           string                                                                `json:"country_region_id,omitempty"`           // 国家 / 地区
	Name                      []*GetCoreHrNationalIDTypeRespNationalIDTypeName                      `json:"name,omitempty"`                        // 名称
	Active                    bool                                                                  `json:"active,omitempty"`                      // 是否启用
	ValidationRule            string                                                                `json:"validation_rule,omitempty"`             // 校验规则
	ValidationRuleDescription []*GetCoreHrNationalIDTypeRespNationalIDTypeValidationRuleDescription `json:"validation_rule_description,omitempty"` // 校验规则描述
	Code                      string                                                                `json:"code,omitempty"`                        // 编码
	IdentificationType        *GetCoreHrNationalIDTypeRespNationalIDTypeIdentificationType          `json:"identification_type,omitempty"`         // 证件类型
	CustomFields              []*GetCoreHrNationalIDTypeRespNationalIDTypeCustomField               `json:"custom_fields,omitempty"`               // 自定义字段
}

// GetCoreHrNationalIDTypeRespNationalIDTypeCustomField ...
type GetCoreHrNationalIDTypeRespNationalIDTypeCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHrNationalIDTypeRespNationalIDTypeIdentificationType ...
type GetCoreHrNationalIDTypeRespNationalIDTypeIdentificationType struct {
	EnumName string                                                                `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHrNationalIDTypeRespNationalIDTypeIdentificationTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHrNationalIDTypeRespNationalIDTypeIdentificationTypeDisplay ...
type GetCoreHrNationalIDTypeRespNationalIDTypeIdentificationTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHrNationalIDTypeRespNationalIDTypeName ...
type GetCoreHrNationalIDTypeRespNationalIDTypeName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHrNationalIDTypeRespNationalIDTypeValidationRuleDescription ...
type GetCoreHrNationalIDTypeRespNationalIDTypeValidationRuleDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHrNationalIDTypeResp ...
type getCoreHrNationalIDTypeResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHrNationalIDTypeResp `json:"data,omitempty"`
}