```text

docker run -d --restart=unless-stopped \
  -p 8090:80 -p 444:443 \
  --privileged \
  rancher/rancher:v2.6.9
  
```
 docker exec -ti interesting_brown reset-password


## 查看pod 状态

```bash
kubectl -n  kube-ops describe   pod runner-svzx4ujw-project-1-concurrent-0qd6hq
```

## 删除pod

kubectl -n kube-ops delete pods runner-svzx4ujw-project-1-concurrent-0qd6hq --grace-period=0 --force

## 查看pod 中的container

kubectl get pods --all-namespaces -o=jsonpath='{range .items\[_\]}{"\n"}{.metadata.name}{":\t"}{range .spec.containers\[_\]}{.image}{", "}{end}{end}' \| sort

## 查看pod 中container 的日志

kubectl logs  -c 


## 重装
kubeadm reset



## nginx-ingress-controller
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.7.0/deploy/static/provider/baremetal/deploy.yaml
```
image mirror
```
docker pull m.daocloud.io/registry.k8s.io/ingress-nginx/kube-webhook-certgen:v20230312-helm-chart-4.5.2-28-g66a760794@sha256:01d181618f270f2a96c04006f33b2699ad3ccb02da48d0f89b22abce084b292f

```