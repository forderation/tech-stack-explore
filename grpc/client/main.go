package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/forderation/tech-stack-explore/grpc/common/config"
	"github.com/forderation/tech-stack-explore/grpc/common/model"
	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "n001",
		Name:     "zaaki zaki",
		Password: "hooman try to b3st",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	user2 := model.User{
		Id:       "n001",
		Name:     "trifebsi",
		Password: "let the w0rld run",
		Gender:   model.UserGender(model.UserGender_value["FEMALE"]),
	}

	userService := serviceUser()
	fmt.Println("\n", "==========> user test")

	// Register user 1
	userService.Register(context.Background(), &user1)
	userService.Register(context.Background(), &user2)

	garageService := serviceGarage()
	fmt.Println("\n", "======> Garage Test")
	garage1 := model.Garage{
		Id:   "q001",
		Name: "Quel'thalas",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}
	garage2 := model.Garage{
		Id:   "q002",
		Name: "Jakarta",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	// Add both garage 1 to user1
	garageService.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	// Show all garages user 1
	garagesUser1, err := garageService.List(context.Background(), &model.GarageUserId{
		UserId: user1.Id,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	garagesUser1Bytes, _ := json.Marshal(garagesUser1.List)
	log.Println("user1 garages: ", string(garagesUser1Bytes))

	// Add garage 2 only to user2
	garageService.Add(context.Background(), &model.GarageAndUserId{
		UserId: user2.Id,
		Garage: &garage2,
	})

	// Show all garages user 2
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	garagesUser2, err := garageService.List(ctx, &model.GarageUserId{
		UserId: user2.Id,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	garageUser2Bytes, _ := json.Marshal(garagesUser2.List)
	log.Println("user2 garages:", string(garageUser2Bytes))
}
