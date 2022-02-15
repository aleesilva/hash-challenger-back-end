package integrations

import (
	"context"
	"log"
	"os"

	"github.com/aleesilva/hash-challenger-back-end/pb"

	"google.golang.org/grpc"
)

func GetDiscount(productId int) (float32, error) {

	conn, err := grpc.Dial("grpc-service:"+os.Getenv("GRPC_SERVICE_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewDiscountClient(conn)

	req := &pb.GetDiscountRequest{
		ProductID: int32(productId),
	}

	res, err := client.GetDiscount(context.Background(), req)
	if err != nil {
		return float32(0), err
	}

	return res.GetPercentage(), nil
}
