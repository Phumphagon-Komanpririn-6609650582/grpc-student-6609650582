package main

import (
	"context"
	"log"
	"net"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

func (s *server) GetStudent(ctx context.Context, req *pb.StudentRequest) (*pb.StudentResponse, error) {
	log.Printf("Received request for student ID: %d", req.Id)

	return &pb.StudentResponse{
		Id:    req.Id,
		Name:  "Alice Johnson",
		Major: "Computer Science",
		Email: "alice.joh@dome.tu.ac.th",
		Phone: "081-111-1111", //เพิ่ม phone
	}, nil
}

// เมธอด ListStudents จาก array
func (s *server) ListStudents(ctx context.Context, req *pb.Empty) (*pb.StudentListResponse, error) {
	log.Println("Received request for ListStudents")

	// List นักเรียน
	students := []*pb.StudentResponse{
		{Id: 101, Name: "Alice Johnson", Major: "Computer Science", Email: "alice.joh@dome.tu.ac.th", Phone: "081-111-1111"},
		{Id: 102, Name: "Phumphagon Komanpririn", Major: "Computer Science", Email: "phumphagon.kom@dome.tu.ac.th", Phone: "099-999-9999"},
		{Id: 103, Name: "Alex Manuel", Major: "Computer Science", Email: "alex.man@dome.tu.ac.th", Phone: "066-666-666"},
	}

	return &pb.StudentListResponse{
		Student: students,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStudentServiceServer(grpcServer, &server{})

	log.Println("gRPC Server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
