package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/WeixinX/graduation-project-service/service_load_balancer/load_balancer"
)

type ConfigParams struct {
	LBName              string                      `json:"load_balancer_name"`       // LB service 名称
	Port                string                      `json:"port"`                     // LB service 运行端口
	UpstreamServiceName string                      `json:"upstream_service_name"`    // 上游服务名称
	InstanceList        *load_balancer.InstanceList `json:"downstream_instance_list"` // 下游服务实例信息
}

var CONFIG_PARAMS *ConfigParams

func ConfigSetUp(configFile string) *ConfigParams {
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Failed to read json configuration file! err: ", err)
		return nil
	}

	configParams := &ConfigParams{}
	err = json.Unmarshal(bytes, configParams)
	if err != nil {
		fmt.Println("json configuration file parsing failed! err: ", err)
		return nil
	}

	if len(configParams.InstanceList.Instances) > configParams.InstanceList.Total {
		configParams.InstanceList.Total = len(configParams.InstanceList.Instances)
	}

	// 建立 InstanceID 与 InstanceList 下标的映射
	instanceMap := make(map[string]int, configParams.InstanceList.Total)
	for idx, instance := range configParams.InstanceList.Instances {
		instanceMap[instance.InstanceID] = idx
	}
	configParams.InstanceList.InstanceMap = instanceMap

	return configParams
}
