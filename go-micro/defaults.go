package micro

import (
	"github.com/mjlee1983/go-micro/go-micro/v2/client"
	"github.com/mjlee1983/go-micro/go-micro/v2/debug/trace"
	"github.com/mjlee1983/go-micro/go-micro/v2/server"
	"github.com/mjlee1983/go-micro/go-micro/v2/store"

	// set defaults
	gcli "github.com/mjlee1983/go-micro/go-micro/v2/client/grpc"
	memTrace "github.com/mjlee1983/go-micro/go-micro/v2/debug/trace/memory"
	gsrv "github.com/mjlee1983/go-micro/go-micro/v2/server/grpc"
	memoryStore "github.com/mjlee1983/go-micro/go-micro/v2/store/memory"
)

func init() {
	// default client
	client.DefaultClient = gcli.NewClient()
	// default server
	server.DefaultServer = gsrv.NewServer()
	// default store
	store.DefaultStore = memoryStore.NewStore()
	// set default trace
	trace.DefaultTracer = memTrace.NewTracer()
}
