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

// TranslateText 机器翻译 (MT), 支持以下语种互译:
//
// "zh": 汉语；
// "zh-Hant": 繁体汉语；
// "en": 英语；
// "ja": 日语；
// "ru": 俄语；
// "de": 德语；
// "fr": 法语；
// "it": 意大利语；
// "pl": 波兰语；
// "th": 泰语；
// "hi": 印地语；
// "id": 印尼语；
// "es": 西班牙语；
// "pt": 葡萄牙语；
// "ko": 朝鲜语；
// "vi": 越南语；
// 单租户限流: 20QPS, 同租户下的应用没有限流, 共享本租户的 20QPS 限流。免费版不支持调用。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/translation-v1/text/translate
// new doc: https://open.feishu.cn/document/server-docs/ai/translation-v1/translate
func (r *AIService) TranslateText(ctx context.Context, request *TranslateTextReq, options ...MethodOptionFunc) (*TranslateTextResp, *Response, error) {
	if r.cli.mock.mockAITranslateText != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#TranslateText mock enable")
		return r.cli.mock.mockAITranslateText(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "TranslateText",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/translation/v1/text/translate",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(translateTextResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAITranslateText mock AITranslateText method
func (r *Mock) MockAITranslateText(f func(ctx context.Context, request *TranslateTextReq, options ...MethodOptionFunc) (*TranslateTextResp, *Response, error)) {
	r.mockAITranslateText = f
}

// UnMockAITranslateText un-mock AITranslateText method
func (r *Mock) UnMockAITranslateText() {
	r.mockAITranslateText = nil
}

// TranslateTextReq ...
type TranslateTextReq struct {
	SourceLanguage string                      `json:"source_language,omitempty"` // 源语言, 示例值: "zh"
	Text           string                      `json:"text,omitempty"`            // 源文本, 示例值: "尝试使用一下飞书吧"
	TargetLanguage string                      `json:"target_language,omitempty"` // 目标语言, 示例值: "en"
	Glossary       []*TranslateTextReqGlossary `json:"glossary,omitempty"`        // 请求级术语表, 携带术语, 仅在本次翻译中生效（最多能携带 128个术语词）
}

// TranslateTextReqGlossary ...
type TranslateTextReqGlossary struct {
	From string `json:"from,omitempty"` // 原文, 示例值: "飞书"
	To   string `json:"to,omitempty"`   // 译文, 示例值: "Lark"
}

// TranslateTextResp ...
type TranslateTextResp struct {
	Text string `json:"text,omitempty"` // 翻译后的文本
}

// translateTextResp ...
type translateTextResp struct {
	Code int64              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *TranslateTextResp `json:"data,omitempty"`
}
