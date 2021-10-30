package main

import (
	"bytes"
	"fmt"
	"grpc-learn/model"
	"strings"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	user1 := &model.User{
		Id:       "u0001",
		Name:     "Kharisma Muzaki",
		Password: "f0r th3 p4ssword",
		Gender:   model.UserGender_FEMALE,
	}

	_ = &model.UserList{
		List: []*model.User{
			user1,
		},
	}

	garage1 := &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	garageList := &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	_ = &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	fmt.Printf("#=== Original\n		%#v", user1)
	fmt.Printf("\n\n#=== As string\n 	%v \n", user1.String())

	// ====== as json string
	var buf bytes.Buffer

	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("\n\n#=== As JSON string\n %v \n", buf.String())

	// ======= convert from json string
	jsonString := `{"list":[{"id":"g001","name":"Kalimdor","coordinate":{"latitude":23.221285,"longitude":53.220333}}]}`
	buf2 := strings.NewReader(jsonString)
	protoObject := new(model.GarageList)

	err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)

	if err2 != nil {
		panic(err2)
	}

	fmt.Printf("\n\n#=== As string from JSON -> Proto\n %v \n", protoObject.String())
}
