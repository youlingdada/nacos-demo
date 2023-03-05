package main

import (
	"context"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/youlingdada/nacos-demo/grpc/user"
	"github.com/youlingdada/nacos-demo/naming"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	// 获取服务
	serviceName := "userService"
	groupName := "user"

	naming.InitNacos()
	instances, err := naming.Client.SelectInstances(vo.SelectInstancesParam{
		ServiceName: serviceName,
		GroupName:   groupName,
	})
	if err != nil {
		log.Fatalf("获取服务失败,err:%v", err)
	}

	if len(instances) == 0 {
		log.Fatalf("没有找到相应的服务,err:%v", err)
	}
	instance := instances[0]
	log.Printf("选择到一个服务,ip:%s,port:%d", instance.Ip, instance.Port)
	addr := fmt.Sprintf("%s:%d", instance.Ip, instance.Port)

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Login(ctx, &user.LoginRequest{Username: "yxd", Password: "123456"})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Printf("login info: %v", r)
}
