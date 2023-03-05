package main

import (
	"context"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/youlingdada/nacos-demo/grpc/user"
	"github.com/youlingdada/nacos-demo/naming"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func (us *UserService) Login(ctx context.Context, request *user.LoginRequest) (*user.LoginResponse, error) {
	log.Printf("登录成功,username:%s", request.Username)
	return &user.LoginResponse{Code: 200, Message: "登录成功", Data: "this is a user"}, nil
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func main() {
	port := 8989
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &UserService{})
	log.Printf("server listening at %v", lis.Addr())

	// 注册服务
	naming.InitNacos()

	ip, err := GetOutBoundIP()
	if err != nil {
		log.Printf("获取主机ip失败,err:%v", err)
	}
	instance, err := naming.Client.RegisterInstance(vo.RegisterInstanceParam{
		ServiceName: "userService",
		GroupName:   "user",
		Enable:      true,
		Healthy:     true,
		Ip:          ip,
		Port:        uint64(port),
		Weight:      10,
	})
	if err != nil {
		log.Printf("注册服务失败")
	} else {
		log.Printf("注册服务状态：%v,服务名称：%s,服务ip: %s,服务端口：%d", instance, "userService", ip, port)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
