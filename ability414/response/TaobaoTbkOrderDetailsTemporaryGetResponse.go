package response

import (
	"github.com/ldxtechteam/cpstbsdk/ability414/domain"
)

type TaobaoTbkOrderDetailsTemporaryGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   PublisherOrderDto
	*/
	Data domain.TaobaoTbkOrderDetailsTemporaryGetOrderPage `json:"data,omitempty" `
}
