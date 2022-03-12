package main

import (
	"fmt"
	"os"

	"github.com/un4ckn0wl3z/what-is-my-isp/service"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "<phone number>")
		return
	}

	number := os.Args[1]
	fmt.Println("Target number:", number)
	service := service.GetIsp(number)
	fmt.Println("Target service:", service)

}
