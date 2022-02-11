package grpc

import (
	"context"
	"log"

	env "github.com/aleesilva/hash-challenger-back-end/config"
	"github.com/aleesilva/hash-challenger-back-end/pb"

	"google.golang.org/grpc"
)

func GetDiscount(productId int) float32 {

	conn, err := grpc.Dial("localhost:"+env.LoadDotEnvVariable("GRPC_SERVICE_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewDiscountClient(conn)

	req := &pb.GetDiscountRequest{
		ProductID: int32(productId),
	}

	res, err := client.GetDiscount(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	return res.GetPercentage()
}
