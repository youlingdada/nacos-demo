# nacos-demo
nacos go-sdk实践，使用nacos go sdk 实现配置管理，服务治理。同通过nacos 服务治理与grpc实现分布式微服务

### nacos 配置中心和服务治理
* 在代码中配置ip和端口
* 启动根目录的main.go
* 相关接口测试(部分接口)
  * 发布配置
    ```shell
    http://localhost:8080/config/publish?dataId=test-data&group=test-group&content=hello%20world
    ```
  * 获取配置
    ```shell
    http://localhost:8080/config/get?dataId=test-data&group=test-group
    ```
  * 注册服务
    ```shell
    http://localhost:8080/naming/register?ip=5.5.5.5&port=8080&serviceName=test-service&groupName=test-group&weight=10&enable=true&healthy=true
    ```
  * 获取服务
    ```shell
    http://localhost:8080/naming/getService?serviceName=test-service&groupName=test-group
    ```

### grpc 使用nacos作为服务注册中心
* 服务端：
```shell
go run ./grpc/user/service/server.go
```
* 客户端
```shell
go run ./grpc/user/service/client.go
```