```bash

etcdctl  -w table --cacert=/etc/ssl/etcd/ssl/ca.pem --cert=/etc/ssl/etcd/ssl/node-node1.pem --key=/etc/ssl/etcd/ssl/node-node1-key.pem --endpoints=https://192.168.1.121:2379,https://192.168.1.122:2379,https://192.168.1.120:2379 \
endpoint status --cluster



etcdctl  --cacert=/etc/ssl/etcd/ssl/ca.pem --cert=/etc/ssl/etcd/ssl/node-node1.pem --key=/etc/ssl/etcd/ssl/node-node1-key.pem --endpoints=https://192.168.1.121:2379,https://192.168.1.122:2379,https://192.168.1.120:2379 \
-u root:123456 auth disable


etcdctl   --cacert=/etc/ssl/etcd/ssl/ca.pem --cert=/etc/ssl/etcd/ssl/node-node1.pem --key=/etc/ssl/etcd/ssl/node-node1-key.pem --endpoints=https://192.168.1.121:2379,https://192.168.1.122:2379,https://192.168.1.120:2379 \
get --prefix --limit=2 


```

```bash

ResolvePools
pl, err := c.IPPools().List(ctx, options.ListOptions{})
type ipPools struct {
	client client
}.List()

if err := r.client.resources.List(ctx, opts, apiv3.KindIPPool, apiv3.KindIPPoolList, res); err != nil {
		return nil, err
}

func (c *resources) List(ctx context.Context, opts options.ListOptions, kind, listKind string, listObj resourceList) error {
    ... 
    kvps, err := c.backend.List(ctx, list, opts.ResourceVersion) }


func (c *customK8sResourceClient) List(ctx context.Context, list model.ListInterface, revision string) (*model.KVPairList, error) {






```