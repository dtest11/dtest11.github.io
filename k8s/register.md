```json
{
  "registry-mirrors": [
    "http://hub-mirror.com:5000"
  ],
  "insecure-registry": [
    "http://hub-mirror.com:5000"
  ]
}
```

 sudo vim etc/docker/daemon.json 
 sudo systemctl daemon-reload 
 sudo systemctl restart docker
 sudo systemctl status docker

curl http://hub-mirror.com:5000/v2/library/redis/tags/list




docker run  --restart always \
--name registry-proxy \
-d \
-v /home/bong/data:/var/lib/registry \
-v /home/bong/config.yaml:/etc/docker/registry/config.yml \
-p 5000:5000 registry

sudo vim  /etc/containerd/config.toml 
sudo systemctl restart containerd.service
sudo systemctl status containerd.service

 curl http://127.0.0.1:5000/v2/_catalog
