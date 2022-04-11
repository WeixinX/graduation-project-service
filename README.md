# Deploy

部署服务环境，其中包括 minikube、tracing、storage、monitoring、demo-service 和 lb-service

## Minikube

暴露 k8s dashboard

```shell
// sh 命令在 在 [WeixinX/graduation-project](https://github.com/WeixinX/graduation-project) 的 scripts/minikube 中
$ cd /home/ubuntu/goworkspace/src/WeixinX/graduation-project

$ sh ./scripts/minikube/expose_dashboard.sh
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
$ sudo systemctl restart kubelet
```

## Tracing & Storage

启动 Jaeger, Elasticsearch 和 Kibana

```shell
// 文件在 [WeixinX/graduation-project](https://github.com/WeixinX/graduation-project) 的 manifest/tracing 中
$ cd /home/ubuntu/goworkspace/src/WeixinX/graduation-project

// tracing
$ kubectl create namespace tracing
$ kubectl apply -f ./manifest/tracing/jaeger-all.yml
$ sh ./scripts/minikube/expose_jaeger.sh

// elastic for storing trace data
$ kubectl create namespace elastic
$ kubectl apply -f ./manifest/tracing/es.yml
$ kubectl apply -f ./manifest/tracing/kibana.yml
$ sh ./scripts/minikube/expose_elastic.sh
```

## Monitoring

启动 Prometheus 和 Grafana

```shell
// 文件在 [WeixinX/graduation-project](https://github.com/WeixinX/graduation-project) 的 manifest/monitoring 中
$ cd /home/ubuntu/goworkspace/src/WeixinX/graduation-project

$ kubectl create namespace monitoring
// Prometheus
$ kubectl apply -f ./manifest/monitoring/prometheus-config.yaml
$ kubectl apply -f ./manifest/monitoring/prometheus-deployment.yaml
$ kubectl apply -f ./manifest/monitoring/prometheus-service.yaml

// Grafana
$ kubectl apply -f ./manifest/monitoring/grafana-deployment.yaml
$ kubectl apply -f ./manifest/monitoring/grafana-service.yaml
```

## Service

将所有服务打包成 docker image, 将他们部署到 minikube 环境中

### 构建服务二进制文件

```shell
$ cd /home/ubuntu/goworkspace/src/WeixinX/graduation-project-service

$ sh ./scripts/build_all_services.sh
```

### 构建服务镜像

```shell
$ sh ./scripts/build_all_images.sh minikube
```

### 部署服务到 minikube 环境

```shell
$ kubectl create namespace service
$ kubectl apply -f ./manifests/kubernetes/demo-and-lb-deployment.yml
$ kubectl apply -f ./manifests/kubernetes/demo-and-lb-service.yml
```

### 暴露服务

```shell
$ sh ./scripts/expose_all_service.sh
```
