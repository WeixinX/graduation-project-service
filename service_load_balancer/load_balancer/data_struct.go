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
	ServiceName string     `json:"service_name"`
	Instances   []Instance `json:"instances"`
	ReplicaNum  int        `json:"replica_num"`
	Total       int        `json:"total"`
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

	return list
}
