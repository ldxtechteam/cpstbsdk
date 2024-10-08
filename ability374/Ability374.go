package ability374

import (
	"errors"
	"github.com/ldxtechteam/cpstbsdk"
	"github.com/ldxtechteam/cpstbsdk/ability374/request"
	"github.com/ldxtechteam/cpstbsdk/ability374/response"
	"github.com/ldxtechteam/cpstbsdk/util"
	"log"
)

type Ability374 struct {
	Client *topsdk.TopClient
}

func NewAbility374(client *topsdk.TopClient) *Ability374 {
	return &Ability374{client}
}

/*
淘宝客-推广者-官方活动转链
*/
func (ability *Ability374) TaobaoTbkActivityInfoGet(req *request.TaobaoTbkActivityInfoGetRequest) (*response.TaobaoTbkActivityInfoGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability374 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.activity.info.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkActivityInfoGetResponse{}
	if err != nil {
		log.Println("taobaoTbkActivityInfoGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
