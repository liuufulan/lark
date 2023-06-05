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

// DownloadCoreHrPersonFile 根据ID下载文件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/file/get
func (r *CoreHrService) DownloadCoreHrPersonFile(ctx context.Context, request *DownloadCoreHrPersonFileReq, options ...MethodOptionFunc) (*DownloadCoreHrPersonFileResp, *Response, error) {
	if r.cli.mock.mockCoreHrDownloadCoreHrPersonFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#DownloadCoreHrPersonFile mock enable")
		return r.cli.mock.mockCoreHrDownloadCoreHrPersonFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "DownloadCoreHrPersonFile",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/files/:id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(downloadCoreHrPersonFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrDownloadCoreHrPersonFile mock CoreHrDownloadCoreHrPersonFile method
func (r *Mock) MockCoreHrDownloadCoreHrPersonFile(f func(ctx context.Context, request *DownloadCoreHrPersonFileReq, options ...MethodOptionFunc) (*DownloadCoreHrPersonFileResp, *Response, error)) {
	r.mockCoreHrDownloadCoreHrPersonFile = f
}

// UnMockCoreHrDownloadCoreHrPersonFile un-mock CoreHrDownloadCoreHrPersonFile method
func (r *Mock) UnMockCoreHrDownloadCoreHrPersonFile() {
	r.mockCoreHrDownloadCoreHrPersonFile = nil
}

// DownloadCoreHrPersonFileReq ...
type DownloadCoreHrPersonFileReq struct {
	ID string `path:"id" json:"-"` // 上传文件ID, 示例值: "150018109586e8ea745e47ae8feb3722dbe1d03a181336393633393133303431393831343930373235150200"
}

// downloadCoreHrPersonFileResp ...
type downloadCoreHrPersonFileResp struct {
	Code int64                         `json:"code,omitempty"`
	Msg  string                        `json:"msg,omitempty"`
	Data *DownloadCoreHrPersonFileResp `json:"data,omitempty"`
}

func (r *downloadCoreHrPersonFileResp) SetReader(file io.Reader) {
	if r.Data == nil {
		r.Data = &DownloadCoreHrPersonFileResp{}
	}
	r.Data.File = file
}

// DownloadCoreHrPersonFileResp ...
type DownloadCoreHrPersonFileResp struct {
	File io.Reader `json:"file,omitempty"`
}
