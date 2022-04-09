package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConfigParams struct {
	ServiceName        string            `json:"service_name"` // 该服务名称
	InstanceID         string            `json:"instance_id"`  // 该服务实例标识符
	Port               string            `json:"port"`         // 该服务实例运行端口
	JaegerAgent        string            `json:"jaeger_agent"` // 设置 jaeger agent 主机:端口
	DownstreamCallList []Downstream      `json:"downstream_call_list"`
	DownstreamCallPair map[string]string // 方便用下游服务名称映射到调用 URL
}

type Downstream struct {
	ServiceName string `json:"service_name"` // 下游服务名称
	LBCallURL   string `json:"lb_call_url"`  // 该服务与下游实例中间的 LB 调用 URL
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

	configParams.DownstreamCallPair = make(map[string]string)
	if configParams != nil {
		for _, downstream := range configParams.DownstreamCallList {
			configParams.DownstreamCallPair[downstream.ServiceName] = downstream.LBCallURL
		}
	}

	return configParams
}
