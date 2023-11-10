package advertiser

import (
	"github.com/xiaoshouchen/kwai-marketing-api/core"
	"github.com/xiaoshouchen/kwai-marketing-api/model/advertiser"
)

// UpdateBudget 修改账户预算
func UpdateBudget(clt *core.SDKClient, accessToken string, req *advertiser.UpdateBudgetRequest) error {
	return clt.Post(accessToken, req, nil)
}
