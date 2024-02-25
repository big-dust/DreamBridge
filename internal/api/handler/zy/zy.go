package zy

import (
	"context"
	"encoding/json"
	"github.com/big-dust/DreamBridge/internal/api/response"
	"github.com/big-dust/DreamBridge/internal/api/service/zy"
	"github.com/big-dust/DreamBridge/internal/api/types"
	"github.com/big-dust/DreamBridge/internal/pkg/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary	模拟报考
// @Description
// @Accept		json;multipart/form-data
// @Produce	json
// @Param		Authorization	header		string										true	"token"
// @Success	200				{object}	response.OkMsgDataResp[types.ZYMockResp]	"获取成功"
// @Failure	400				{object}	response.FailMsgResp						"获取模拟报考信息失败"
// @Router		/api/v1/zy/mock [get]
func Mock(c *gin.Context) {
	uid := c.GetInt("uid")
	owner := "/api/v1/zy/mock" + strconv.Itoa(uid)
	//取缓存
	cache, err := zy.GetResp(owner)
	if err == nil {
		resp := &types.ZYMockResp{}
		if err := json.Unmarshal(cache, resp); err == nil {
			response.OkMsgData(c, "获取成功", resp)
			return
		} else {
			common.LOG.Error("json.Unmarshal(cache, resp): Error:" + err.Error())
		}
	} else {
		common.LOG.Info("Get Cache Error: " + err.Error())
	}
	resp := &types.ZYMockResp{}
	errChan := make(chan error)
	p := context.Background()
	child, cancel := context.WithCancel(p)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		default:
		}
		//保 都是前闭后开
		bs, err := zy.GetSchools(ctx, uid, -10, -5)
		if err != nil {
			errChan <- err
			return
		}
		resp.BaoSchools = bs
		errChan <- nil
	}(child)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		default:
		}
		//稳
		ws, err := zy.GetSchools(child, uid, -5, 5)
		if err != nil {
			errChan <- err
			return
		}
		resp.WenSchools = ws
		errChan <- nil
	}(child)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		default:
		}
		//冲
		cs, err := zy.GetSchools(child, uid, 5, 11)
		if err != nil {
			errChan <- err
			return
		}
		resp.ChongSchools = cs
		errChan <- nil
	}(child)
	for i := 0; i < 3; i++ {
		if err := <-errChan; err != nil {
			cancel()
			response.FailMsg(c, "获取\"模拟报考信息失败\": "+err.Error())
			return
		}
	}
	cache, _ = json.Marshal(resp)
	if err = zy.SetResp(owner, cache); err != nil {
		common.LOG.Error("缓存失败: Owner: " + owner)
	}
	response.OkMsgData(c, "获取成功", resp)
	cancel()
}
