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

### flannel


不同node上的pod的通信流程：

1. pod中产生数据，根据pod的路由信息，将数据发送到Cni0
2. Cni0 根据节点的路由表，将数据发送到隧道设备flannel.1
3. Flannel.1查看数据包的目的ip，从flanneld获得对端隧道设备的必要信息，封装数据包。
4. Flannel.1将数据包发送到对端设备。对端节点的网卡接收到数据包，发现数据包为overlay数据包，解开外层封装，并发送内层封装到flannel.1设备。
5.Flannel.1设备查看数据包，根据路由表匹配，将数据发送给Cni0设备。
6. Cni0匹配路由表，发送数据给网桥上对应的端口。



