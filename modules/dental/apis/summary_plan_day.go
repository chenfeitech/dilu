package apis

import (
	"dilu/modules/dental/models"
	"dilu/modules/dental/service"
	"dilu/modules/dental/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SummaryPlanDayApi struct {
	base.BaseApi
}

var ApiSummaryPlanDay = SummaryPlanDayApi{}

// QueryPage 获取每日总结计划列表
// @Summary Page接口
// @Tags SummaryPlanDay
// @Accept application/json
// @Product application/json
// @Param data body dto.SummaryPlanDayGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SummaryPlanDay}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/summary-plan-day/page [post]
// @Security Bearer
func (e *SummaryPlanDayApi) QueryPage(c *gin.Context) {
	var req dto.SummaryPlanDayGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SummaryPlanDay, 10)
	var total int64

	var model models.SummaryPlanDay
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSummaryPlanDay.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取每日总结计划
// @Summary 获取每日总结计划
// @Tags SummaryPlanDay
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SummaryPlanDay} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/summary-plan-day/get [post]
// @Security Bearer
func (e *SummaryPlanDayApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SummaryPlanDay
	if err := service.SerSummaryPlanDay.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建每日总结计划
// @Summary 创建每日总结计划
// @Tags SummaryPlanDay
// @Accept application/json
// @Product application/json
// @Param data body dto.SummaryPlanDayDto true "body"
// @Success 200 {object} base.Resp{data=models.SummaryPlanDay} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/summary-plan-day/create [post]
// @Security Bearer
func (e *SummaryPlanDayApi) Create(c *gin.Context) {
	var req dto.SummaryPlanDayDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SummaryPlanDay
	copier.Copy(&data, req)
	if err := service.SerSummaryPlanDay.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新每日总结计划
// @Summary 更新每日总结计划
// @Tags SummaryPlanDay
// @Accept application/json
// @Product application/json
// @Param data body dto.SummaryPlanDayDto true "body"
// @Success 200 {object} base.Resp{data=models.SummaryPlanDay} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/summary-plan-day/update [post]
// @Security Bearer
func (e *SummaryPlanDayApi) Update(c *gin.Context) {
	var req dto.SummaryPlanDayDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SummaryPlanDay
	copier.Copy(&data, req)
	if err := service.SerSummaryPlanDay.Save(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除每日总结计划
// @Summary 删除每日总结计划
// @Tags SummaryPlanDay
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SummaryPlanDay} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/summary-plan-day/del [post]
// @Security Bearer
func (e *SummaryPlanDayApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSummaryPlanDay.DelIds(&models.SummaryPlanDay{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
