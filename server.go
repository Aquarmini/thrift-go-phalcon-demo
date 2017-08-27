package main

import (
	"fmt"
	//"os"
	"phalcon/demo/micro/impl"
	"phalcon/demo/micro/service"
	"thrift"
	"os"
)

const (
	NetworkAddr = "0.0.0.0:10086"
)

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
	//server = thrift.NewTSimpleServer4(processor1, serverTransport, transportFactory, protocolFactory)

	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()

}
