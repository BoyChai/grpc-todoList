package discovery

import (
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/resolver"
)

type Resolver struct {
	schema      string
	EtcdAddres  []string
	DialTimeout int

	closeCh      chan struct{}
	watchCh      clientv3.WatchChan
	cli          *clientv3.Client
	keyPrifix    string
	srvAddrsList []resolver.Address

	cc     resolver.ClientConn
	logger *logrus.Logger
}
