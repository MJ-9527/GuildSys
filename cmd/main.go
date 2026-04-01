//package main
//
//import "github.com/MJ-9527/GulidSys/internal/api"
//
//func main() {
//	r := api.NewRouter()
//	err := r.Run(":8080")
//	if err != nil {
//		return
//	}
//}

package main

import (
	"fmt"

	"github.com/MJ-9527/GulidSys/internal/service"
)

func main() {

	u, err := service.RegisterUser("alice", "123456")
	fmt.Println(u, err)
}
