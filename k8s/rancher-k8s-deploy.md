# rancher-k8s-deploy


## Rancher install

如果之前安装过,需要清理 /opt/rancher 文件夹

```text
docker run -d -p 8080:80 -p 443:443 \
  --restart=unless-stopped \
  -v /opt/rancher:/var/lib/rancher \
  --privileged \
  rancher/rancher:latest
```

## 选择

![&#x56FE;&#x7247;&#x540D;&#x79F0;](https://lin19999.oss-cn-beijing.aliyuncs.com/rancher_dep.png)

