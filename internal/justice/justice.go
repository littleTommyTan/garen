package justice

import (
	"github.com/littletommytan/garen/internal/justice/hellogrpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func SetupGrpcJustice() (j *grpc.Server) {
	j = grpc.NewServer()

	// grpc services assemble
	hellogrpc.Assemble(j)

	lis, err := net.Listen("tcp", ":2233")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		logrus.Infof("grpc server is running ...")
		if err := j.Serve(lis); err != nil {
			log.Print(err)
		}
	}()
	return
}
