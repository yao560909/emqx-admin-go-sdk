package nodes

import "github.com/yao560909/emqx-admin-go-sdk/core"

type ListNodesReq struct {
	apiReq *core.APIReq
}

type ListNodesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Node
}
