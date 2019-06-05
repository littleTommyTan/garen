package greeterendpoint

import (
	"context"
	"github.com/littletommytan/garen/internal/gokit/greeterservice"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a greeter service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	HealthEndpoint   endpoint.Endpoint // used by Consul for the healthcheck
	GreetingEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns service Endoints, and wires in all the provided
// middlewares.
func MakeServerEndpoints(s greeterservice.Service) Endpoints {
	var healthEndpoint endpoint.Endpoint
	{
		healthEndpoint = MakeHealthEndpoint(s)
	}

	var greetingEndpoint endpoint.Endpoint
	{
		greetingEndpoint = MakeGreetingEndpoint(s)
	}

	return Endpoints{
		HealthEndpoint:   healthEndpoint,
		GreetingEndpoint: greetingEndpoint,
	}
}

// MakeHealthEndpoint constructs a Health greeterendpoint wrapping the service.
func MakeHealthEndpoint(s greeterservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		healthy := s.Health()
		return HealthResponse{Healthy: healthy}, nil
	}
}

// MakeGreetingEndpoint constructs a Greeter greeterendpoint wrapping the service.
func MakeGreetingEndpoint(s greeterservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GreetingRequest)
		greeting := s.Greeting(req.Name)
		return GreetingResponse{Greeting: greeting}, nil
	}
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so if they've
// failed, and if so encode them using a separate write path based on the error.
type Failer interface {
	Failed() error
}

// HealthRequest collects the request parameters for the Health method.
type HealthRequest struct{}

// HealthResponse collects the response values for the Health method.
type HealthResponse struct {
	Healthy bool  `json:"healthy,omitempty"`
	Err     error `json:"err,omitempty"`
}

// Failed implements Failer.
func (r HealthResponse) Failed() error { return r.Err }

// GreetingRequest collects the request parameters for the Greeting method.
type GreetingRequest struct {
	Name string `json:"name,omitempty"`
}

// GreetingResponse collects the response values for the Greeting method.
type GreetingResponse struct {
	Greeting string `json:"greeting,omitempty"`
	Err      error  `json:"err,omitempty"`
}

// Failed implements Failer.
func (r GreetingResponse) Failed() error { return r.Err }
