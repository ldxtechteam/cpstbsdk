package response

import (
	"github.com/ldxtechteam/cpstbsdk/ability425/domain"
)

type TaobaoTbkScPublisherInfoGetResponse struct {

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
	Data domain.TaobaoTbkScPublisherInfoGetData `json:"data,omitempty" `
}
