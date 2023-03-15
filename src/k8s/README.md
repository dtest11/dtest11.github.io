# k8s

##  istio
1. [https://jimmysong.io/istio-handbook/concepts/what-is-service-mesh.html](https://jimmysong.io/istio-handbook/concepts/what-is-service-mesh.html)


* pod
* deployment
* service 
* DaemonSet(每个Node上运行一个 Pod 的副本，日志收集等)
* StatefulSet
* Federation
* Volume
* Persistent Volume，PV/pvc
* node
* secret
* user account /sevice account
* namespace
* rabc
* etcd
* CRI/CNI/CSI
### network
cacio
flannel

### Pause / infra
linux namespce / cgroups

## pod 
### 生命周期： 
* pending
* running
* successed
* failed
* unknown

### 探针
* exec
* tcp socket
* http get 

### pod hook

## Node
## label
* lable selector
## annotation
## taint / toleration(污点/容忍)

## statefulset

StatefulSet 适用于有以下某个或多个需求的应用：

    稳定，唯一的网络标志。
    稳定，持久化存储。
    有序，优雅地部署和 scale。
    有序，优雅地删除和终止。
    有序，自动的滚动升级。

## Job
## CronJob
## HPA

## service
### iptables
### kube-proxy
### ipvs :hash 数据结构
### cluster ip
### Node port
### loadbalancer
### externalName

## ingress

### ingress controller
### traefik ingress controller


## service account 
### RBAC








