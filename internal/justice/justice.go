package justice

import (
	"github.com/sirupsen/logrus"
	"github.com/tommytan/garen/internal/justice/hellogrpc"
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
		logrus.Infof("grpc server running...")
		log.Print("")
		if err := j.Serve(lis); err != nil {
			log.Print(err)
		}
	}()
	return
}
