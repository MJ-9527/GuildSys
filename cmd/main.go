package main

import "github.com/MJ-9527/GulidSys/internal/service"

func main() {
	//r := api.NewRouter()
	//err := r.Run(":8080")
	//if err != nil {
	//	return
	//}

	_, err := service.RegisterUser("123", "9735")
	if err != nil {
		return
	}

	_, err = service.Login("123", "9735")
	if err != nil {
		return
	}
}
