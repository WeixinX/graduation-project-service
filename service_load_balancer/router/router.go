package router

import (
	"github.com/WeixinX/graduation-project-service/service_load_balancer/load_balancer"
	"github.com/gin-gonic/gin"
)

func RouterSetUp(engine *gin.Engine) {
	engine.Use()

	apiGroup := engine.Group("/lb_api")
	{
		// 调用下游服务
		apiGroup.GET("/call", load_balancer.CallDown)
		apiGroup.POST("/call", load_balancer.CallDown)

		// 扩容
		apiGroup.POST("/scaling", load_balancer.Scaling)

		// 转移/均衡
		apiGroup.POST("/balancing", load_balancer.Balancing)

		// 重调度
		apiGroup.POST("/reschedule", load_balancer.Reschedule)

		// 获取下游实例副本数
		apiGroup.GET("/get_replica", load_balancer.GetReplicaNum)
	}

}
