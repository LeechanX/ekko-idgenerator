package main

import (
	"context"
	"fmt"

	ekko_client "github.com/leechanx/ekko-idgenerator/client"
)

func main() {
	client, err := ekko_client.NewEkkoClient([]string{"127.0.0.1:8972"})
	//client, err := ekko_client.NewEkkoClient([]string{"127.0.0.1:8972"}, ekko_client.WithFallback())
	if err != nil {
		panic(err)
	}
	defer client.Close()

	for i := 0; i < 10; i++ {
		uid, err := client.IDGen(context.TODO(), 0)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(uid)
		}
	}

}
