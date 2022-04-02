package router

import (
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WeixinX/graduation-project-service/service_load_balancer/config"
	"github.com/WeixinX/graduation-project-service/service_load_balancer/load_balancer"

	"github.com/gin-gonic/gin"
)

func TestScaling(t *testing.T) {
	configFile := "../config/config_test.json"
	config.CONFIG_PARAMS = config.ConfigSetUp(configFile)
	load_balancer.INSTANCE_LIST = config.CONFIG_PARAMS.InstanceList

	fmt.Println("before action:")
	PrintInstanceListPretty(load_balancer.INSTANCE_LIST)

	r := gin.Default()
	RouterSetUp(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/lb_api/scaling?add_num=1", nil)
	r.ServeHTTP(w, req)

	fmt.Println("after action:")
	PrintInstanceListPretty(load_balancer.INSTANCE_LIST)

}

func TestBalancing(t *testing.T) {
	configFile := "../config/config_test.json"
	config.CONFIG_PARAMS = config.ConfigSetUp(configFile)
	load_balancer.INSTANCE_LIST = config.CONFIG_PARAMS.InstanceList

	fmt.Println("before action:")
	PrintInstanceListPretty(load_balancer.INSTANCE_LIST)

	r := gin.Default()
	RouterSetUp(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/lb_api/balancing?instance_id=service_a_2&weight=10", nil)
	r.ServeHTTP(w, req)

	fmt.Println("after action:")
	PrintInstanceListPretty(load_balancer.INSTANCE_LIST)
}

func TestReschedule(t *testing.T) {
	configFile := "../config/config_test.json"
	config.CONFIG_PARAMS = config.ConfigSetUp(configFile)
	load_balancer.INSTANCE_LIST = config.CONFIG_PARAMS.InstanceList

	fmt.Println("before action:")
	PrintInstanceListPretty(load_balancer.INSTANCE_LIST)

	r := gin.Default()
	RouterSetUp(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/lb_api/reschedule?instance_id=service_a_2", nil)
	r.ServeHTTP(w, req)

	fmt.Println("after action:")
	PrintInstanceListPretty(load_balancer.INSTANCE_LIST)
}

func TestGetReplicaNum(t *testing.T) {
	r := gin.Default()
	RouterSetUp(r)

	var w *httptest.ResponseRecorder
	var req *http.Request

	fmt.Println("error case:")
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/lb_api/get_replica", nil)
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())

	fmt.Println("success case:")
	configFile := "../config/config_test.json"
	config.CONFIG_PARAMS = config.ConfigSetUp(configFile)
	load_balancer.INSTANCE_LIST = config.CONFIG_PARAMS.InstanceList
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/lb_api/get_replica", nil)
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
}

func PrintInstanceListPretty(list *load_balancer.InstanceList) {
	for _, instance := range list.Instances {
		fmt.Printf("%+v\n", instance)
	}
	fmt.Println("Replica Num: ", list.ReplicaNum)
	fmt.Println("Total: ", list.Total)
	fmt.Println("Instance Map: ", list.InstanceMap)
}
