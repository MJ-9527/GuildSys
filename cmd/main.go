package main

import (
	"fmt"

	"github.com/MJ-9527/GulidSys/internal/repo"
)

func main() {
	//r := api.NewRouter()
	//err := r.Run(":8080")
	//if err != nil {
	//	return
	//}
	g, err := repo.CreateGuild("Knights", 4562)
	fmt.Println(g, err)

	g2, err2 := repo.GetGuildByID(1)
	fmt.Println(g2, err2)

	m, err := repo.AddMember(1, 2, "member")
	fmt.Println(m, err)

	members, _ := repo.GetMembersByGuild(1)
	fmt.Println(members)
}
