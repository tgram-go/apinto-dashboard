package controller

import (
	"github.com/eolinker/apinto-dashboard/access"
	"github.com/eolinker/apinto-dashboard/common"
	"github.com/eolinker/apinto-dashboard/dto"
	"github.com/eolinker/apinto-dashboard/service"
	"github.com/eolinker/eosc/common/bean"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type clusterCertificateController struct {
	clusterCertificateService service.IClusterCertificateService
}

func RegisterClusterCertificateRouter(router gin.IRoutes) {
	c := &clusterCertificateController{}
	bean.Autowired(&c.clusterCertificateService)

	router.POST("/cluster/:cluster_name/certificate", genAccessHandler(access.ClusterEdit), c.post)
	router.PUT("/cluster/:cluster_name/certificate/:certificate_id", genAccessHandler(access.ClusterEdit), c.put)
	router.DELETE("/cluster/:cluster_name/certificate/:certificate_id", genAccessHandler(access.ClusterEdit), c.del)
	router.GET("/cluster/:cluster_name/certificates", genAccessHandler(access.ClusterView, access.ClusterEdit), c.gets)
}

// gets 获取证书列表
func (c *clusterCertificateController) gets(ginCtx *gin.Context) {
	namespaceId := getNamespaceId(ginCtx)
	clusterName := ginCtx.Param("cluster_name")

	list, err := c.clusterCertificateService.QueryList(ginCtx, namespaceId, clusterName)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
		return
	}
	dtoList := make([]*dto.ClusterCertificateOut, 0, len(list))
	for _, val := range list {
		cert, err := common.ParseCert(val.Key, val.Pem)
		if err != nil {
			ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
			return
		}
		dtoList = append(dtoList, &dto.ClusterCertificateOut{
			Id:           val.Id,
			ClusterId:    val.ClusterId,
			Name:         cert.Leaf.Subject.CommonName,
			ValidTime:    common.TimeToStr(cert.Leaf.NotAfter),
			OperatorName: val.OperatorName,
			CreateTime:   common.TimeToStr(val.CreateTime),
			UpdateTime:   common.TimeToStr(val.UpdateTime),
		})
	}
	m := common.Map[string, interface{}]{}
	m["certificates"] = dtoList
	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(m))
}

// post 新增
func (c *clusterCertificateController) post(ginCtx *gin.Context) {
	namespaceId := getNamespaceId(ginCtx)
	clusterName := ginCtx.Param("cluster_name")
	operator := getUserId(ginCtx)
	input := &dto.ClusterCertificateInput{}
	err := ginCtx.BindJSON(input)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
		return
	}
	if len(input.Key) == 0 || len(input.Pem) == 0 {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult("key or pem is null"))
		return
	}

	pem, _ := common.Base64Decode(input.Pem)
	key, _ := common.Base64Decode(input.Key)

	if err = c.clusterCertificateService.Insert(ginCtx, operator, namespaceId, clusterName, string(key), string(pem)); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
		return
	}
	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))

}

// put 修改
func (c *clusterCertificateController) put(ginCtx *gin.Context) {
	namespaceId := getNamespaceId(ginCtx)
	clusterName := ginCtx.Param("cluster_name")
	certificateIdStr := ginCtx.Param("certificate_id")
	certificateId, _ := strconv.Atoi(certificateIdStr)
	operator := getUserId(ginCtx)
	if certificateId <= 0 {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult("certificate_id is 0"))
		return
	}
	input := &dto.ClusterCertificateInput{}
	err := ginCtx.BindJSON(input)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
		return
	}

	if len(input.Key) == 0 || len(input.Pem) == 0 {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult("key or pem is null"))
		return
	}

	pem, _ := common.Base64Decode(input.Pem)
	key, _ := common.Base64Decode(input.Key)

	if err = c.clusterCertificateService.Update(ginCtx, operator, namespaceId, certificateId, clusterName, string(key), string(pem)); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
		return
	}
	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))

}

// del 删除
func (c *clusterCertificateController) del(ginCtx *gin.Context) {
	namespaceId := getNamespaceId(ginCtx)
	clusterName := ginCtx.Param("cluster_name")
	certificateIdStr := ginCtx.Param("certificate_id")
	certificateId, _ := strconv.Atoi(certificateIdStr)

	err := c.clusterCertificateService.DeleteById(ginCtx, namespaceId, clusterName, certificateId)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
		return
	}
	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))

}
