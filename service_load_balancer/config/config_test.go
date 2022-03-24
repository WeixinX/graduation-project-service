package config

import (
	"fmt"
	"testing"
)

func TestConfigSetUp(t *testing.T) {
	configFile := "./config_test.json"
	configParams := ConfigSetUp(configFile)
	if configParams != nil {
		fmt.Printf("lb service name: %s\n", configParams.LBName)
		fmt.Printf("lb service listent port: %s\n\n", configParams.Port)
		fmt.Printf("upstream service name: %s\n\n", configParams.UpstreamServiceName)
		fmt.Printf("downstream service info:\n")
		fmt.Printf("instance replice number: %d\n", configParams.InstanceList.ReplicaNum)
		fmt.Printf("init total: %d\n", configParams.InstanceList.Total)
		fmt.Printf("service name: %s\n", configParams.InstanceList.ServiceName)
		fmt.Println("instance info:")
		for _, instance := range configParams.InstanceList.Instances {
			fmt.Printf("{id=%s, current_weight=%d, effective_weight=%d, call_url=%s}\n",
				instance.InstanceID, instance.CurrentWeight, instance.EffectiveWeight, instance.CallURL)
		}
	} else {
		t.Error("ConfigSetUp error!")
	}
}
