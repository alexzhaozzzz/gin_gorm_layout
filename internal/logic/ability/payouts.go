package ability

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/ability"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/common"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewPayouts() *Payouts {
	return &Payouts{}
}

type Payouts struct {
}

// GetListByPage ...
// @Summary 获取多个标签
// @Tags todos
// @Accept json
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=dto.PayoutsReq} "成功"
// @Router /api/v1/tags [get]
func (s Payouts) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("PayoutsReq GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.PayoutsReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("PayoutsReq GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("PayoutsReq GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewPayoutsData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("PayoutsReq GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount(cIds, req)
	if err != nil {
		logrus.Errorf("PayoutsReq GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	pageResp := &dto.Pagination{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		TotalNum:  int(total),
	}

	resp := map[string]interface{}{"list": list, "page": pageResp}
	c.RenderSuccess(resp)
	return
}

func (s Payouts) GetAuditListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("PayoutsReq GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.PayoutsAuditReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("PayoutsReq GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("PayoutsReq GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewPayoutsData()
	list, err := d.ListByStatePage(cIds, 0, req)
	if err != nil {
		logrus.Errorf("PayoutsReq GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	respList := make([]dto.PayoutsAuditResp, 0)
	if len(list) > 0 {

		dpp := ability.NewPlayerProfileData()

		for _, v := range list {
			profileInfo, err := dpp.InfoById(v.Playerid)
			if err != nil {
				continue
			}

			respList = append(respList, dto.PayoutsAuditResp{
				Id:            v.Id,
				Orderid:       v.Orderid,
				Playerid:      v.Playerid,
				Amount:        v.Amount,
				State:         v.State,
				Des:           v.Des,
				Createtime:    v.Createtime,
				Paytime:       v.Paytime,
				Checktime:     v.Checktime,
				Coin:          v.Coin,
				Channel:       v.Channel,
				ExOrderid:     v.ExOrderid,
				Bankinfo:      v.Bankinfo,
				ExResp:        v.ExResp,
				BackCoin:      v.BackCoin,
				Operator:      v.Operator,
				Fee:           v.Fee,
				From:          v.From,
				Partnerid:     v.Partnerid,
				Platform:      v.Platform,
				After:         v.After,
				Kind:          v.Kind,
				ShowErr:       v.ShowErr,
				CardId:        v.CardId,
				PayCreatetime: v.PayCreatetime,
				Retry:         v.Retry,
				RetryState:    v.RetryState,
				QueryState:    v.QueryState,
				QueryResp:     v.QueryResp,
				LastOrderId:   v.LastOrderId,
				Red:           v.Red,
				WdFlag:        v.WdFlag,
				UserFlag:      v.UserFlag,
				UserKind:      profileInfo.NKind,
			})
		}
	}

	total, err := d.ListStateCount(cIds, 0, req)
	if err != nil {
		logrus.Errorf("PayoutsReq GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	pageResp := &dto.Pagination{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		TotalNum:  int(total),
	}

	resp := map[string]interface{}{"list": respList, "page": pageResp}
	c.RenderSuccess(resp)
	return
}

func (s Payouts) GetUserBaseInfo(c *ginx.Context) {
	req := &dto.UserBaseReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("PayoutsReq GetUserFoundationInfo ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dp := ability.NewPlayerData()
	dpInfo, err := dp.InfoById(req.PlayerId)
	if err != nil {
		logrus.Errorf("PayoutsReq dp.InfoById Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dpd := ability.NewPlayerDeviceData()
	dpdInfo, err := dpd.InfoById(uint64(req.PlayerId))
	if err != nil {
		logrus.Errorf("PayoutsReq dpd.InfoById Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dpp := ability.NewPlayerProfileData()
	dppInfo, err := dpp.InfoById(uint64(req.PlayerId))
	if err != nil {
		logrus.Errorf("PayoutsReq dpp.InfoById Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dpb := ability.NewPlayerBankData()
	dpbInfo, err := dpb.InfoById(uint64(req.PlayerId))
	if err != nil {
		logrus.Errorf("PayoutsReq dpb.InfoById Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	respInfo := dto.UserBaseResp{
		PlayerId:   req.PlayerId,
		Tag:        dppInfo.NTag,
		UType:      dppInfo.NType,
		NickName:   dpInfo.Nickname,
		Card:       dpbInfo.NBankcardid,
		CardName:   dpbInfo.NName,
		Phone:      dpbInfo.NPhone,
		DeviceCode: dpdInfo.NDevCode,
	}

	resp := map[string]interface{}{"info": respInfo}
	c.RenderSuccess(resp)
	return
}

func (s Payouts) GetUserListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("PayoutsReq GetUserListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.UserListReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("PayoutsReq GetUserListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("PayoutsReq GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewPayoutsData()
	list, err := d.ListByPlayerPage(cIds, req.PlayerId, req)
	if err != nil {
		logrus.Errorf("PayoutsReq GetUserListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListPlayerCount(cIds, req.PlayerId)
	if err != nil {
		logrus.Errorf("PayoutsReq GetUserListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	pageResp := &dto.Pagination{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		TotalNum:  int(total),
	}

	resp := map[string]interface{}{"list": list, "page": pageResp}
	c.RenderSuccess(resp)
	return
}