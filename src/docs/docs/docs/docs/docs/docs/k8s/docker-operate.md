---
title: docker-operate
date: '2020-01-04T14:53:27.000Z'
tags:
  - k8s
---

# docker-operate

## 停止所有的container

```text
docker stop $(docker ps -q)
```

## 清除 所有的docker container

```text
docker rm $(docker ps -a -q)
```

## 清除 所有的磁盘挂载

```text
docker volume prune
```

## without sudo

```bash
   newgrp docker  
   sudo groupadd docker
   docker run hello-world
   sudo usermod -aG docker ${USER}
   docker run hello-world
   sudo systemctl restart docker
   docker run hello-world
```

