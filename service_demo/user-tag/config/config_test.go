package config

import (
	"fmt"
	"testing"
)

func TestConfigSetUp(t *testing.T) {
	configFile := "./config_test.json"
	CONFIG_PARAMS = ConfigSetUp(configFile)
	if CONFIG_PARAMS != nil {
		fmt.Printf("service_name: %s\n", CONFIG_PARAMS.ServiceName)
		fmt.Printf("instance_id: %s\n", CONFIG_PARAMS.InstanceID)
		fmt.Printf("port: %s\n\n", CONFIG_PARAMS.Port)
		fmt.Println("downstream_call_list info:")
		for _, downstream := range CONFIG_PARAMS.DownstreamCallList {
			fmt.Printf("{%s=%s}\n", downstream.ServiceName, downstream.LBCallURL)
		}
		fmt.Println("downstream_call_pair:")
		for _, downstream := range CONFIG_PARAMS.DownstreamCallList {
			fmt.Printf("{%s=%s}\n", downstream.ServiceName, CONFIG_PARAMS.DownstreamCallPair[downstream.ServiceName])
		}
	}
}
