package client

import (
	"context"
	"time"

	"github.com/leechanx/ekko-idgenerator/protobuf"

	"github.com/smallnest/rpcx/client"
)

type EkkoClient struct {
	xclient client.XClient
	option  option
}

func NewEkkoClient(addrs []string, opts ...OpOption) (*EkkoClient, error) {
	var kvpars []*client.KVPair
	for _, addr := range addrs {
		kvpars = append(kvpars, &client.KVPair{
			Key: addr,
		})
	}
	ekkoClient := &EkkoClient{}
	d, err := client.NewMultipleServersDiscovery(kvpars)
	if err != nil {
		return ekkoClient, err
	}

	rpcxOption := client.DefaultOption
	rpcxOption.IdleTimeout = time.Millisecond * 100
	rpcxOption.Retries = 2

	ekkoClient.xclient = client.NewXClient("Ekko", client.Failfast, client.RoundRobin, d, rpcxOption)
	for _, optFunc := range opts {
		optFunc(&ekkoClient.option)
	}
	return ekkoClient, nil
}

func (e *EkkoClient) Close() {
	e.xclient.Close()
}

func (c *EkkoClient) IDGen(ctx context.Context, product int32) (uint64, error) {
	if c.option.fallback {
		if len(c.option.uidList) == 0 {
			// default count = 0 will get 1024 results
			uids, err := c.MultiIDGen(ctx, product, 0)
			if err == nil {
				c.option.uidList = uids
			} else {
				return 0, err
			}
		}
		uid := c.option.uidList[0]
		c.option.uidList = c.option.uidList[1:]
		return uid, nil
	}
	req := protobuf.GetUniqueIDRequest{Product: product}
	rsp := &protobuf.GetUniqueIDResponse{}
	err := c.xclient.Call(ctx, "GetUniqueID", req, rsp)
	if err != nil {
		return 0, err
	}
	return rsp.Uid, nil
}

func (c *EkkoClient) MultiIDGen(ctx context.Context, product int32, count uint32) (uids []uint64, err error) {
	req := &protobuf.MGetUniqueIDRequest{
		Product: product,
		Count:   count,
	}
	rsp := &protobuf.MGetUniqueIDResponse{}
	err = c.xclient.Call(ctx, "MGetUniqueID", req, rsp)
	if err != nil {
		return nil, err
	}
	return spreadOutUidList(rsp.LowerUid, rsp.UpperUid), nil
}
