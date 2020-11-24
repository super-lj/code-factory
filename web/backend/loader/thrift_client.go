package loader

import (
	"web-backend/thrift/ci"

	"github.com/apache/thrift/lib/go/thrift"
)

var CIBackendAddr string = "localhost:9090"

func CreateCIBackendClient(addr string) (*ci.CIBackendClient, *thrift.TSocket, error) {
	transport, err := thrift.NewTSocket(addr)
	if err != nil {
		return nil, nil, err
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := ci.NewCIBackendClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		return nil, nil, err
	}

	return client, transport, nil
}
