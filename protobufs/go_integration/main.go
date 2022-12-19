package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	pb "go_integration/proto"
)

func getUser() *pb.User {
	return &pb.User{
		Id:   1,
		Name: "Sharat",
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		D: &pb.Dummy{
			Id:   1,
			Name: "Sharat",
		},
		DMul: []*pb.Dummy{
			{
				Id:   2,
				Name: "Sha",
			},
			{
				Id:   3,
				Name: "Sha",
			},
			{
				Id:   4,
				Name: "Sha",
			},
		},
	}
}

func doEnum() *pb.PersonalityTraits {
	return &pb.PersonalityTraits{
		EyeColor: pb.EyeColor_EYECOLOR_BLACK,
	}
}

func selOneOf(message interface{}) {

	switch x := message.(type) {
	case *pb.Choice_Id:
		fmt.Println(message.(*pb.Choice_Id).Id)
	case *pb.Choice_Msg:
		fmt.Println(message.(*pb.Choice_Msg).Msg)
	default:
		fmt.Errorf("hello %v", x)
	}
}

func doMaps() *pb.Maps {
	return &pb.Maps{
		Ids: map[string]*pb.IdWrapper{
			"id1": {Id: 1},
			"id2": &pb.IdWrapper{Id: 2},
			"id3": &pb.IdWrapper{Id: 3},
		},
	}
}

func fileMarshalUnmarshalExample(p proto.Message) {
	path := "simple.bin"
	writeToFile(path, p)
	pAnother := &pb.User{} // why &? https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
	readFromFile(path, pAnother)
	fmt.Println(pAnother)
}

func main() {
	fmt.Println(getUser())
	fmt.Println(doComplex())
	fmt.Println(doEnum())
	selOneOf(&pb.Choice_Msg{Msg: "as"})
	selOneOf(&pb.Choice_Id{Id: 1})
	fmt.Println(doMaps())
	fileMarshalUnmarshalExample(getUser())
	fmt.Println(toJSON(getUser()))
	fmt.Println(toJSON(doComplex()))
	fmt.Println(toJSON(doEnum()))
	fmt.Println(toJSON(doMaps()))

	a, _ := toJSON(getUser())
	b := pb.User{}
	fmt.Println(fromJSON(a, &b))
	fmt.Println(b.Id, b.Name, &b)
	fmt.Println(b.GetName())
}
