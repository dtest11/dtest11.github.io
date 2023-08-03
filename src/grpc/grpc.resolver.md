服务发现与负载均衡
常用策略：

1. RoundRobin（轮询）
2. Weight-RoundRobin（加权轮询）
不同的后端服务器可能机器的配置和当前系统的负载并不相同，因此它们的抗压能力也不相同。给配置高、负载低的机器配置更高的权重，而配置低、负载高的机器，给其分配较低的权重，降低其系统负载，加权轮询能很好地处理这一问题，并将请求顺序且按照权重分配到后端。
3. Random（随机）
4. Weight-Random（加权随机）
通过系统的随机算法，根据后端服务器的列表随机选取其中的一台服务器进行访问
源地址哈希法
源地址哈希的思想是根据获取客户端的 IP 地址，通过哈希函数计算得到的一个数值，用该数值对服务器列表的大小进行取模运算，得到的结果便是客服端要访问服务器的序号。采用源地址哈希法进行负载均衡，同一 IP 地址的客户端，当后端服务器列表不变时，它每次都会映射到同一台后端服务器进行访问
5. 最小连接数法
最小连接数算法比较灵活和智能，由于后端服务器的配置不尽相同，对于请求的处理有快有慢，它是根据后端服务器当前的连接情况，动态地选取其中当前积压连接数最少的一台服务器来处理当前的请求，尽可能地提高后端服务的利用效率，将负责合理地分流到每一台服务器
6. 一致性哈希算法
常见的是 Ketama 算法，该算法是用来解决 cache 失效导致的缓存穿透的问题的，当然也可以适用于 gRPC 长连接的场景

## dns 
[https://github.com/grpc/grpc-go/blob/master/internal/resolver/dns/dns_resolver.go](https://github.com/grpc/grpc-go/blob/master/internal/resolver/dns/dns_resolver.go)


## etcd 服务发现

```go


package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"time"

	"google.golang.org/grpc"
)

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"http://s1004.lab.org:2379"},
		DialTimeout: 5 * time.Second,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}

	etcdResolver, err := resolver.NewBuilder(cli)
	conn, err := grpc.Dial("etcd:///foo/bar/my-service", grpc.WithResolvers(etcdResolver))
	if err != nil {
		fmt.Printf("dial service(%s) by etcd resolver server error (%v)", addr, err.Error())
		panic(err)
	}
	defer conn.Close()
}

func Register() {
	config := clientv3.Config{
		Endpoints:   []string{"http://s1004.lab.org:2379"},
		DialTimeout: 5 * time.Second,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}
	em, err := endpoints.NewManager(cli, "foo/bar/my-service")
	if err != nil {
		panic(err)
	}
	err = em.AddEndpoint(context.TODO(), "foo/bar/my-service/e1", endpoints.Endpoint{Addr: "1.2.3.4"})
	if err != nil {
		panic(err)
	}
}


```

[https://etcd.io/docs/v3.3/dev-guide/grpc_naming/](https://etcd.io/docs/v3.3/dev-guide/grpc_naming/)

[https://pandaychen.github.io/2019/07/11/GRPC-SERVICE-DISCOVERY/](https://pandaychen.github.io/2019/07/11/GRPC-SERVICE-DISCOVERY/)
