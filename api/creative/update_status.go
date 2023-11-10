package creative

import (
	"github.com/xiaoshouchen/kwai-marketing-api/core"
	"github.com/xiaoshouchen/kwai-marketing-api/model/creative"
)

// UpdateStatus 修改创意状态
func UpdateStatus(clt *core.SDKClient, accessToken string, req *creative.UpdateStatusRequest) ([]uint64, error) {
	var resp creative.UpdateStatusResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.CreativeIDs, nil
}
