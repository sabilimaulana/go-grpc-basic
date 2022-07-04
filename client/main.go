package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sabilimaulana/go-grpc-basic/common/config"
	"github.com/sabilimaulana/go-grpc-basic/common/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "u001",
		Name:     "Suzuka",
		Password: "baby",
		Gender:   model.UserGender(model.UserGender_value["FEMALE"]),
	}

	user2 := model.User{
		Id:       "u002",
		Name:     "Moa",
		Password: "metal",
		Gender:   model.UserGender(model.UserGender_value["FEMALE"]),
	}

	garage1 := model.Garage{
		Id:   "g001",
		Name: "Going Merry",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	garage2 := model.Garage{
		Id:   "g002",
		Name: "Thousand Sunny",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	garage3 := model.Garage{
		Id:   "g003",
		Name: "The Oro Jackson",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	user := serviceUser()

	fmt.Println("\n", "===========> user test")

	// register user1
	user.Register(context.Background(), &user1)

	// register user2
	user.Register(context.Background(), &user2)

	// show all registered users
	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	garage := serviceGarage()

	fmt.Println("\n", "===========> garage test A")

	// add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	// add garage2 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage2,
	})

	// show all garages of user1
	res2, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))

	fmt.Println("\n", "===========> garage test B")

	// add garage3 to user2
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user2.Id,
		Garage: &garage3,
	})

	// show all garages of user2
	res3, err := garage.List(context.Background(), &model.GarageUserId{UserId: user2.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res3String, _ := json.Marshal(res3.List)
	log.Println(string(res3String))
}
