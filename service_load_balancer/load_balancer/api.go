package load_balancer

import (
	"fmt"
	"net/http"
	"strconv"

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
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "The load balancing algorithm failed to select an instance!",
		})

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
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "The load balancer request failed: " + err.Error(),
			})
		}
		// 把相应信息透传回上游
		ctx.JSON(http.StatusOK, resp)
	}
}

// Scaling 扩容
func Scaling(ctx *gin.Context) {
	addNumStr := ctx.DefaultQuery("add_num", "0")
	addNum, _ := strconv.Atoi(addNumStr)
	INSTANCE_LIST.AddReplica(addNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": map[string]interface{}{
			"current_replica_num": INSTANCE_LIST.ReplicaNum,
			"total_replica_num":   INSTANCE_LIST.Total,
		},
	})
}

// Balancing 转移/均衡
func Balancing(ctx *gin.Context) {
	instanceID := ctx.DefaultQuery("instance_id", "nil")
	weightStr := ctx.DefaultQuery("weight", "0")
	weight, _ := strconv.Atoi(weightStr)
	if instanceID == "nil" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "The instance ID parameter is invalid!",
		})

	} else {
		if idx, ok := INSTANCE_LIST.InstanceMap[instanceID]; !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "The instance ID parameter is invalid!",
			})
		} else {
			INSTANCE_LIST.AddEffectiveWeight(idx, weight)
			ctx.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	}
}

// Reschedule 重调度
func Reschedule(ctx *gin.Context) {
	// 把指定实例从列表中删除然后加入一个实例来模拟
	instanceID := ctx.DefaultQuery("instance_id", "nil")
	if instanceID == "nil" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "The instance ID parameter is invalid!",
		})

	} else {
		if _, ok := INSTANCE_LIST.InstanceMap[instanceID]; !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "The instance ID parameter is invalid!",
			})

		} else {
			INSTANCE_LIST.RemoveInstance(instanceID)
			ctx.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	}
}

// GetReplicaNum 获取下游实例当前副本数
func GetReplicaNum(ctx *gin.Context) {
	if INSTANCE_LIST == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "The list of downstream service instances is not initialized!",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": map[string]interface{}{
				"current_replica_num": INSTANCE_LIST.ReplicaNum,
				"total_replica_num":   INSTANCE_LIST.Total,
			},
		})
	}
}
