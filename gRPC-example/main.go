package main



import (
	"context"
	"flag"
	"log"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// pb "./firsttest/helloworld"
	pb "helloworlddd"
)


const (
	defaultName = "safekey"
)


var (
	addr = flag.String("addr", "172.17.17.113:50051", "address to connect")
	name = flag.String("name", defaultName, "name to connect to")
)



func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSafeKeyProtoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GenerateNewSafeKey(ctx, &pb.SafeKeyNameProtoMessage{Name: *name})
	if err != nil {
		log.Fatalf("could not generate new safe key: %v", err)
	}
	log.Printf("generated new safe key: %s", r.GetHash())
}




