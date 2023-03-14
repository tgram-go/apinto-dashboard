package audit_controller

import (
	"context"
	"fmt"
	"github.com/eolinker/apinto-dashboard/access"
	controller2 "github.com/eolinker/apinto-dashboard/controller"
	"github.com/eolinker/apinto-dashboard/dto"
	"github.com/eolinker/apinto-dashboard/enum"
	"github.com/eolinker/apinto-dashboard/service/audit-service"
	"github.com/eolinker/eosc/common/bean"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type auditLogController struct {
	auditLogService audit_service.IAuditLogService
}

func RegisterAuditLogRouter(router gin.IRoutes) {
	a := &auditLogController{}
	bean.Autowired(&a.auditLogService)

	router.GET("/audit-logs", controller2.GenAccessHandler(access.AuditLogView), a.getLogs)
	router.GET("/audit-log", controller2.GenAccessHandler(access.AuditLogView), a.getDetail)
	router.GET("/audit-log/kinds", a.getTargets)
}

func (a *auditLogController) getLogs(ginCtx *gin.Context) {
	namespaceId := controller2.GetNamespaceId(ginCtx)

	operateType := ginCtx.Query("operate_type")
	kind := ginCtx.Query("kind")
	keyword := ginCtx.Query("keyword")
	startStr := ginCtx.Query("start")
	endStr := ginCtx.Query("end")
	pageNumStr := ginCtx.Query("page_num")
	pageSizeStr := ginCtx.Query("page_size")
	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageNum == 0 {
		pageNum = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}

	//判断操作目标合不合法
	if kind != "" && !enum.IsLogKindExist(kind) {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("kind %s is illegal. ", kind)))
		return
	}

	operate := 0
	if operateType != "" {
		//若operteType或非法,则会为0
		operate = int(enum.GetLogOperateIndex(operateType))
		if operate == 0 {
			ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("operate_type %s is illegal. ", operateType)))
			return
		}
	}

	var start, end int64
	var err error
	if startStr != "" {
		start, err = strconv.ParseInt(startStr, 10, 64)
		if err != nil {
			ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("start %s is illegal. ", startStr)))
			return
		}
	}
	if endStr != "" {
		end, err = strconv.ParseInt(endStr, 10, 64)
		if err != nil {
			ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("end %s is illegal. ", endStr)))
			return
		}
	}

	logList, total, err := a.auditLogService.GetLogsList(context.Background(), namespaceId, operate, kind, keyword, start, end, pageNum, pageSize)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("get audit-log fail. err: %s. ", err)))
		return
	}

	data := make(map[string]interface{})
	data["items"] = logList
	data["total"] = total
	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(data))
}

func (a *auditLogController) getDetail(ginCtx *gin.Context) {
	logIDStr := ginCtx.Query("log_id")
	logID, err := strconv.Atoi(logIDStr)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("log_id %s is illegal. err: %s", logIDStr, err)))
		return
	}

	details, err := a.auditLogService.GetLogDetail(context.Background(), logID)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("get log detail fail. err: %s. ", err)))
		return
	}
	data := make(map[string]interface{})
	data["args"] = details
	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(data))
}

func (a *auditLogController) getTargets(ginCtx *gin.Context) {
	data := make(map[string]interface{})
	data["items"] = enum.GetLogKinds()
	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(data))
}
