package request

import (
	"github.com/ldxtechteam/cpstbsdk/ability414/domain"
	"github.com/ldxtechteam/cpstbsdk/util"
)

type TaobaoTbkRelationRefundRequest struct {
	/*
	   参数option     */
	SearchOption *domain.TaobaoTbkRelationRefundTopApiRefundRptOption `json:"search_option" required:"true" `
}

func (s *TaobaoTbkRelationRefundRequest) SetSearchOption(v domain.TaobaoTbkRelationRefundTopApiRefundRptOption) *TaobaoTbkRelationRefundRequest {
	s.SearchOption = &v
	return s
}

func (req *TaobaoTbkRelationRefundRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SearchOption != nil {
		paramMap["search_option"] = util.ConvertStruct(*req.SearchOption)
	}
	return paramMap
}

func (req *TaobaoTbkRelationRefundRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
