package handler

import (
	"ci-backend/thrift/ci"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
)

func InitThriftServer(addr string) *thrift.TSimpleServer {
	pr := ci.NewCIBackendProcessor(&CIBackendServiceHandler{})
	tp, err := thrift.NewTServerSocket(addr)
	if err != nil {
		log.Fatalf("Cannot start thrift server: %v", err)
	}
	return thrift.NewTSimpleServer2(pr, tp)
}
