package app

import (
	"github.com/chengleqi/kubesphere-core/pkg/models/controllers"
	"k8s.io/klog/v2"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/chengleqi/kubesphere-core/pkg/apis/v1alpha"
)

type kubeSphereServer struct {
	address net.IP
	port    int
}

func (s *kubeSphereServer) run() {
	stopChan := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	// start controllers
	go controllers.Run(stopChan, &wg)

	// start gin server
	go func() {
		r := v1alpha.SetupRouter()
		err := r.Run(s.address.String() + ":" + strconv.Itoa(s.port))
		if err != nil {
			klog.Fatal("gin server start failed")
		}
	}()

	// end
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	close(stopChan)
	wg.Wait()
}

func newKubeSphereServer() *kubeSphereServer {
	return &kubeSphereServer{
		address: net.ParseIP("127.0.0.1"),
		port:    9999,
	}
}

func Run() {
	server := newKubeSphereServer()
	server.run()
}
