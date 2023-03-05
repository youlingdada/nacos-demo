package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
	"strconv"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/config/publish", PublishConfig)
	router.GET("/config/get", GetConfig)
	router.GET("/config/addListen", AddConfigListen)
	router.GET("/config/cancelListen", CancelConfigListen)
	router.GET("/config/del", DelConfig)
	router.GET("/config/search", SearchConfig)

	router.GET("/naming/register", RegisterServiceInstance)
	router.GET("naming/update", UpdateServiceInstance)
	router.GET("/naming/getInstance", GetServiceInstance)
	router.GET("/naming/getAllInstance", GetAllServiceInstance)
	router.GET("/naming/getService", GetService)
	router.GET("/naming/getAllService", GetAllService)
	router.GET("naming/subscribe", Subscribe)
	router.GET("naming/unSubscribe", UnSubscribe)

	return router
}

// PublishConfig 发布配置
// http://localhost:8080/config/publish?dataId=test-data&group=test-group&content=hello%20world
func PublishConfig(c *gin.Context) {
	dataId := c.Query("dataId")
	group := c.Query("group")
	content := c.Query("content")

	config, err := ConfigClient.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: content,
	})
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", config)
	}
}

// GetConfig 获取配置
// http://localhost:8080/config/get?dataId=test-data&group=test-group
func GetConfig(c *gin.Context) {
	dataId := c.Query("dataId")
	group := c.Query("group")
	config, err := ConfigClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})

	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", config)
	}
}

// AddConfigListen 添加配置监听
// http://localhost:8080/config/addListen?dataId=test-data&group=test-group
func AddConfigListen(c *gin.Context) {
	dataId := c.Query("dataId")
	group := c.Query("group")

	err := ConfigClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("namespace: " + namespace + "config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})
	if err != nil {
		log.Printf("%v", err)
	}
}

// CancelConfigListen 取消配置监听
// http://localhost:8080/config/cancelListen?dataId=test-data&group=test-group
func CancelConfigListen(c *gin.Context) {
	dataId := c.Query("dataId")
	group := c.Query("group")

	err := ConfigClient.CancelListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		log.Printf("%v", err)
	}
}

// DelConfig 删除配置
// http://localhost:8080/config/del?dataId=test-data&group=test-group
func DelConfig(c *gin.Context) {
	dataId := c.Query("dataId")
	group := c.Query("group")
	config, err := ConfigClient.DeleteConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", config)
	}
}

// SearchConfig 搜索配置
// http://localhost:8080/config/search
func SearchConfig(c *gin.Context) {
	searchPage, _ := ConfigClient.SearchConfig(vo.SearchConfigParam{
		Search:   "blur",
		DataId:   "",
		Group:    "",
		PageNo:   1,
		PageSize: 10,
	})
	fmt.Printf("Search config:%+v \n", searchPage)
}

// RegisterServiceInstance 注册服务
// http://localhost:8080/naming/register?ip=5.5.5.5&port=8080&serviceName=test-service&groupName=test-group&weight=10&enable=true&healthy=true
func RegisterServiceInstance(c *gin.Context) {
	ip := c.Query("ip")
	port, _ := strconv.ParseInt(c.Query("port"), 10, 10)
	serviceName := c.Query("serviceName")
	groupName := c.Query("groupName")
	weight, _ := strconv.ParseInt(c.Query("weight"), 10, 10)
	enable, _ := strconv.ParseBool(c.Query("enable"))
	healthy, _ := strconv.ParseBool(c.Query("healthy"))
	registerServiceInstance(NamingClient, vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        uint64(port),
		ServiceName: serviceName,
		GroupName:   groupName,
		Weight:      float64(weight),
		Enable:      enable,
		Healthy:     healthy,
	})
}

// UpdateServiceInstance 更新服务示例
// http://localhost:8080/naming/update?ip=5.5.5.5&port=8080&serviceName=test-service&groupName=test-group&weight=10&enable=true&healthy=true
func UpdateServiceInstance(c *gin.Context) {
	ip := c.Query("ip")
	port, _ := strconv.ParseInt(c.Query("port"), 10, 10)
	serviceName := c.Query("serviceName")
	groupName := c.Query("groupName")
	weight, _ := strconv.ParseInt(c.Query("weight"), 10, 10)
	enable, _ := strconv.ParseBool(c.Query("enable"))
	healthy, _ := strconv.ParseBool(c.Query("healthy"))
	updateServiceInstance(NamingClient, vo.UpdateInstanceParam{
		Ip:          ip,
		Port:        uint64(port),
		ServiceName: serviceName,
		Weight:      float64(weight),
		Enable:      enable,
		GroupName:   groupName,
		Healthy:     healthy,
	})
}

// GetService 获取指定服务
// http://localhost:8080/naming/getService?serviceName=test-service&groupName=test-group
func GetService(c *gin.Context) {
	serviceName := c.Query("serviceName")
	groupName := c.Query("groupName")

	getService(NamingClient, vo.GetServiceParam{
		ServiceName: serviceName,
		GroupName:   groupName,
	})
}

// GetAllServiceInstance 获取注册的所有服务
// http://localhost:8080/naming/getAllInstance
func GetAllServiceInstance(c *gin.Context) {
	getAllService(NamingClient, vo.GetAllServiceInfoParam{})
}

// GetServiceInstance 获取所有服务实例
// 可以选择健康与否
// http://localhost:8080/naming/getInstance?iserviceName=test-service&groupName=test-group&healthy=true
func GetServiceInstance(c *gin.Context) {
	serviceName := c.Query("serviceName")
	groupName := c.Query("groupName")
	healthy, _ := strconv.ParseBool(c.Query("healthy"))
	selectInstances(NamingClient, vo.SelectInstancesParam{
		ServiceName: serviceName,
		GroupName:   groupName,
		HealthyOnly: healthy,
	})
}

// GetAllService
// 获取所有服务
// http://localhost:8080/naming/getAllService?serviceName=test-service&groupName=test-group
func GetAllService(c *gin.Context) {
	serviceName := c.Query("serviceName")
	groupName := c.Query("groupName")

	selectAllInstances(NamingClient, vo.SelectAllInstancesParam{
		ServiceName: serviceName,
		GroupName:   groupName,
	})
}

// Subscribe 订阅
// http://localhost:8080/naming/subscribe?serviceName=test-service&groupName=test-group
func Subscribe(c *gin.Context) {
	serviceName := c.Query("serviceName")
	groupName := c.Query("groupName")

	subscribe(NamingClient, &vo.SubscribeParam{
		ServiceName: serviceName,
		GroupName:   groupName,
		SubscribeCallback: func(services []model.Instance, err error) {
			log.Printf("instance : %v, err:%v", services, err)
		},
	})
}

// UnSubscribe 取消订阅
// http://localhost:8080/naming/unSubscribe?serviceName=test-service&groupName=test-group
func UnSubscribe(c *gin.Context) {
	serviceName := c.Query("serviceName")
	groupName := c.Query("groupName")
	unSubscribe(NamingClient, &vo.SubscribeParam{
		ServiceName: serviceName,
		GroupName:   groupName,
		SubscribeCallback: func(services []model.Instance, err error) {
			log.Printf("instance : %v, err:%v", services, err)
		},
	})
}

var ConfigClient config_client.IConfigClient
var NamingClient naming_client.INamingClient

// InitNacos 初始化nacos配置
// 保证8848 和 9848端口开放
func InitNacos() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("4.4.4.4", 8848, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create config client
	clientConfig, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	// create naming client
	clientNaming, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	ConfigClient = clientConfig
	NamingClient = clientNaming
}

func main() {
	InitNacos()
	router := InitRouter()

	err := router.Run(":8080")
	if err != nil {
		log.Printf("%v", err)
	}
}
