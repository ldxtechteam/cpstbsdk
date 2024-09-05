package response

import (
	"github.com/ldxtechteam/cpstbsdk/ability371/domain"
)

type TaobaoTbkCouponGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   data
	*/
	Data domain.TaobaoTbkCouponGetMapData `json:"data,omitempty" `
}
