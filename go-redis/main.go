package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

type User struct {
	Id   int64
	Name string `json:"name"`
}

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic("Error Connecting with Redis")
	}
	fmt.Println("pong: ", pong)
	user := User{
		Id:   5,
		Name: "akipiD",
	}

	obj, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	keyStr := strconv.Itoa(int(user.Id))
	err = client.Set(keyStr, obj, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get(keyStr).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("val :", val)

	var newUser User
	err = json.Unmarshal([]byte(val), &newUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, _ := client.Del("1").Result()
	fmt.Println("resp :", resp)
}
