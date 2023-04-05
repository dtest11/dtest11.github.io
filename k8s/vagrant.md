```bash

vagrant init hashicorp/bionic64
vagrant box  add bento/ubuntu-18.04
vagrant up --provider=virtualbox

重新加载
vagrant reload
stop:

vagrant halt  

```

bridge init
```bash
master.vm.network "public_network",bridge: "wlp58s0",ip: "192.168.1.120"

```
config like this to get name 
#### yaml 内容
Vagrantfile.yaml
hyperv

