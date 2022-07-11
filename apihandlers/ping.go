package apihandlers

import (
	"fmt"
	"net/http"
	"setu/handler"
)

// PingHandler is a Basic ping utility for the service
type PingHandler struct {
	BaseHandler
}

func (p *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := handler.RouteApiCall(p, r, []string{http.MethodGet})
	response.RenderResponse(w)
}

// Get method for PingHandler
func (p *PingHandler) Get(r *http.Request) handler.ServiceResponse {
	fmt.Println("Got Ping Request")
	return handler.Response200OK("PONG")
}
