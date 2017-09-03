package main

import (
	"fmt"
	"phalcon/demo/micro/impl"
	"phalcon/demo/micro/service"
	"thrift"
	"os"
)
import log "github.com/sirupsen/logrus"

const (
	NetworkAddr = "0.0.0.0:10086"
)

func init() {
	//go get -u github.com/sirupsen/logrus
	fmt.Println("INIT")
	file, _ := os.OpenFile("go.server.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	log.SetOutput(file);
}

func main() {
	//transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	processor := thrift.NewTMultiplexedProcessor();
	processor.RegisterProcessor("app", service.NewAppProcessor(&impl.App{}));
	processor.RegisterProcessor("user", service.NewUserProcessor(&impl.User{}));
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)

	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()

}
