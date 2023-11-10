package unit

import (
	"github.com/xiaoshouchen/kwai-marketing-api/core"
	"github.com/xiaoshouchen/kwai-marketing-api/model/unit"
)

// UpdateUnion 修改联盟定投广告组
// 【注】白名单账户支持定投联盟广告，修改广告组接口不变，有以下限制：1. 仅支持部分定向。
func UpdateUnion(clt *core.SDKClient, accessToken string, req *unit.UpdateUnionRequest) (uint64, error) {
	var resp unit.UpdateUnionResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.UnitID, nil
}
