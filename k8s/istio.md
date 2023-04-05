- xDS协议

xDS协议是Envoy获取配置信息的传输协议,grpc-stream
ADS是一种xDS的实现, 它基于gRPC长连接。 gRPC的实现是承载在HTTP/2之上。

- 流量镜像
- 请求路由管理-负载均衡，重试，熔断
- 调试
- 安全
- 可观测性
- 故障注入

* bug
请求中断

* [install](https://istio.io/latest/zh/docs/setup/getting-started/)

```bash
 kubectl get ingress --all-namespaces
 kubectl get gateway --all-namespaces

 ```
 卸载
 ```bash
 https://istio.io/latest/zh/docs/setup/install/istioctl/
 ```

virtual service 搭配gateway 使用，virtual service 定义各种规则，gateway选择pod