package user

import (
	"github.com/big-dust/DreamBridge/internal/api/response"
	"github.com/big-dust/DreamBridge/internal/api/service/user"
	"github.com/big-dust/DreamBridge/internal/api/types"
	"github.com/gin-gonic/gin"
	"strings"
)

// @Summary		设置用户报考信息
// @Description	设置用户信息接口
// @Accept			json;multipart/form-data
// @Produce		json
// @Param			Authorization		header		string					true	"token"
// @Param			province			formData		string					true	"省份"
// @Param			exam_type			formData		string					true	"考试类型"
// @Param			school_type			formData		string					true	"学校类型"
// @Param			subject				formData		object					true	"科目"
// @Param			subject.formData.physics		formData		bool					true	"是否选择物理学"
// @Param			subject.formData.history		formData		bool					true	"是否选择历史学"
// @Param			subject.formData.chemistry	formData		bool					true	"是否选择化学"
// @Param			subject.formData.biology		formData		bool					true	"是否选择生物学"
// @Param			subject.formData.geography	formData		bool					true	"是否选择地理学"
// @Param			subject.formData.politics	formData		bool					true	"是否选择政治学"
// @Param			score				formData		int						true	"分数"
// @Param			province_rank		formData		int						true	"省份排名"
// @Param			holland				formData		string					true	"霍兰德"	Enums(conventional,investigative,realistic,enterprising,artistic,social)
// @Param			interests			formData		[]string				true	"兴趣列表"
// @Success		200					{object}	response.OkMsgResp		"更新信息成功"
// @Failure		400					{object}	response.FailMsgResp	"更新信息失败"
// @Router			/api/v1/user/info [post]
func SetInfo(c *gin.Context) {
	req := &types.UserSetInfoReq{}
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, "更新信息失败: "+err.Error())
		return
	}
	uid := c.GetInt("uid")
	if err := user.SetUserInfo(uid, req); err != nil {
		response.FailMsg(c, "更新信息失败: "+err.Error())
		return
	}
	response.OkMsg(c, "更新信息成功")
}

// @Summary		获取用户报考信息
// @Description	获取用户信息接口
// @Accept			json;multipart/form-data
// @Produce		json
// @Param			Authorization		header		string					true	"token"
// @Success		200					{object}	response.OkMsgDataResp[types.UserGetInfoResp]		"获取信息成功"
// @Failure		400					{object}	response.FailMsgResp	"获取信息失败"
// @Router			/api/v1/user/info [get]
func GetInfo(c *gin.Context) {
	uid := c.GetInt("uid")
	info, err := user.GetUserInfo(uid)
	if err != nil {
		response.FailMsg(c, "获取信息失败: "+err.Error())
		return
	}
	resp := &types.UserGetInfoResp{
		Province:   info.Province,
		ExamType:   info.ExamType,
		SchoolType: info.SchoolType,
		Subject: types.Subject{
			Physics:   info.Physics,
			History:   info.History,
			Chemistry: info.Chemistry,
			Biology:   info.Biology,
			Geography: info.Geography,
			Politics:  info.Politics,
		},
		Score:        info.Score,
		ProvinceRank: info.ProvinceRank,
		Holland:      info.Holland,
		Interests:    strings.Split(info.Interests, " "),
	}
	response.OkMsgData(c, "更新信息成功", resp)
}
