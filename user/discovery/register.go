package discovery

import (
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
)

type Register struct {
	EtcdAddres  []string
	DialTimeout int
	closeCh     chan struct{}
	leasesID    clientv3.LeaseID
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse

	srvInfo Server
	SrvTTL  int64
	cki     *clientv3.Client
	logger  *logrus.Logger
}
