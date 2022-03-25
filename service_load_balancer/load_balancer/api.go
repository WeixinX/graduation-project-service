package load_balancer

import (
	"fmt"
	"net/http"

	"github.com/WeixinX/graduation-project-service/service_load_balancer/request"
	"github.com/gin-gonic/gin"
)

// CallDown POST调用下游服务
func CallDown(ctx *gin.Context) {
	// 通过负载均衡策略选择下游实例进行调用
	willCallInstance := WRR(INSTANCE_LIST)

	// postman测试时，解开注释调试信息
	//testMessage := fmt.Sprintf("The instance being called this time: %s\n", willCallInstance.InstanceID)
	//testMessage += "Weights for the next selection reference:\n{"
	//for i := 0; i < INSTANCE_LIST.ReplicaNum; i++ {
	//	testMessage += fmt.Sprintf("%s=%d",
	//		INSTANCE_LIST.Instances[i].InstanceID, INSTANCE_LIST.Instances[i].CurrentWeight)
	//
	//	if i != INSTANCE_LIST.ReplicaNum-1 {
	//		testMessage += fmt.Sprintf(", ")
	//	}
	//}
	//testMessage += "}"
	//fmt.Println(testMessage)
	//ctx.JSON(http.StatusOK, gin.H{"test message": testMessage})

	if willCallInstance == nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": "The load balancing algorithm failed to select an instance!"})

	} else {
		requestParams := request.ReqParams{
			UrlStr: willCallInstance.CallURL,
			Method: ctx.Request.Method,
			Header: ctx.Request.Header,
			Body:   ctx.Request.Body,
		}

		fmt.Printf("LB selected %s\n", willCallInstance.InstanceID)
		resp, err := request.HttpDo(ctx, &requestParams)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": "The load balancer request failed: " + err.Error()})
		}
		// 把相应信息透传回上游
		ctx.JSON(http.StatusOK, resp)
	}
}

// Scale 扩容
func Scale(ctx *gin.Context) {

}

// Transfer 转移/均衡
func Transfer(ctx *gin.Context) {

}

// Reschedule 重调度
func Reschedule(ctx *gin.Context) {

}

// GetReplicaNum 获取下游实例当前副本数
func GetReplicaNum(ctx *gin.Context) {
	if INSTANCE_LIST == nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": "The list of downstream service instances is not initialized!"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   map[string]interface{}{"replica_num": INSTANCE_LIST.ReplicaNum},
		})
	}
}
