package page

import (
	"github.com/xiaoshouchen/kwai-marketing-api/core"
	"github.com/xiaoshouchen/kwai-marketing-api/model/page"
)

// List 获取魔力建站落地页组信息列表
func List(clt *core.SDKClient, accessToken string, req *page.ListRequest) (*page.ListResponse, error) {
	var resp page.ListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
