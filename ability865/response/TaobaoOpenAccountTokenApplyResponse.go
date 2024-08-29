package response

import (
	"github.com/ldxtechteam/cpstbsdk/ability865/domain"
)

type TaobaoOpenAccountTokenApplyResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回的token结果
	*/
	Data domain.TaobaoOpenAccountTokenApplyOpenAccountTokenApplyResult `json:"data,omitempty" `
}
