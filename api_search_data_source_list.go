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

// GetSearchDataSourceList 批量获取创建的数据源信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/list
func (r *SearchService) GetSearchDataSourceList(ctx context.Context, request *GetSearchDataSourceListReq, options ...MethodOptionFunc) (*GetSearchDataSourceListResp, *Response, error) {
	if r.cli.mock.mockSearchGetSearchDataSourceList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#GetSearchDataSourceList mock enable")
		return r.cli.mock.mockSearchGetSearchDataSourceList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "GetSearchDataSourceList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/data_sources",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getSearchDataSourceListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchGetSearchDataSourceList mock SearchGetSearchDataSourceList method
func (r *Mock) MockSearchGetSearchDataSourceList(f func(ctx context.Context, request *GetSearchDataSourceListReq, options ...MethodOptionFunc) (*GetSearchDataSourceListResp, *Response, error)) {
	r.mockSearchGetSearchDataSourceList = f
}

// UnMockSearchGetSearchDataSourceList un-mock SearchGetSearchDataSourceList method
func (r *Mock) UnMockSearchGetSearchDataSourceList() {
	r.mockSearchGetSearchDataSourceList = nil
}

// GetSearchDataSourceListReq ...
type GetSearchDataSourceListReq struct {
	View      *int64  `query:"view" json:"-"`       // 回包数据格式, 0-全量数据；1-摘要数据, 注: 摘要数据仅包含"id", "name", "state", 示例值: 0, 可选值有: 0: 全量数据, 1: 摘要数据
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 最大值: `50`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "PxZFma9OIRhdBlT/dOYNiu2Ro8F2WAhcby7OhOijfljZ"
}

// GetSearchDataSourceListResp ...
type GetSearchDataSourceListResp struct {
	HasMore   bool                               `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                             `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetSearchDataSourceListRespItem `json:"items,omitempty"`      // 数据源中的数据记录
}

// GetSearchDataSourceListRespItem ...
type GetSearchDataSourceListRespItem struct {
	ID               string                                          `json:"id,omitempty"`                // 数据源的唯一标识
	Name             string                                          `json:"name,omitempty"`              // data_source的展示名称
	State            int64                                           `json:"state,omitempty"`             // 数据源状态, 0-已上线, 1-未上线。如果未填, 默认是未上线状态, 可选值有: 0: 已上线, 1: 未上线
	Description      string                                          `json:"description,omitempty"`       // 对于数据源的描述
	CreateTime       string                                          `json:"create_time,omitempty"`       // 创建时间, 使用Unix时间戳, 单位为“秒”
	UpdateTime       string                                          `json:"update_time,omitempty"`       // 更新时间, 使用Unix时间戳, 单位为“秒”
	IsExceedQuota    bool                                            `json:"is_exceed_quota,omitempty"`   // 是否超限
	IconURL          string                                          `json:"icon_url,omitempty"`          // 数据源在 search tab 上的展示图标路径
	Template         string                                          `json:"template,omitempty"`          // 数据源采用的展示模版名称
	SearchableFields []string                                        `json:"searchable_fields,omitempty"` // 【已废弃, 如有定制需要请使用“数据范式”接口】描述哪些字段可以被搜索
	I18nName         *GetSearchDataSourceListRespItemI18nName        `json:"i18n_name,omitempty"`         // 数据源的国际化展示名称
	I18nDescription  *GetSearchDataSourceListRespItemI18nDescription `json:"i18n_description,omitempty"`  // 数据源的国际化描述
	SchemaID         string                                          `json:"schema_id,omitempty"`         // 数据源关联的 schema 标识
}

// GetSearchDataSourceListRespItemI18nDescription ...
type GetSearchDataSourceListRespItemI18nDescription struct {
	ZhCn string `json:"zh_cn,omitempty"` // 国际化字段: 中文
	EnUs string `json:"en_us,omitempty"` // 国际化字段: 英文
	JaJp string `json:"ja_jp,omitempty"` // 国际化字段: 日文
}

// GetSearchDataSourceListRespItemI18nName ...
type GetSearchDataSourceListRespItemI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 国际化字段: 中文
	EnUs string `json:"en_us,omitempty"` // 国际化字段: 英文
	JaJp string `json:"ja_jp,omitempty"` // 国际化字段: 日文
}

// getSearchDataSourceListResp ...
type getSearchDataSourceListResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *GetSearchDataSourceListResp `json:"data,omitempty"`
}
