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

// GetApproval 根据 Approval Code 获取某个审批定义的详情, 用于构造创建审批实例的请求。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/get
func (r *ApprovalService) GetApproval(ctx context.Context, request *GetApprovalReq, options ...MethodOptionFunc) (*GetApprovalResp, *Response, error) {
	if r.cli.mock.mockApprovalGetApproval != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#GetApproval mock enable")
		return r.cli.mock.mockApprovalGetApproval(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "GetApproval",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/approvals/:approval_code",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApprovalResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalGetApproval mock ApprovalGetApproval method
func (r *Mock) MockApprovalGetApproval(f func(ctx context.Context, request *GetApprovalReq, options ...MethodOptionFunc) (*GetApprovalResp, *Response, error)) {
	r.mockApprovalGetApproval = f
}

// UnMockApprovalGetApproval un-mock ApprovalGetApproval method
func (r *Mock) UnMockApprovalGetApproval() {
	r.mockApprovalGetApproval = nil
}

// GetApprovalReq ...
type GetApprovalReq struct {
	ApprovalCode string  `path:"approval_code" json:"-"` // 审批定义 Code, 示例值: "7C468A54-8745-2245-9675-08B7C63E7A85"
	Locale       *string `query:"locale" json:"-"`       // 语言可选值, 示例值: "zh-CN", 可选值有: zh-CN: 中文, en-US: 英文, ja-JP: 日文
}

// GetApprovalResp ...
type GetApprovalResp struct {
	ApprovalName string                   `json:"approval_name,omitempty"` // 审批名称
	Status       string                   `json:"status,omitempty"`        // 审批定义状态, 可选值有: ACTIVE: 已启用, INACTIVE: 已停用, DELETED: 已删除, UNKNOWN: 未知
	Form         ApprovalWidgetList       `json:"form,omitempty"`          // 控件信息, 见下方form字段说明
	NodeList     []*GetApprovalRespNode   `json:"node_list,omitempty"`     // 节点信息
	Viewers      []*GetApprovalRespViewer `json:"viewers,omitempty"`       // 可见人列表
}

// GetApprovalRespNode ...
type GetApprovalRespNode struct {
	Name                string                                    `json:"name,omitempty"`                  // 节点名称
	NeedApprover        bool                                      `json:"need_approver,omitempty"`         // 是否发起人自选节点 true - 发起审批时需要提交审批人
	NodeID              string                                    `json:"node_id,omitempty"`               // 节点 ID
	CustomNodeID        string                                    `json:"custom_node_id,omitempty"`        // 节点自定义 ID, 如果没有设置则不返回
	NodeType            string                                    `json:"node_type,omitempty"`             // 审批方式, 可选值有: AND: 会签, OR: 或签, SEQUENTIAL: 依次审批, CC_NODE: 抄送节点
	ApproverChosenMulti bool                                      `json:"approver_chosen_multi,omitempty"` // 是否支持多选: true-支持, 发起、结束节点该值无意义
	ApproverChosenRange []*GetApprovalRespNodeApproverChosenRange `json:"approver_chosen_range,omitempty"` // 自选范围
}

// GetApprovalRespNodeApproverChosenRange ...
type GetApprovalRespNodeApproverChosenRange struct {
	ApproverRangeType int64    `json:"approver_range_type,omitempty"` // 指定范围: 0-all, 1-指定角色, 2-指定人员, 可选值有: 0: 全公司范围, 1: 指定角色范围, 2: 指定用户范围
	ApproverRangeIDs  []string `json:"approver_range_ids,omitempty"`  // 根据上面的type, 分别存放角色id与userid, type为0时本字段为空列表
}

// GetApprovalRespViewer ...
type GetApprovalRespViewer struct {
	Type   string `json:"type,omitempty"`    // 可见人类型, 可选值有: TENANT: 租户内可见, DEPARTMENT: 指定部门, USER: 指定用户, ROLE: 指定角色, USER_GROUP: 指定用户组, NONE: 任何人都不可见
	ID     string `json:"id,omitempty"`      // 在可见人类型为DEPARTMENT时, id为部门的id ；在可见人类型为USER时, id为用户的id ；在可见人类型为ROLE时, id为角色的id ；在可见人类型为USER_GROUP时, id为用户组的id
	UserID string `json:"user_id,omitempty"` // 在可见人类型为USER时, 表示可见人用户id
}

// getApprovalResp ...
type getApprovalResp struct {
	Code int64            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *GetApprovalResp `json:"data,omitempty"`
}