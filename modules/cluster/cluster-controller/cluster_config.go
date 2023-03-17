package cluster_controller

import (
	"fmt"
	"github.com/eolinker/apinto-dashboard/access"
	"github.com/eolinker/apinto-dashboard/controller"
	"github.com/eolinker/apinto-dashboard/modules/base/namespace-controller"
	"github.com/eolinker/apinto-dashboard/modules/cluster"
	"github.com/eolinker/eosc/common/bean"
	"github.com/gin-gonic/gin"
	"net/http"
)

type clusterConfigController struct {
	configService cluster.IClusterConfigService
}

func RegisterClusterConfigRouter(router gin.IRoutes) {
	c := &clusterConfigController{}
	bean.Autowired(&c.configService)

	router.GET("/cluster/:cluster_name/configuration/:type", controller.GenAccessHandler(access.ClusterView, access.ClusterEdit), c.get)
	router.PUT("/cluster/:cluster_name/configuration/:type", controller.GenAccessHandler(access.ClusterEdit), c.edit)
	router.PUT("/cluster/:cluster_name/configuration/:type/enable", controller.GenAccessHandler(access.ClusterEdit), c.enable)
	router.PUT("/cluster/:cluster_name/configuration/:type/disable", controller.GenAccessHandler(access.ClusterEdit), c.disable)
}

func (c *clusterConfigController) get(ginCtx *gin.Context) {
	namespaceId := namespace_controller.GetNamespaceId(ginCtx)
	clusterName := ginCtx.Param("cluster_name")
	configType := ginCtx.Param("type")

	if !c.configService.IsConfigTypeExist(configType) {
		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("get %s fail. err: %s doesn't exist. ", configType, configType)))
		return
	}

	info, err := c.configService.Get(ginCtx, namespaceId, clusterName, configType)
	if err != nil {
		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("get %s fail. err: %s ", configType, err)))
		return
	}

	data := make(map[string]interface{})
	if info != nil {
		data[configType] = info
	}

	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(data))
}

func (c *clusterConfigController) edit(ginCtx *gin.Context) {
	namespaceId := namespace_controller.GetNamespaceId(ginCtx)
	clusterName := ginCtx.Param("cluster_name")
	configType := ginCtx.Param("type")
	operator := controller.GetUserId(ginCtx)

	if !c.configService.IsConfigTypeExist(configType) {
		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("edit %s fail. err: %s doesn't exist. ", configType, configType)))
		return
	}

	body, err := ginCtx.GetRawData()
	if err != nil {
		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("edit %s fail. err: %s ", configType, err)))
		return
	}

	if err = c.configService.CheckInput(configType, body); err != nil {
		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("edit %s fail. err: %s ", configType, err)))
		return
	}

	err = c.configService.Edit(ginCtx, namespaceId, operator, clusterName, configType, body)
	if err != nil {
		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("edit %s fail. err: %s ", configType, err)))
		return
	}

	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
}

func (c *clusterConfigController) enable(ginCtx *gin.Context) {
	namespaceId := namespace_controller.GetNamespaceId(ginCtx)
	clusterName := ginCtx.Param("cluster_name")
	configType := ginCtx.Param("type")
	operator := controller.GetUserId(ginCtx)

	if !c.configService.IsConfigTypeExist(configType) {
		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("enable %s fail. err: %s doesn't exist. ", configType, configType)))
		return
	}

	if err := c.configService.Enable(ginCtx, namespaceId, operator, clusterName, configType); err != nil {
		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("enable %s fail. err: %s  ", configType, err)))
		return
	}
	ginCtx.JSON(http.StatusOK, controller.Result{
		Msg: fmt.Sprintf("已启用%s", configType),
	})
}

func (c *clusterConfigController) disable(ginCtx *gin.Context) {
	namespaceId := namespace_controller.GetNamespaceId(ginCtx)
	clusterName := ginCtx.Param("cluster_name")
	configType := ginCtx.Param("type")
	operator := controller.GetUserId(ginCtx)

	if !c.configService.IsConfigTypeExist(configType) {
		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("disable %s fail. err: %s doesn't exist. ", configType, configType)))
		return
	}

	if err := c.configService.Disable(ginCtx, namespaceId, operator, clusterName, configType); err != nil {
		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("disable %s fail. err: %s  ", configType, err)))
		return
	}
	ginCtx.JSON(http.StatusOK, controller.Result{
		Msg: fmt.Sprintf("已停用%s", configType),
	})
}