package response

import (
	"github.com/ldxtechteam/cpstbsdk/defaultability/domain"
)

type TaobaoTbkDgTpwdRiskReportResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   无
	*/
	Result domain.TaobaoTbkDgTpwdRiskReportResult `json:"result,omitempty" `
}
