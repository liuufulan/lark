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
	"io"
)

// UpdateACSUserFace 用户需要录入人脸图片才可以使用门禁考勤机。使用该 API 上传门禁用户的人脸图片。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user-face/update
func (r *ACSService) UpdateACSUserFace(ctx context.Context, request *UpdateACSUserFaceReq, options ...MethodOptionFunc) (*UpdateACSUserFaceResp, *Response, error) {
	if r.cli.mock.mockACSUpdateACSUserFace != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] ACS#UpdateACSUserFace mock enable")
		return r.cli.mock.mockACSUpdateACSUserFace(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "ACS",
		API:                   "UpdateACSUserFace",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/acs/v1/users/:user_id/face",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		IsFile:                true,
	}
	resp := new(updateACSUserFaceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockACSUpdateACSUserFace mock ACSUpdateACSUserFace method
func (r *Mock) MockACSUpdateACSUserFace(f func(ctx context.Context, request *UpdateACSUserFaceReq, options ...MethodOptionFunc) (*UpdateACSUserFaceResp, *Response, error)) {
	r.mockACSUpdateACSUserFace = f
}

// UnMockACSUpdateACSUserFace un-mock ACSUpdateACSUserFace method
func (r *Mock) UnMockACSUpdateACSUserFace() {
	r.mockACSUpdateACSUserFace = nil
}

// UpdateACSUserFaceReq ...
type UpdateACSUserFaceReq struct {
	UserID     string    `path:"user_id" json:"-"`       // 用户 ID, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	UserIDType *IDType   `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Files      io.Reader `json:"files,omitempty"`        // 人脸图片内容, 示例值: jpg图片
	FileType   FileType  `json:"file_type,omitempty"`    // 文件类型, 可选的类型有jpg, png, 示例值: "jpg"
	FileName   string    `json:"file_name,omitempty"`    // 带后缀的文件名, 示例值: "efeqz12f.jpg"
}

// UpdateACSUserFaceResp ...
type UpdateACSUserFaceResp struct {
}

// updateACSUserFaceResp ...
type updateACSUserFaceResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *UpdateACSUserFaceResp `json:"data,omitempty"`
}