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

// GetHireTalentList 根据更新时间获取人才列表, 仅支持获取默认字段信息, 获取详细信息可调用「获取人才详细」接口。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/talent/list
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/candidate-management/talent/list
func (r *HireService) GetHireTalentList(ctx context.Context, request *GetHireTalentListReq, options ...MethodOptionFunc) (*GetHireTalentListResp, *Response, error) {
	if r.cli.mock.mockHireGetHireTalentList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#GetHireTalentList mock enable")
		return r.cli.mock.mockHireGetHireTalentList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireTalentList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/talents",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireTalentListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireTalentList mock HireGetHireTalentList method
func (r *Mock) MockHireGetHireTalentList(f func(ctx context.Context, request *GetHireTalentListReq, options ...MethodOptionFunc) (*GetHireTalentListResp, *Response, error)) {
	r.mockHireGetHireTalentList = f
}

// UnMockHireGetHireTalentList un-mock HireGetHireTalentList method
func (r *Mock) UnMockHireGetHireTalentList() {
	r.mockHireGetHireTalentList = nil
}

// GetHireTalentListReq ...
type GetHireTalentListReq struct {
	UpdateStartTime *string `query:"update_start_time" json:"-"` // 最早更新时间, 毫秒级时间戳, 示例值: 1618500278663
	UpdateEndTime   *string `query:"update_end_time" json:"-"`   // 最晚更新时间, 毫秒级时间戳, 示例值: 1618500278663
	PageSize        *int64  `query:"page_size" json:"-"`         // 分页大小, 示例值: 10, 最大值: `20`
	PageToken       *string `query:"page_token" json:"-"`        // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: eyJvZmZzZXQiOjEwLCJ0aW1lc3RhbXAiOjE2Mjc1NTUyMjM2NzIsImlkIjpudWxsfQ[
	UserIDType      *IDType `query:"user_id_type" json:"-"`      // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以 people_admin_id 来识别用户, 默认值: `people_admin_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	QueryOption     *string `query:"query_option" json:"-"`      // 请求控制参数, 示例值: ignore_empty_error, 可选值有: ignore_empty_error: 忽略结果为空时的报错
}

// GetHireTalentListResp ...
type GetHireTalentListResp struct {
	HasMore   bool                         `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                       `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetHireTalentListRespItem `json:"items,omitempty"`      // 列表
}

// GetHireTalentListRespItem ...
type GetHireTalentListRespItem struct {
	ID                        string                                            `json:"id,omitempty"`                          // 人才ID
	IsInAgencyPeriod          bool                                              `json:"is_in_agency_period,omitempty"`         // 是否在猎头保护期, 可选值有: false: 未在猎头保护期, true: 在猎头保护期
	IsOnboarded               bool                                              `json:"is_onboarded,omitempty"`                // 是否已入职, 可选值有: false: 未入职, true: 已入职
	BasicInfo                 *GetHireTalentListRespItemBasicInfo               `json:"basic_info,omitempty"`                  // 基础信息
	EducationList             []*GetHireTalentListRespItemEducation             `json:"education_list,omitempty"`              // 教育经历
	CareerList                []*GetHireTalentListRespItemCareer                `json:"career_list,omitempty"`                 // 工作经历
	ProjectList               []*GetHireTalentListRespItemProject               `json:"project_list,omitempty"`                // 项目经历
	WorksList                 []*GetHireTalentListRespItemWorks                 `json:"works_list,omitempty"`                  // 作品
	AwardList                 []*GetHireTalentListRespItemAward                 `json:"award_list,omitempty"`                  // 获奖
	LanguageList              []*GetHireTalentListRespItemLanguage              `json:"language_list,omitempty"`               // 语言能力
	SnsList                   []*GetHireTalentListRespItemSns                   `json:"sns_list,omitempty"`                    // 社交账号
	ResumeSourceList          []*GetHireTalentListRespItemResumeSource          `json:"resume_source_list,omitempty"`          // 简历来源
	InterviewRegistrationList []*GetHireTalentListRespItemInterviewRegistration `json:"interview_registration_list,omitempty"` // 面试登记表
	ResumeAttachmentIDList    []string                                          `json:"resume_attachment_id_list,omitempty"`   // 简历附件id列表（按照简历创建时间降序）
	TopDegree                 int64                                             `json:"top_degree,omitempty"`                  // 最高学历, 可选值有: 1: 小学, 2: 初中, 3: 专职, 4: 高中, 5: 大专, 6: 本科, 7: 硕士, 8: 博士, 9: 其他
	FirstDegree               int64                                             `json:"first_degree,omitempty"`                // 第一学历, 可选值有: 1: 低于大专, 2: 大专, 3: 本科, 4: 硕士, 5: 博士, 6: 其他, 7: 无
}

// GetHireTalentListRespItemAward ...
type GetHireTalentListRespItemAward struct {
	ID        string `json:"id,omitempty"`         // ID
	Title     string `json:"title,omitempty"`      // 获奖名称
	AwardTime string `json:"award_time,omitempty"` // 获奖时间
	Desc      string `json:"desc,omitempty"`       // 描述
}

// GetHireTalentListRespItemBasicInfo ...
type GetHireTalentListRespItemBasicInfo struct {
	Name                 string                                             `json:"name,omitempty"`                  // 名字
	Mobile               string                                             `json:"mobile,omitempty"`                // 手机
	MobileCode           string                                             `json:"mobile_code,omitempty"`           // 手机国家区号
	MobileCountryCode    string                                             `json:"mobile_country_code,omitempty"`   // 手机国家代码
	Email                string                                             `json:"email,omitempty"`                 // 邮箱
	ExperienceYears      int64                                              `json:"experience_years,omitempty"`      // 工作年限
	Age                  int64                                              `json:"age,omitempty"`                   // 年龄
	Nationality          *GetHireTalentListRespItemBasicInfoNationality     `json:"nationality,omitempty"`           // 国籍
	Gender               int64                                              `json:"gender,omitempty"`                // 性别, 可选值有: 1: 男, 2: 女, 3: 其他
	CurrentCity          *GetHireTalentListRespItemBasicInfoCurrentCity     `json:"current_city,omitempty"`          // 所在地点
	HometownCity         *GetHireTalentListRespItemBasicInfoHometownCity    `json:"hometown_city,omitempty"`         // 家乡
	PreferredCityList    []*GetHireTalentListRespItemBasicInfoPreferredCity `json:"preferred_city_list,omitempty"`   // 意向地点
	IdentificationType   int64                                              `json:"identification_type,omitempty"`   // 证件类型, 可选值有: 1: 中国 - 居民身份证, 2: 护照, 3: 中国 - 港澳居民居住证, 4: 中国 - 台湾居民来往大陆通行证, 5: 其他, 6: 中国 - 港澳居民来往内地通行证, 9: 中国 - 台湾居民居住证
	IdentificationNumber string                                             `json:"identification_number,omitempty"` // 证件号
	Birthday             int64                                              `json:"birthday,omitempty"`              // 生日
	CreatorID            string                                             `json:"creator_id,omitempty"`            // 创建人
	MaritalStatus        int64                                              `json:"marital_status,omitempty"`        // 婚姻状况, 可选值有: 1: 已婚, 2: 未婚
	CurrentHomeAddress   string                                             `json:"current_home_address,omitempty"`  // 家庭住址
	ModifyTime           string                                             `json:"modify_time,omitempty"`           // 修改时间
}

// GetHireTalentListRespItemBasicInfoCurrentCity ...
type GetHireTalentListRespItemBasicInfoCurrentCity struct {
	CityCode string `json:"city_code,omitempty"` // 城市码
	ZhName   string `json:"zh_name,omitempty"`   // 中文名
	EnName   string `json:"en_name,omitempty"`   // 英文名
}

// GetHireTalentListRespItemBasicInfoHometownCity ...
type GetHireTalentListRespItemBasicInfoHometownCity struct {
	CityCode string `json:"city_code,omitempty"` // 城市码
	ZhName   string `json:"zh_name,omitempty"`   // 中文名
	EnName   string `json:"en_name,omitempty"`   // 英文名
}

// GetHireTalentListRespItemBasicInfoNationality ...
type GetHireTalentListRespItemBasicInfoNationality struct {
	NationalityCode string `json:"nationality_code,omitempty"` // 国家编码
	ZhName          string `json:"zh_name,omitempty"`          // 中文名
	EnName          string `json:"en_name,omitempty"`          // 英文名
}

// GetHireTalentListRespItemBasicInfoPreferredCity ...
type GetHireTalentListRespItemBasicInfoPreferredCity struct {
	CityCode string `json:"city_code,omitempty"` // 城市码
	ZhName   string `json:"zh_name,omitempty"`   // 中文名
	EnName   string `json:"en_name,omitempty"`   // 英文名
}

// GetHireTalentListRespItemCareer ...
type GetHireTalentListRespItemCareer struct {
	ID         string  `json:"id,omitempty"`          // ID
	Company    string  `json:"company,omitempty"`     // 公司名称
	Title      string  `json:"title,omitempty"`       // 职位名称
	Desc       string  `json:"desc,omitempty"`        // 描述
	StartTime  string  `json:"start_time,omitempty"`  // 开始时间
	EndTime    string  `json:"end_time,omitempty"`    // 结束时间
	CareerType int64   `json:"career_type,omitempty"` // 经历类型, 可选值有: 1: 实习经历, 2: 工作经历
	TagList    []int64 `json:"tag_list,omitempty"`    // 工作经历标签, 可选值有: 5: 百度 阿里 腾讯, 6: 头条, 美团, 滴滴, 7: 其它大厂
}

// GetHireTalentListRespItemEducation ...
type GetHireTalentListRespItemEducation struct {
	ID              string  `json:"id,omitempty"`               // ID
	Degree          int64   `json:"degree,omitempty"`           // 学位, 可选值有: 1: 小学, 2: 初中, 3: 专职, 4: 高中, 5: 大专, 6: 本科, 7: 硕士, 8: 博士, 9: 其他
	School          string  `json:"school,omitempty"`           // 学校
	FieldOfStudy    string  `json:"field_of_study,omitempty"`   // 专业
	StartTime       string  `json:"start_time,omitempty"`       // 开始时间
	EndTime         string  `json:"end_time,omitempty"`         // 结束时间
	EducationType   int64   `json:"education_type,omitempty"`   // 学历类型, 可选值有: 1: 海外及港台, 2: 统招全日制, 3: 非全日制, 4: 自考, 5: 其他
	AcademicRanking int64   `json:"academic_ranking,omitempty"` // 成绩排名, 可选值有: 5: 前 5 %, 10: 前 10 %, 20: 前 20 %, 30: 前 30 %, 50: 前 50 %, -1: 其他
	TagList         []int64 `json:"tag_list,omitempty"`         // 教育经历标签, 可选值有: 1: 985学校, 2: 211学校, 3: 一本, 4: 国外院校QS200
}

// GetHireTalentListRespItemInterviewRegistration ...
type GetHireTalentListRespItemInterviewRegistration struct {
	ID               string `json:"id,omitempty"`                // ID
	RegistrationTime int64  `json:"registration_time,omitempty"` // 创建时间
}

// GetHireTalentListRespItemLanguage ...
type GetHireTalentListRespItemLanguage struct {
	ID          string `json:"id,omitempty"`          // ID
	Language    int64  `json:"language,omitempty"`    // 语言, 可选值有: 1: 英语, 2: 法语, 3: 日语, 4: 韩语, 5: 德语, 6: 俄语, 7: 西班牙语, 8: 葡萄牙语, 9: 阿拉伯语, 10: 印地语, 11: 印度斯坦语, 12: 孟加拉语, 13: 豪萨语, 14: 旁遮普语, 15: 波斯语, 16: 斯瓦西里语, 17: 泰卢固语, 18: 土耳其语, 19: 意大利语, 20: 爪哇语, 21: 泰米尔语, 22: 马拉地语, 23: 越南语, 24: 普通话, 25: 粤语
	Proficiency int64  `json:"proficiency,omitempty"` // 精通程度, 可选值有: 1: 入门, 2: 日常会话, 3: 商务会话, 4: 无障碍沟通, 5: 母语
}

// GetHireTalentListRespItemProject ...
type GetHireTalentListRespItemProject struct {
	ID        string `json:"id,omitempty"`         // ID
	Name      string `json:"name,omitempty"`       // 项目名称
	Role      string `json:"role,omitempty"`       // 项目角色
	Link      string `json:"link,omitempty"`       // 项目链接
	Desc      string `json:"desc,omitempty"`       // 描述
	StartTime string `json:"start_time,omitempty"` // 开始时间
	EndTime   string `json:"end_time,omitempty"`   // 结束时间
}

// GetHireTalentListRespItemResumeSource ...
type GetHireTalentListRespItemResumeSource struct {
	ID     string `json:"id,omitempty"`      // ID
	ZhName string `json:"zh_name,omitempty"` // 中文名
	EnName string `json:"en_name,omitempty"` // 英文名
}

// GetHireTalentListRespItemSns ...
type GetHireTalentListRespItemSns struct {
	ID      string `json:"id,omitempty"`       // ID
	SnsType int64  `json:"sns_type,omitempty"` // SNS名称, 可选值有: 1: 领英, 2: 脉脉, 3: 微信, 4: 微博, 5: Github, 6: 知乎, 7: 脸书, 8: 推特, 9: Whatsapp, 10: 个人网站, 11: QQ
	Link    string `json:"link,omitempty"`     // URL/ID
}

// GetHireTalentListRespItemWorks ...
type GetHireTalentListRespItemWorks struct {
	ID   string `json:"id,omitempty"`   // ID
	Link string `json:"link,omitempty"` // 作品链接
	Desc string `json:"desc,omitempty"` // 描述
	Name string `json:"name,omitempty"` // 作品附件名称, 若需获取作品附件预览信息可调用「获取附件预览信息」接口
}

// getHireTalentListResp ...
type getHireTalentListResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *GetHireTalentListResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
