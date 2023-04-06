package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	//"io"
	"jjchavarrg.com/go/grpc/testpb"
	"log"
	"time"
)

func main() {
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)
	log.Printf("--- DoUnary(c) ------------")
	DoUnary(c)
	log.Printf("--- DoClienStreaming(c) ---")
	DoClienStreaming(c)
	//DoServerStreaming(c)
	//DoBidirectionalStreaming(c)
}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "t1",
	}

	res, err := c.GetTest(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GetTest: %v\n", err)
	}
	log.Printf("Response from GetTest: %v\n", res)
}

func DoClienStreaming(c testpb.TestServiceClient) {
	questions := []*testpb.Question{
		{
			Id:       "q1t1",
			Answer:   "Azul",
			Question: "Color asociado a Golang",
			TestId:   "t1",
		}, {
			Id:       "q2t1",
			Answer:   "Google",
			Question: "Empresa que desarrolla el lenguaje  Golang",
			TestId:   "t1",
		}, {
			Id:       "q3t1",
			Answer:   "Backend",
			Question: "Especialidad de Golang",
			TestId:   "t1",
		},
	}

	stream, err := c.SetQuestions(context.Background())
	if err != nil {
		log.Fatalf("Error while calling SetQuestions: %v\n", err)
	}
	for _, question := range questions {
		log.Println("Sending question: ", question.Id)
		stream.Send(question)
		time.Sleep(2 * time.Second)
	}

	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}
	log.Printf("Response from SetQuestions: %v\n", msg)
}
