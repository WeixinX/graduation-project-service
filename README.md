# Deploy

部署服务环境，其中包括 minikube、tracing、storage、monitoring、demo-service 和 lb-service

## Pre-work

需要先安装:
- [docker](https://docs.docker.com/engine/install/ubuntu/#install-using-the-convenience-script)
- [minikube](https://minikube.sigs.k8s.io/docs/start)
- [kubectl](https://kubernetes.io/zh/docs/tasks/tools/install-kubectl-linux/)

## Minikube

暴露 k8s dashboard

```shell
$ cd /home/ubuntu/goworkspace/src/WeixinX/graduation-project-service

$ chmod +777 ./scripts/minikube/* && ./scripts/minikube/expose_dashboard.sh
```

调整 minikube kubelet 的 `--housekeeping_interval` (kubelet 采集 cAdvisor 数据的周期)

```shell
// 1. 进入 minikube
$ minikube ssh

// 2. 查看 kubelet 运行情况
$ systemctl status kubelet

/* 可以看到运行选项中 --housekeeping-interval=5m
● kubelet.service - kubelet: The Kubernetes Node Agent
     Loaded: loaded (/lib/systemd/system/kubelet.service; disabled; vendor preset: enabled)
    Drop-In: /etc/systemd/system/kubelet.service.d
             └─10-kubeadm.conf
     Active: active (running) since Wed 2022-04-06 13:41:20 UTC; 11min ago
       Docs: http://kubernetes.io/docs/
   Main PID: 3118664 (kubelet)
      Tasks: 33 (limit: 9540)
     Memory: 85.1M
     CGroup: /docker/46dfcefbb539255399f2143993306cbc7db4e7d436a722beaecd0a058aff68d5/system.slice/kubelet.service
             ├─3118664 /var/lib/minikube/binaries/v1.23.3/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --config=/var/lib/kubelet/config.yaml --container-runtime=docker --hostname-override=
minikube --housekeeping-interval=5m --kubeconfig=/etc/kubernetes/kubelet.conf --node-ip=192.168.49.2

*/

// 3. 修改 --housekeeping-interval, 例如: --housekeeping-interval=5s
$ sudo vi /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

// 4. 重启 kubelet
$ sudo systemctl daemon-reload
$ sudo systemctl restart kubelet
```

## Tracing & Storage

启动 Jaeger、Elasticsearch 和 Kibana

```shell
// tracing
$ kubectl create namespace tracing
$ kubectl apply -f ./manifests/kubernetes/tracing/jaeger-all.yml
$ ./scripts/minikube/expose_jaeger.sh

// elastic for storing trace data
$ kubectl create namespace elastic
$ kubectl apply -f ./manifests/kubernetes/tracing/es.yml
$ kubectl apply -f ./manifests/kubernetes/tracing/kibana.yml
$ ./scripts/minikube/expose_elastic.sh
```

## Monitoring

启动 Prometheus 和 Grafana

```shell
$ kubectl create namespace monitoring

// Prometheus
$ kubectl apply -f ./manifests/kubernetes/monitoring/prometheus-config.yaml
$ kubectl apply -f ./manifests/kubernetes/monitoring/prometheus-deployment.yaml
$ kubectl apply -f ./manifests/kubernetes/monitoring/prometheus-service.yaml

// Grafana
$ kubectl apply -f ./manifests/kubernetes/monitoring/grafana-deployment.yaml
$ kubectl apply -f ./manifests/kubernetes/monitoring/grafana-service.yaml

// expose
$ ./scripts/minikube/expose_monitoring.sh
```

## Service

将所有服务打包成 docker image, 将他们部署到 minikube 环境中

### 构建服务二进制文件

```shell
$ cd ./scripts && chmod +777 ./*
$ ./build_all_services.sh
```

### 构建服务镜像

```shell
$ ./build_all_images.sh minikube
```

### 部署服务到 minikube 环境

```shell
$ kubectl create namespace service
$ kubectl apply -f ./manifests/kubernetes/service/demo-and-lb-deployment.yml
$ kubectl apply -f ./manifests/kubernetes/service/demo-and-lb-service.yml
```

### 暴露服务

```shell
$ ./expose_all_service.sh
```
