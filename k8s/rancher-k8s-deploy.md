---
title: rancher-k8s-deploy
date: '2020-01-03T23:01:27.000Z'
tags:
  - k8s
---

# rancher-k8s-deploy

![&#x56FE;&#x7247;&#x540D;&#x79F0;](https://lin19999.oss-cn-beijing.aliyuncs.com/rancher.svg)

## Rancher install

如果之前安装过,需要清理 /opt/rancher 文件夹

```text
docker run -d -p 80:80 -p 443:443 \
  --restart=unless-stopped \
  -v /opt/rancher:/var/lib/rancher \
  rancher/rancher:latest
```

## 选择

![&#x56FE;&#x7247;&#x540D;&#x79F0;](https://lin19999.oss-cn-beijing.aliyuncs.com/rancher_dep.png)

