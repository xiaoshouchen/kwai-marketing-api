package preput

import (
	"github.com/xiaoshouchen/kwai-marketing-api/core"
	"github.com/xiaoshouchen/kwai-marketing-api/model/tool/preput"
)

// PredicationTaskList 投前预估列表页接口
func PredicationTaskList(clt *core.SDKClient, accessToken string, req *preput.PredicationTaskListRequest) (*preput.PredicationTaskListResponse, error) {
	var resp preput.PredicationTaskListResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
