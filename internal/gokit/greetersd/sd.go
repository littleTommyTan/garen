package greetersd

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
)

// ConsulRegister method.
func ConsulRegister() (registar *consulsd.Registrar) {

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	// Service discovery domain. In this example we use Consul.
	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		consulConfig.Address = "localhost:8500"
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			_ = logger.Log("err", err)
			os.Exit(1)
		}
		client = consulsd.NewClient(consulClient)
	}

	check := api.AgentServiceCheck{
		HTTP:                           "http://localhost:8500",
		Interval:                       "3s",
		Timeout:                        "3s",
		Notes:                          "Basic health checks",
		Method:                         "Get",
		DeregisterCriticalServiceAfter: "30s",
	}

	num := rand.Intn(1000000)
	asr := api.AgentServiceRegistration{
		ID:      "go-kit-srv-greeter-" + strconv.Itoa(num),
		Name:    "go-kit-srv-greeter",
		Address: "192.168.1.4",
		Port:    5500,
		Tags:    []string{"go-kit", "greeter"},
		Check:   &check,
	}
	registar = consulsd.NewRegistrar(client, &asr, logger)
	return
}
