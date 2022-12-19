package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"os"
)

// proto.Message is the top level interface that all other proto messages implement
func writeToFile(filename string, pb proto.Message) {
	out, err := proto.Marshal(pb)

	if err != nil {
		fmt.Println("Marshaling error", err)
		return
	}

	err = os.WriteFile(filename, out, 0644)
	if err != nil {
		fmt.Println("Write to file error", err)
		return
	}
	fmt.Println("Done writing to", filename)

}

func readFromFile(filename string, pb proto.Message) {
	out, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("can't read file", err)
		return
	}

	if err = proto.Unmarshal(out, pb); err != nil {
		fmt.Println("Unmarshal error", err)
	}
	fmt.Println("Done reading")
}



