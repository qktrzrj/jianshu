package resolver

import (
	"context"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"google.golang.org/grpc/resolver"
	"sync"
)

var defaultPrefix = "hello/grpc"

type builder struct {
	prefix string
	client clientv3.Client
}

func (b *builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	b.prefix = defaultPrefix + "/" + target.Endpoint
	adders, err := b.parseTarget()
	if err != nil {
		return nil, err
	}

	var rds []resolver.Address
	for _, a := range adders {
		rds = append(rds, resolver.Address{Addr: a})
	}
	cc.UpdateState(resolver.State{Addresses: rds})
	ctx, cancel := context.WithCancel(context.Background())
	e := &etcdResolver{
		adders: rds,
		ctx:    ctx,
		cancel: cancel,
		cc:     cc,
		rn:     make(chan struct{}, 1),
		b:      b,
	}

	e.wg.Add(1)
	go e.watcher()
	e.ResolveNow(resolver.ResolveNowOptions{})
	return e, nil
}

func (b *builder) Scheme() string {
	return "etcd"
}

type etcdResolver struct {
	adders []resolver.Address
	ctx    context.Context
	cancel context.CancelFunc
	cc     resolver.ClientConn
	rn     chan struct{}
	wg     sync.WaitGroup
	b      *builder
}

func (e *etcdResolver) ResolveNow(options resolver.ResolveNowOptions) {
	select {
	case e.rn <- struct{}{}:
	default:
	}
}

func (e *etcdResolver) Close() {
	e.cancel()
	e.wg.Wait()
}

func (e *etcdResolver) watcher() {
	defer e.wg.Done()

	wch := e.b.client.Watch(e.ctx, e.b.prefix, clientv3.WithPrefix())
	for {
		select {
		case <-e.ctx.Done():
			return
		case wr := <-wch:
			for _, ev := range wr.Events {
				switch ev.Type {
				case mvccpb.PUT:
					e.adders = append(e.adders)
					e.cc.UpdateState(resolver.State{
						Addresses: e.adders,
					})
				case mvccpb.DELETE:
					for i, a := range e.adders {
						if a.Addr == string(ev.Kv.Value) {
							e.adders = append(e.adders[:i], e.adders[i+1:]...)
							break
						}
					}
				}
			}
		}
	}
}

func (b *builder) parseTarget() (adders []string, err error) {
	if b.prefix == "" {
		return nil, errors.New("etcd resolver: must provider target serverName")
	}
	resp, err := b.client.Get(context.Background(), b.prefix, clientv3.WithPrefix())
	if err != nil {
		return
	}
	adders = extractAdders(resp)
	return
}

func extractAdders(resp *clientv3.GetResponse) []string {
	var addrs []string
	if resp == nil || resp.Count == 0 {
		return addrs
	}
	for k := range resp.Kvs {
		if v := resp.Kvs[k].Value; v != nil {
			addrs = append(addrs, string(v))
		}
	}
	return addrs
}
