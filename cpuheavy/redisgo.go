package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"strconv"
)

type ServeHandler struct {
	client *redis.Client
}

func (sh *ServeHandler) serve(w http.ResponseWriter, r *http.Request) {
	tmp := fibonaccired(90, sh.client)
	fmt.Fprintln(w, strconv.Itoa(tmp))
}

func fibonaccired(num int, client *redis.Client) int {
	if num <= 1 {
		return num
	}
	val, err := client.Get(strconv.Itoa(num)).Result()
	if err == redis.Nil {
		res := fibonaccired(num-1, client) + fibonaccired(num-2, client)
		client.Set(strconv.Itoa(num), res, 0)
		return res
	} else if err != nil {
		panic(err)
	} else {
		ret, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		return ret
	}

}


func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password:"",
		DB: 0,
	})
	myHandler := &ServeHandler{client:client}

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	http.HandleFunc("/", myHandler.serve)
	http.ListenAndServe(":1235", nil)
}
