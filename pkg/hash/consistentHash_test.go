package hash

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func Test(key *testing.T) {

	NewHash := NewHashConsistent()

	for i := 0; i < 3; i++ {
		port, _ := strconv.Atoi(fmt.Sprintf("808%d", i))
		NewHash.Add(NewNode(i, 2, port, fmt.Sprintf("127.0.0.%d", i)))
	}

	//玩家id
	log.Println(NewHash.nodes)
	userId3 := "322232324"
	user3HashId := NewHash.CheckHashKey(userId3)

	log.Fatalln("userId:", userId3, "|hashId:", user3HashId)
}
