package creative

import (
	"github.com/xiaoshouchen/kwai-marketing-api/core"
	"github.com/xiaoshouchen/kwai-marketing-api/model/creative"
)

// ActionBarTextList 获取行动号召按钮
func ActionBarTextList(clt *core.SDKClient, accessToken string, req *creative.ActionBarTextListRequest) ([]string, error) {
	var resp creative.ActionBarTextListResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.ActionBarText, nil
}
