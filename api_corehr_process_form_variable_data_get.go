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

// GetCoreHrProcessFormVariableData 获取流程表单数据。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/process-form_variable_data/get
func (r *CoreHrService) GetCoreHrProcessFormVariableData(ctx context.Context, request *GetCoreHrProcessFormVariableDataReq, options ...MethodOptionFunc) (*GetCoreHrProcessFormVariableDataResp, *Response, error) {
	if r.cli.mock.mockCoreHrGetCoreHrProcessFormVariableData != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#GetCoreHrProcessFormVariableData mock enable")
		return r.cli.mock.mockCoreHrGetCoreHrProcessFormVariableData(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "GetCoreHrProcessFormVariableData",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/processes/:process_id/form_variable_data",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHrProcessFormVariableDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrGetCoreHrProcessFormVariableData mock CoreHrGetCoreHrProcessFormVariableData method
func (r *Mock) MockCoreHrGetCoreHrProcessFormVariableData(f func(ctx context.Context, request *GetCoreHrProcessFormVariableDataReq, options ...MethodOptionFunc) (*GetCoreHrProcessFormVariableDataResp, *Response, error)) {
	r.mockCoreHrGetCoreHrProcessFormVariableData = f
}

// UnMockCoreHrGetCoreHrProcessFormVariableData un-mock CoreHrGetCoreHrProcessFormVariableData method
func (r *Mock) UnMockCoreHrGetCoreHrProcessFormVariableData() {
	r.mockCoreHrGetCoreHrProcessFormVariableData = nil
}

// GetCoreHrProcessFormVariableDataReq ...
type GetCoreHrProcessFormVariableDataReq struct {
	ProcessID string `path:"process_id" json:"-"` // 流程ID, 示例值: "123456987"
}

// GetCoreHrProcessFormVariableDataResp ...
type GetCoreHrProcessFormVariableDataResp struct {
	FieldVariableValues []*GetCoreHrProcessFormVariableDataRespFieldVariableValue `json:"field_variable_values,omitempty"` // 流程变量
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValue struct {
	VariableApiName string                                                               `json:"variable_api_name,omitempty"` // 变量api名称
	VariableName    *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableName  `json:"variable_name,omitempty"`     // 变量名称的i18n描述
	VariableValue   *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValue `json:"variable_value,omitempty"`    // 变量值的对象
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableName ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableName struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValue struct {
	TextValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueTextValue       `json:"text_value,omitempty"`       // 文本变量对象
	NumberValue     *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueNumberValue     `json:"number_value,omitempty"`     // 数值变量对象
	DateValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueDateValue       `json:"date_value,omitempty"`       // 日期变量对象
	EmploymentValue *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEmploymentValue `json:"employment_value,omitempty"` // 员工变量对象
	DateTimeValue   *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueDateTimeValue   `json:"date_time_value,omitempty"`  // 日期时间变量对象
	EnumValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEnumValue       `json:"enum_value,omitempty"`       // 枚举变量对象
	NullValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueNullValue       `json:"null_value,omitempty"`       // 空变量对象
	BoolValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueBoolValue       `json:"bool_value,omitempty"`       // 布尔变量对象
	DepartmentValue *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueDepartmentValue `json:"department_value,omitempty"` // 部门变量对象
	FileValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueFileValue       `json:"file_value,omitempty"`       // 文件变量对象
	I18nValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueI18nValue       `json:"i18n_value,omitempty"`       // i18n变量对象
	ObjectValue     *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueObjectValue     `json:"object_value,omitempty"`     // 对象变量
	ListValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValue       `json:"list_value,omitempty"`       // 列表对象
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueBoolValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueBoolValue struct {
	Value bool `json:"value,omitempty"` // 布尔变量的值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueDateTimeValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueDateTimeValue struct {
	Value int64  `json:"value,omitempty"` // 毫秒的时间戳
	Zone  string `json:"zone,omitempty"`  // 时区
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueDateValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueDateValue struct {
	Value int64 `json:"value,omitempty"` // 日期变量的值, 从1970起的天数
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueDepartmentValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueDepartmentValue struct {
	Value string `json:"value,omitempty"` // 部门ID
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEmploymentValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEmploymentValue struct {
	Value  string `json:"value,omitempty"`   // employmentID
	UserID string `json:"user_id,omitempty"` // 员工ID 如3158117
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEnumValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEnumValue struct {
	Value string                                                                            `json:"value,omitempty"` // 枚举值
	Name  *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueName `json:"name,omitempty"`  // 枚举的名称
	Desc  *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueDesc `json:"desc,omitempty"`  // 枚举的描述
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueDesc ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueDesc struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueName ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueName struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueFileValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueFileValue struct {
	SourceType int64  `json:"source_type,omitempty"` // 文件源类型（1BPM; 2主数据）
	FileID     string `json:"file_id,omitempty"`     // 文件id
	FileName   string `json:"file_name,omitempty"`   // 文件名称
	Length     int64  `json:"length,omitempty"`      // 文件长度
	MimeType   string `json:"mime_type,omitempty"`   // mime type
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueI18nValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueI18nValue struct {
	Value *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueI18nValueValue `json:"value,omitempty"` // i18n值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueI18nValueValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueI18nValueValue struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValue struct {
	Values []*GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValue `json:"values,omitempty"` // 列表值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValue struct {
	TextValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueTextValue       `json:"text_value,omitempty"`       // 文本变量对象
	NumberValue     *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueNumberValue     `json:"number_value,omitempty"`     // 数值变量对象
	DateValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateValue       `json:"date_value,omitempty"`       // 日期变量对象
	EmploymentValue *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEmploymentValue `json:"employment_value,omitempty"` // 员工变量对象
	DateTimeValue   *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateTimeValue   `json:"date_time_value,omitempty"`  // 日期时间变量对象
	EnumValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValue       `json:"enum_value,omitempty"`       // 枚举变量对象
	NullValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueNullValue       `json:"null_value,omitempty"`       // 空变量对象
	BoolValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueBoolValue       `json:"bool_value,omitempty"`       // 布尔变量对象
	DepartmentValue *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDepartmentValue `json:"department_value,omitempty"` // 部门变量对象
	FileValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueFileValue       `json:"file_value,omitempty"`       // 文件变量对象
	I18nValue       *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValue       `json:"i18n_value,omitempty"`       // i18n变量对象
	ObjectValue     *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueObjectValue     `json:"object_value,omitempty"`     // 对象变量
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueBoolValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueBoolValue struct {
	Value bool `json:"value,omitempty"` // 布尔变量的值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateTimeValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateTimeValue struct {
	Value int64  `json:"value,omitempty"` // 毫秒的时间戳
	Zone  string `json:"zone,omitempty"`  // 时区
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateValue struct {
	Value int64 `json:"value,omitempty"` // 日期变量的值, 从1970起的天数
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDepartmentValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDepartmentValue struct {
	Value string `json:"value,omitempty"` // 部门ID
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEmploymentValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEmploymentValue struct {
	Value  string `json:"value,omitempty"`   // employmentID
	UserID string `json:"user_id,omitempty"` // 员工ID 如3158117
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValue struct {
	Value string                                                                                          `json:"value,omitempty"` // 枚举值
	Name  *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueName `json:"name,omitempty"`  // 枚举的名称
	Desc  *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueDesc `json:"desc,omitempty"`  // 枚举的描述
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueDesc ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueDesc struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueName ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueName struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueFileValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueFileValue struct {
	SourceType int64  `json:"source_type,omitempty"` // 文件源类型（1BPM; 2主数据）
	FileID     string `json:"file_id,omitempty"`     // 文件id
	FileName   string `json:"file_name,omitempty"`   // 文件名称
	Length     int64  `json:"length,omitempty"`      // 文件长度
	MimeType   string `json:"mime_type,omitempty"`   // mime type
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValue struct {
	Value *GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValueValue `json:"value,omitempty"` // i18n值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValueValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValueValue struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueNumberValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueNumberValue struct {
	Value string `json:"value,omitempty"` // 数值类型变量的值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueObjectValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueObjectValue struct {
	Value     string `json:"value,omitempty"`       // 对象ID
	WkApiName string `json:"wk_api_name,omitempty"` // 主数据apiName
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueTextValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueTextValue struct {
	Value string `json:"value,omitempty"` // 文本类型变量的值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueNumberValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueNumberValue struct {
	Value string `json:"value,omitempty"` // 数值类型变量的值
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueObjectValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueObjectValue struct {
	Value     string `json:"value,omitempty"`       // 对象ID
	WkApiName string `json:"wk_api_name,omitempty"` // 主数据apiName
}

// GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueTextValue ...
type GetCoreHrProcessFormVariableDataRespFieldVariableValueVariableValueTextValue struct {
	Value string `json:"value,omitempty"` // 文本类型变量的值
}

// getCoreHrProcessFormVariableDataResp ...
type getCoreHrProcessFormVariableDataResp struct {
	Code int64                                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHrProcessFormVariableDataResp `json:"data,omitempty"`
}
