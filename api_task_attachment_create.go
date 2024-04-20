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

// CreateTaskAttachment 为特定资源上传附件。本接口可以支持一次上传多个附件, 最多5个。每个附件尺寸不超过50MB, 格式不限。
//
// 上传请求体的格式为"form-data"。若希望上传多个附件, 则提供多个"file"字段即可。返回的附件顺序将会与输入的file顺序保持一致。
// 目前资源类型仅支持"task", `resource_id`需要填写任务的GUID。
// 为任务上传附件需要任务的可编辑权限
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/attachment/upload
func (r *TaskService) CreateTaskAttachment(ctx context.Context, request *CreateTaskAttachmentReq, options ...MethodOptionFunc) (*CreateTaskAttachmentResp, *Response, error) {
	if r.cli.mock.mockTaskCreateTaskAttachment != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#CreateTaskAttachment mock enable")
		return r.cli.mock.mockTaskCreateTaskAttachment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CreateTaskAttachment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/attachments/upload",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createTaskAttachmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskCreateTaskAttachment mock TaskCreateTaskAttachment method
func (r *Mock) MockTaskCreateTaskAttachment(f func(ctx context.Context, request *CreateTaskAttachmentReq, options ...MethodOptionFunc) (*CreateTaskAttachmentResp, *Response, error)) {
	r.mockTaskCreateTaskAttachment = f
}

// UnMockTaskCreateTaskAttachment un-mock TaskCreateTaskAttachment method
func (r *Mock) UnMockTaskCreateTaskAttachment() {
	r.mockTaskCreateTaskAttachment = nil
}

// CreateTaskAttachmentReq ...
type CreateTaskAttachmentReq struct {
	UserIDType   *IDType   `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ResourceType *string   `json:"resource_type,omitempty"` // 附件归属资源的类型, 示例值: "task", 默认值: `task`
	ResourceID   string    `json:"resource_id,omitempty"`   // 附件要归属资源的id。例如, 要给任务添加附件, 这里要填入任务GUID。任务GUID可以通过[任务相关接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/overview)获得, 示例值: "fe96108d-b004-4a47-b2f8-6886e758b3a5", 最大长度: `100` 字符
	File         io.Reader `json:"file,omitempty"`          // 要上传的文件, 单请求支持最多5个文件。上传结果的顺序将和请求中文件的顺序保持一致, 示例值: file binary
}

// CreateTaskAttachmentResp ...
type CreateTaskAttachmentResp struct {
	Items []*CreateTaskAttachmentRespItem `json:"items,omitempty"` // 上传的附件列表
}

// CreateTaskAttachmentRespItem ...
type CreateTaskAttachmentRespItem struct {
	Guid       string                                `json:"guid,omitempty"`        // 附件GUID
	FileToken  string                                `json:"file_token,omitempty"`  // 附件在云文档系统中的token
	Name       string                                `json:"name,omitempty"`        // 附件名
	Size       int64                                 `json:"size,omitempty"`        // 附件的字节大小
	Resource   *CreateTaskAttachmentRespItemResource `json:"resource,omitempty"`    // 附件归属的资源
	Uploader   *CreateTaskAttachmentRespItemUploader `json:"uploader,omitempty"`    // 附件上传者
	IsCover    bool                                  `json:"is_cover,omitempty"`    // 是否是封面图
	UploadedAt string                                `json:"uploaded_at,omitempty"` // 上传时间戳(ms)
}

// CreateTaskAttachmentRespItemResource ...
type CreateTaskAttachmentRespItemResource struct {
	Type string `json:"type,omitempty"` // 资源类型
	ID   string `json:"id,omitempty"`   // 资源ID
}

// CreateTaskAttachmentRespItemUploader ...
type CreateTaskAttachmentRespItemUploader struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// createTaskAttachmentResp ...
type createTaskAttachmentResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *CreateTaskAttachmentResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}