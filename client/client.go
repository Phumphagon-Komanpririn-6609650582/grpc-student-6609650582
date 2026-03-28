package main

import (
	"context"
	"log"
	"time"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetStudent(ctx, &pb.StudentRequest{
		Id: 101,
	})

	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	log.Printf("Student Info:")
	log.Printf("ID: %d", res.Id)
	log.Printf("Name: %s", res.Name)
	log.Printf("Major: %s", res.Major)
	log.Printf("Email: %s", res.Email)
	log.Printf("Phone: %s", res.Phone)

	//เรียก ListStudents
	resList, err := client.ListStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling ListStudents: %v", err)
	}

	//แสดงนักศึกษาทั้งหทดพร้อม เบอร์โทร
	var i = 1
	log.Printf("\n=== List of All Students ===")
	for _, student := range resList.Student {
		log.Printf("Student Info: %d", i)
		log.Printf("ID: %d", student.Id)
		log.Printf("Name: %s", student.Name)
		log.Printf("Major: %s", student.Major)
		log.Printf("Email: %s", student.Email)
		log.Printf("Phone: %s", student.Phone)
		log.Printf("-------------------------")
		i++
	}
}
