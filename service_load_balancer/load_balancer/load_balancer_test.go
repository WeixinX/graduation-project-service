package load_balancer

import (
	"fmt"
	"testing"
)

func TestNewInstanceList(t *testing.T) {
	configFile := "./instances_test.json"
	list := NewInstanceList(configFile)
	if list != nil {
		for _, instance := range list.Instances {
			fmt.Printf("%+v\n", instance)
		}
		fmt.Println("Replica Num: ", list.ReplicaNum)
		fmt.Println("Total: ", list.Total)
		fmt.Println("Instance Map: ", list.InstanceMap)
	} else {
		t.Error("NewInstanceList error!")
	}
}

func TestWRR(t *testing.T) {
	configFile := "./instances_test.json"
	list := NewInstanceList(configFile)
	if list != nil {
		// 循环7次看效果
		testLen := 3
		for i := 1; i <= testLen; i++ {
			fmt.Println("No.", i, "-------------")

			fmt.Println("before WRR: ")
			//fmt.Print("{")
			//for j := 0; j < list.ReplicaNum; j++ {
			//	fmt.Printf("%s=%d", list.Instances[j].InstanceID, list.Instances[j].CurrentWeight)
			//	if j != list.ReplicaNum-1 {
			//		fmt.Print(", ")
			//	}
			//}
			//fmt.Println("}")

			choose := WRR(list)
			fmt.Println("choosed instance: ", choose.InstanceID)

			fmt.Println("after WRR: ")
			fmt.Print("{")
			for j := 0; j < list.ReplicaNum; j++ {
				fmt.Printf("%s=%d", list.Instances[j].InstanceID, list.Instances[j].CurrentWeight)
				if j != list.ReplicaNum-1 {
					fmt.Print(", ")
				}
			}
			fmt.Println("}")
			fmt.Println()
		}
	} else {
		t.Error("NewInstanceList error!")
	}
}
