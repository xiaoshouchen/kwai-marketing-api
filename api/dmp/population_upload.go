package dmp

import (
	"github.com/xiaoshouchen/kwai-marketing-api/core"
	"github.com/xiaoshouchen/kwai-marketing-api/model/dmp"
)

// PopulationUpload 人群包上传接口
func PopulationUpload(clt *core.SDKClient, accessToken string, req *dmp.PopulationUploadRequest) (*dmp.Population, error) {
	var resp dmp.Population
	err := clt.Upload(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
