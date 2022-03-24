package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"WeixinX/graduation-project/service_load_balancer/load_balancer"
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

	return configParams
}
