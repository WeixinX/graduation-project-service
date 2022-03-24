package load_balancer

import (
	"math/rand"
	"time"
)

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

/*
	加权轮询算法 Weighted Round Robin, WRR
	过程:
		1. 遍历实例列表 instance.CurrentWeight += instance.EffectiveWeight, total += instance.EffectiveWeight
		2. 选出实例中 CurrentWeight 最大的作为此次负载实例，如果一样大则随机选择
		3. 对于选处的实例，执行 instance.CurrentWeight -= Total
*/

func WRR(list *InstanceList) (retInstance *Instance) {
	if list == nil || list.Instances == nil || len(list.Instances) == 0 {
		return nil
	}

	maxWeight := INT_MIN
	total := 0
	randomList := make([]*Instance, 0, list.ReplicaNum)

	// 单元测试 WRR 时，解开下面三块注释
	//fmt.Print("{")

	for i := 0; i < list.ReplicaNum; i++ {
		list.Instances[i].CurrentWeight += list.Instances[i].EffectiveWeight
		total += list.Instances[i].EffectiveWeight

		if list.Instances[i].CurrentWeight > maxWeight {
			maxWeight = list.Instances[i].CurrentWeight

			// 清空数组并添加当前实例 index
			randomList = make([]*Instance, 0, list.ReplicaNum)
			randomList = append(randomList, &list.Instances[i])

		} else if list.Instances[i].CurrentWeight == maxWeight {
			// 将最大相同权重的实例 index 放在一个数组，随机选择
			randomList = append(randomList, &list.Instances[i])
		}

		//fmt.Printf("%s=%d", list.Instances[i].InstanceID, list.Instances[i].CurrentWeight)
		//if i != list.ReplicaNum-1 {
		//	fmt.Print(", ")
		//}
	}
	//fmt.Println("}")

	list.Total = total

	// 相同权重，使用随机算法
	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	ind := random.Intn(len(randomList))

	randomList[ind].CurrentWeight -= total

	return randomList[ind]
}
