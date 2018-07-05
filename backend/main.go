package main

import (
	"flag"
	"fmt"

	"github.com/ComputePractice2018/microblog/backend/utils"
)

func main() {
	name := flag.String("name", "Александр", "имя для преветствия")
	flag.Parse()
	fmt.Println(utils.GetHelloWorldString(*name))
}
