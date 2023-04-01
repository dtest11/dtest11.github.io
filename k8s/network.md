### network namespce 

* veth pair: virtual ethernet
Veth设备的作用主要用来连接两个网络命名空间

* linux bridge

brctl show: 查看网桥
docker network inspect bridge

docker0的主要作用是为Docker容器提供网络隔离和连接

veth pair 一端连接在容器，一端连接在docker0上

* tun/tap 
- tun 虚拟的点对点设备
- tap 虚拟的以太网设备
可以将TCP/IP协议栈处理好的网络包发送给任何一个使用tun/tap驱动的进程，由进程重新处理后发到物理链路中。
tun/tap设备就像是埋在用户程序空间的一个钩子。
它与物理网卡的不同表现在它的数据源不是物理链路，而是来自用户态！这也是tun/tap设备的最大价值所在。
提前“剧透”：flannel的UDP模式的技术要点就是tun/tap设备。


* ipip

IPV In ipv4

* vxlan
虚拟可扩展的局域网

* macvlan
