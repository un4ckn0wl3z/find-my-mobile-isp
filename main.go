package main

import (
	"fmt"
	"os"

	"github.com/un4ckn0wl3z/find-my-mobile-isp/service"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "<phone number>")
		return
	}

	number := os.Args[1]
	fmt.Println("Target number:", number)
	service, c := service.GetISP(number)
	c.Println("Target service:", service)
	c.DisableColor()

}
