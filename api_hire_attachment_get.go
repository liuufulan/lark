// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetHireAttachment 获取招聘系统中附件的元信息，比如文件名、创建时间、文件url等
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/attachment/get
func (r *HireService) GetHireAttachment(ctx context.Context, request *GetHireAttachmentReq, options ...MethodOptionFunc) (*GetHireAttachmentResp, *Response, error) {
	if r.cli.mock.mockHireGetHireAttachment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireAttachment mock enable")
		return r.cli.mock.mockHireGetHireAttachment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireAttachment",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/hire/v1/attachments/:attachment_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireAttachmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHireGetHireAttachment(f func(ctx context.Context, request *GetHireAttachmentReq, options ...MethodOptionFunc) (*GetHireAttachmentResp, *Response, error)) {
	r.mockHireGetHireAttachment = f
}

func (r *Mock) UnMockHireGetHireAttachment() {
	r.mockHireGetHireAttachment = nil
}

type GetHireAttachmentReq struct {
	AttachmentID string `path:"attachment_id" json:"-"` // 附件id, 示例值："6435242341238"
}

type getHireAttachmentResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetHireAttachmentResp `json:"data,omitempty"`
}

type GetHireAttachmentResp struct {
	Attachment *GetHireAttachmentRespAttachment `json:"attachment,omitempty"` // 附件信息
}

type GetHireAttachmentRespAttachment struct {
	ID         string `json:"id,omitempty"`          // 附件id
	URL        string `json:"url,omitempty"`         // 附件的url
	Name       string `json:"name,omitempty"`        // 附件文件名
	Mime       string `json:"mime,omitempty"`        // 媒体类型/MIME
	CreateTime int64  `json:"create_time,omitempty"` // 附件创建时间（单位ms）
}