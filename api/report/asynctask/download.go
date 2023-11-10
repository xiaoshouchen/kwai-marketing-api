package asynctask

import (
	"github.com/xiaoshouchen/kwai-marketing-api/core"
	"github.com/xiaoshouchen/kwai-marketing-api/model/report/asynctask"
)

func Download(clt *core.SDKClient, accessToken string, req *asynctask.DownloadRequest) ([]byte, error) {
	return clt.GetBytes(accessToken, req)
}
