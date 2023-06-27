package cluster_service

import (
	"github.com/eolinker/apinto-dashboard/cache"
	"github.com/eolinker/eosc/common/bean"
	"github.com/go-redis/redis/v8"
)

func init() {

	//clConfigDriverManager := newCLConfigDriverManager()
	//redisDriver := driver2.CreateRedis("redis")
	//clConfigDriverManager.RegisterDriver(cluster.CLConfigRedis, redisDriver)
	//
	//bean.Injection(&clConfigDriverManager)

	iClusterService := newClusterService()
	clusterCertificate := newClusterCertificateService()
	clusterNode := newClusterNodeService()
	//clusterConfig := newClusterConfigService()

	apintoClient := newApintoClientService()

	bean.Injection(&apintoClient)

	bean.Injection(&iClusterService)
	bean.Injection(&clusterCertificate)
	bean.Injection(&clusterNode)
	//bean.Injection(&clusterConfig)

	cache.RegisterCacheInitHandler(func(client *redis.ClusterClient) {
		nodeCache := newINodeCache(client)
		bean.Injection(&nodeCache)
	})
}
