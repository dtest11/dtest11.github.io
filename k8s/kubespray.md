declare -a IPS=(192.168.1.33 192.168.1.34 192.168.1.38)

CONFIG_FILE=inventory/mycluster/hosts.yml python3 contrib/inventory_builder/inventory.py ${IPS[@]}

ansible-playbook -i inventory/mycluster/hosts.yml  --become --become-user=root cluster.yml

https://www.cnblogs.com/huanmin/p/16819079.html

固定虚拟机ip(https://www.cnblogs.com/shangxb/p/virtualbox-set-ubuntu-2004-static-ip.html)

```yaml

 ansible_user: dev
 ansible_sudo_pass: 123456

```


```bash

mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

  ```