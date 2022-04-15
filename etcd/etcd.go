package etcd

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Etcd struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var (
	Cli      *clientv3.Client
	DataChan chan *Etcd
)

//初始化etcd
func Init() (err error) {
	Cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return err
	}
	DataChan = make(chan *Etcd, 10000)
	go SendToEtcd()
	return err
}

//给外部暴露一个chan向里面写如数据
func SendToChan(key, vlaue string) {
	msg := &Etcd{
		Key:   key,
		Value: vlaue,
	}
	DataChan <- msg
}

//从chan中得到数据并写入etcd中
func SendToEtcd() {
	for {
		select {
		case data := <-DataChan:
			//将数据写入etcd中
			_, err := Cli.Put(context.TODO(), data.Key, data.Value)
			if err != nil {
				fmt.Println("send data failed err", err)
				return
			}

		default:
			time.Sleep(time.Second)
		}
	}

}

//返回数据给页面
func GetData() (data []*Etcd) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := Cli.Get(context.TODO(), "", clientv3.WithPrefix())
	// cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return nil
	}
	for _, ev := range resp.Kvs {
		data = append(data, &Etcd{
			Key:   string(ev.Key),
			Value: string(ev.Value),
		})
	}
	return data
}

//关闭etcd
func CloseEtcd(cli *clientv3.Client) {
	cli.Close()
}
