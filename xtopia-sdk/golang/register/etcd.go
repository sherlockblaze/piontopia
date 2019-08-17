package register

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// EtcdServiceRegister etcd register
type EtcdServiceRegister struct {
	EtcdAddr    string
	ServiceName string
	ServiceAddr string
	Timeout     uint32
	TTL         int64
}

// Register Etcd service register
func (sr *EtcdServiceRegister) Register() error {
	var err error

	if cli == nil {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(sr.EtcdAddr, ";"),
			DialTimeout: time.Duration(sr.Timeout) * time.Second,
		})
		if err != nil {
			// TODO: log
			return err
		}
	}

	ticker := time.NewTicker(time.Second * time.Duration(sr.TTL))

	go func() {
		for {
			getResp, err := cli.Get(context.Background(), "/"+schema+"/"+sr.ServiceName+"/"+sr.ServiceAddr)

			if err != nil {
				log.Println(err)
				fmt.Printf("Register:%s", err)
			} else if getResp.Count == 0 {
				err = withAlive(sr.ServiceName, sr.ServiceAddr, sr.TTL)
				if err != nil {
					// TODOL: log
				}
			} else {
				// do nothing now
			}

			<-ticker.C
		}
	}()

	return err
}

func withAlive(name string, addr string, ttl int64) error {
	leaseResp, err := cli.Grant(context.Background(), ttl)
	if err != nil {
		return err
	}

	_, err = cli.Put(context.Background(), "/"+schema+"/"+name+"/"+addr, addr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		fmt.Printf("put etcd error:%s", err)
		return err
	}

	_, err = cli.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		fmt.Printf("keep alive error:%s", err)
		return err
	}
	return nil
}

// UnRegister remove service from etcd
func (sr *EtcdServiceRegister) UnRegister() error {
	if cli != nil {
		cli.Delete(context.Background(), "/"+schema+"/"+sr.ServiceName+"/"+sr.ServiceAddr)
	}
	return nil
}
