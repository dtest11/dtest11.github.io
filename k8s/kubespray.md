declare -a IPS=(192.168.1.33 192.168.1.34 192.168.1.35)

CONFIG_FILE=inventory/mycluster/hosts.yml python3 contrib/inventory_builder/inventory.py ${IPS[@]}


https://www.cnblogs.com/huanmin/p/16819079.html