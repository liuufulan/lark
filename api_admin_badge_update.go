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

// UpdateAdminBadge 通过该接口可以修改勋章的信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/badge/update
func (r *AdminService) UpdateAdminBadge(ctx context.Context, request *UpdateAdminBadgeReq, options ...MethodOptionFunc) (*UpdateAdminBadgeResp, *Response, error) {
	if r.cli.mock.mockAdminUpdateAdminBadge != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Admin#UpdateAdminBadge mock enable")
		return r.cli.mock.mockAdminUpdateAdminBadge(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Admin",
		API:                   "UpdateAdminBadge",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/admin/v1/badges/:badge_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateAdminBadgeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAdminUpdateAdminBadge mock AdminUpdateAdminBadge method
func (r *Mock) MockAdminUpdateAdminBadge(f func(ctx context.Context, request *UpdateAdminBadgeReq, options ...MethodOptionFunc) (*UpdateAdminBadgeResp, *Response, error)) {
	r.mockAdminUpdateAdminBadge = f
}

// UnMockAdminUpdateAdminBadge un-mock AdminUpdateAdminBadge method
func (r *Mock) UnMockAdminUpdateAdminBadge() {
	r.mockAdminUpdateAdminBadge = nil
}

// UpdateAdminBadgeReq ...
type UpdateAdminBadgeReq struct {
	BadgeID     string  `path:"badge_id" json:"-"`      // 勋章ID, 示例值: "m_MzfKDM", 长度范围: `1` ～ `64` 字符
	Name        string  `json:"name,omitempty"`         // 租户内唯一的勋章名称, 最多30个字符, 示例值: "激励勋章"
	Explanation *string `json:"explanation,omitempty"`  // 勋章的描述文案, 最多100个字符, 示例值: "这枚勋章为了激励员工颁发。"
	DetailImage string  `json:"detail_image,omitempty"` // 企业勋章的详情图Key。1.权限校验: 非本租户上传的图片key, 不能直接使用；2.时效校验: 创建勋章, 或者修改勋章图片key时, 需使用1h内上传的图片key, 示例值: "75a1949f-d9df-4b46-bc88-dacc51e88f3j", 最小长度: `1` 字符
	ShowImage   string  `json:"show_image,omitempty"`   // 企业勋章的头像挂饰图Key。1.权限校验: 非本租户上传的图片key, 不能直接使用；2.时效校验: 创建勋章, 或者修改勋章图片key时, 需使用1h内上传的图片key, 示例值: "03daa74a-159f-49e9-963e-b6c4d76103fj", 最小长度: `1` 字符
}

// UpdateAdminBadgeResp ...
type UpdateAdminBadgeResp struct {
	Badge *UpdateAdminBadgeRespBadge `json:"badge,omitempty"` // 勋章信息
}

// UpdateAdminBadgeRespBadge ...
type UpdateAdminBadgeRespBadge struct {
	ID          string `json:"id,omitempty"`           // 租户内勋章的唯一标识, 该值由系统随机生成。
	Name        string `json:"name,omitempty"`         // 租户内唯一的勋章名称, 最多30个字符。
	Explanation string `json:"explanation,omitempty"`  // 勋章的描述文案, 最多100个字符。
	DetailImage string `json:"detail_image,omitempty"` // 企业勋章的详情图Key。1.权限校验: 非本租户上传的图片key, 不能直接使用；2.时效校验: 创建勋章, 或者修改勋章图片key时, 需使用1h内上传的图片key。
	ShowImage   string `json:"show_image,omitempty"`   // 企业勋章的头像挂饰图Key。1.权限校验: 非本租户上传的图片key, 不能直接使用；2.时效校验: 创建勋章, 或者修改勋章图片key时, 需使用1h内上传的图片key。
}

// updateAdminBadgeResp ...
type updateAdminBadgeResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *UpdateAdminBadgeResp `json:"data,omitempty"`
}
