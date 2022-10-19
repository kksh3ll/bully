package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/kksh3ll/bully/bully"
)

var argvPort = flag.Int("port", 8117, "port to listen")
var argvCandidates = flag.String("nodes", "", "comma separated list of nodes.")
var argId = flag.Int("id", 999, "id for node.")

func main() {
	flag.Parse()
	bindAddr := fmt.Sprintf("0.0.0.0:%v", *argvPort)

	ln, err := net.Listen("tcp", bindAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	bly := bully.NewBully(ln, *argId)
	defer bly.Finalize()

	nodeAddr := strings.Split(*argvCandidates, ",")
	dialTimeout := 5 * time.Second

	for _, node := range nodeAddr {
		if len(node) == 0 {
			continue
		}
		err := bly.AddCandidate(node, nil, dialTimeout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v cannot be added: %v\n", node, err)
		}
	}

	fmt.Printf("My ID: %v\n", bly.MyId())
	for {
		time.Sleep(60 * time.Second)
		bly.PrintCandidateList()
		l, _, _ := bly.Leader()
		fmt.Println("leader is ", l.Addr)
	}
}
