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

// ExtractLingoEntity 提取文本中可能成为词条的词语, 且不会过滤已经成为词条的词语。同时返回推荐的别名。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/lingo-v1/entity/extract
//
// Deprecated
func (r *LingoService) ExtractLingoEntity(ctx context.Context, request *ExtractLingoEntityReq, options ...MethodOptionFunc) (*ExtractLingoEntityResp, *Response, error) {
	if r.cli.mock.mockLingoExtractLingoEntity != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Lingo#ExtractLingoEntity mock enable")
		return r.cli.mock.mockLingoExtractLingoEntity(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Lingo",
		API:                   "ExtractLingoEntity",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/lingo/v1/entities/extract",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(extractLingoEntityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockLingoExtractLingoEntity mock LingoExtractLingoEntity method
func (r *Mock) MockLingoExtractLingoEntity(f func(ctx context.Context, request *ExtractLingoEntityReq, options ...MethodOptionFunc) (*ExtractLingoEntityResp, *Response, error)) {
	r.mockLingoExtractLingoEntity = f
}

// UnMockLingoExtractLingoEntity un-mock LingoExtractLingoEntity method
func (r *Mock) UnMockLingoExtractLingoEntity() {
	r.mockLingoExtractLingoEntity = nil
}

// ExtractLingoEntityReq ...
type ExtractLingoEntityReq struct {
	Text *string `json:"text,omitempty"` // 需要被提取词条的文本（不会过滤租户中已成为词条的内容）, 示例值: "飞书词典是一部高效汇聚企业内各类信息, 并可由企业成员参与编辑的在线词典", 最大长度: `128` 字符
}

// ExtractLingoEntityResp ...
type ExtractLingoEntityResp struct {
	EntityWord []*ExtractLingoEntityRespEntityWord `json:"entity_word,omitempty"` // 文本中可能的成为百科词条的实体词
}

// ExtractLingoEntityRespEntityWord ...
type ExtractLingoEntityRespEntityWord struct {
	Name    string   `json:"name,omitempty"`    // 抽取出的词条名
	Aliases []string `json:"aliases,omitempty"` // 词条可能的推荐别名
}

// extractLingoEntityResp ...
type extractLingoEntityResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *ExtractLingoEntityResp `json:"data,omitempty"`
}