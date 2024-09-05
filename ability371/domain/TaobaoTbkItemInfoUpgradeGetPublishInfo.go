package domain


type TaobaoTbkItemInfoUpgradeGetPublishInfo struct {
    /*
        商品信息-收入比率(商品佣金比率+补贴比率)。15.5表示15.5%     */
    IncomeRate  *string `json:"income_rate,omitempty" `

    /*
        前N件佣金信息-前N件佣金生效或预热时透出以下字段     */
    TopnInfo  *TaobaoTbkItemInfoUpgradeGetTopNInfoDTO `json:"topn_info,omitempty" `

    /*
        单品淘礼金今日剩余可创建个数     */
    TljRemainNum  *int64 `json:"tlj_remain_num,omitempty" `

}

func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetIncomeRate(v string) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
    s.IncomeRate = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetTopnInfo(v TaobaoTbkItemInfoUpgradeGetTopNInfoDTO) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
    s.TopnInfo = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetTljRemainNum(v int64) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
    s.TljRemainNum = &v
    return s
}
