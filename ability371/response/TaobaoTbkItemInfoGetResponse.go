package response

import (
	"github.com/ldxtechteam/cpstbsdk/ability371/domain"
)

type TaobaoTbkItemInfoGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   淘宝客商品
	*/
	Results []domain.TaobaoTbkItemInfoGetNTbkItem `json:"results,omitempty" `
}
