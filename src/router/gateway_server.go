package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type GateWayRouter struct {
	*servers.GateWayCommServer
}

func (s *GateWayRouter) Boot() {
	s.GateWayCommServer = servers.NewGateWayCommServer()
}
func (s *GateWayRouter) Register() {

	s.Route("GET", "/local", wego.Handler("local"))

	//consul_demo 微服务上提供的 /demo/two 受限控制
	s.Route(
		"GET",
		"/consul_demo/demo/two",
		wego.Handler("limit_remote"),
	)

}
