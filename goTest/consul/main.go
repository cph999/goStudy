package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"net"
)

type consul struct {
	client *api.Client
}

var consulInstance *consul

// NewConsul 连接至consul服务返回一个consul对象
func NewConsul(addr string) (*consul, error) {
	cfg := api.DefaultConfig()
	cfg.Address = addr
	c, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &consul{c}, nil
}

func initConsul(address string) (*consul, error) {
	return NewConsul(address)
}

// GetOutboundIP 获取本机的出口IP
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

// RegisterService 将gRPC服务注册到consul
func (c *consul) RegisterService(serviceName string, ip string, port int) error {
	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port), // 服务唯一ID
		Name:    serviceName,                                    // 服务名称
		Tags:    []string{"q1mi", "hello"},                      // 为服务打标签
		Address: ip,
		Port:    port,
	}
	return c.client.Agent().ServiceRegister(srv)
}

// ListService 服务发现
func (c *consul) ListService(serviceName string) (map[string]*api.AgentService, error) {
	// c.client.Agent().Service("hello-127.0.0.1-8972")
	return c.client.Agent().ServicesWithFilter("Service== " + serviceName)
}

// Deregister 注销服务
func (c *consul) Deregister(serviceID string) error {
	return c.client.Agent().ServiceDeregister(serviceID)
}

func main() {
	consulInstance, err := initConsul("121.196.223.94:8500")
	if err != nil {
		return
	}
	err = consulInstance.RegisterService("stronger", "1.2.3.4", 12345)
	if err != nil {
		return
	}
	services, err := consulInstance.ListService("stronger")
	if err != nil {
		return
	}
	for key, value := range services {
		fmt.Println(key, value)
	}
}
