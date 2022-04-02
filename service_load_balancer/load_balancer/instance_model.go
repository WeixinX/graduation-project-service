package load_balancer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Instance struct {
	InstanceID      string `json:"id"`
	CallURL         string `json:"call_url"`
	CurrentWeight   int    `json:"current_weight"`
	EffectiveWeight int    `json:"effective_weight"`
}

type InstanceList struct {
	ServiceName string         `json:"service_name"` // 服务名
	Instances   []Instance     `json:"instances"`    // 实例信息列表
	ReplicaNum  int            `json:"replica_num"`  // 当前副本数
	Total       int            `json:"total"`        // 总副本数
	InstanceMap map[string]int `json:"instance_map"` // InstanceID 到实例信息的映射
}

var INSTANCE_LIST *InstanceList

func NewInstanceList(configFile string) *InstanceList {
	// 读取配置文件
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Failed to read json configuration file!")
		return nil
	}

	list := &InstanceList{}
	err = json.Unmarshal(bytes, list)
	if err != nil {
		fmt.Println("json configuration file parsing failed!")
		return nil
	}

	if len(list.Instances) > list.Total {
		list.Total = len(list.Instances)
	}

	// 建立 InstanceID 与 InstanceList 下标的映射
	instanceMap := make(map[string]int, list.Total)
	for idx, instance := range list.Instances {
		instanceMap[instance.InstanceID] = idx
	}
	list.InstanceMap = instanceMap

	return list
}

func (i *InstanceList) AddReplica(addNum int) {
	if i.ReplicaNum+addNum > i.Total {
		i.ReplicaNum = i.Total
	} else {
		i.ReplicaNum += addNum
	}
}

func (i *InstanceList) AddEffectiveWeight(idx, weight int) {
	if idx > len(i.Instances) {
		return
	} else {
		i.Instances[idx].EffectiveWeight += weight
	}
}

func (i *InstanceList) RemoveInstance(instanceID string) {
	instances := make([]Instance, 0, INSTANCE_LIST.ReplicaNum)
	instanceMap := make(map[string]int, INSTANCE_LIST.ReplicaNum)

	for _, instance := range INSTANCE_LIST.Instances {
		if instance.InstanceID != instanceID {
			instances = append(instances, instance)
			instanceMap[instance.InstanceID] = len(instances) - 1
		}
	}

	INSTANCE_LIST.Total = len(instances)
	INSTANCE_LIST.Instances = instances
	INSTANCE_LIST.InstanceMap = instanceMap
}
