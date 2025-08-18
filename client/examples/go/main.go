package main

import (
	"fmt"

	"github.com/Rivalz-ai/proto/client/go/common"
)

func main() {
	//load env
	common.LoadENV()
	fmt.Println(common.GetEndpoint("micros.base.user"))
}
