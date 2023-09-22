package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"encoding/json"
)

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.200.144:6379",
		Password: "",
		DB: 0,
	})
	
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	
	//Set a value in redis
	err = client.Set("name", "Elliot", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	
	//Get a value
	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	
	//Overwrite the value
	err = client.Set("name", 123, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	
	//Get a value
	val, err = client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	
	//Store complex data
	json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
    	if err != nil {
        	fmt.Println(err)
    	}
    	
    	err = client.Set("id1234", json, 0).Err()
    	if err != nil {
        	fmt.Println(err)
    	}
    	val, err = client.Get("id1234").Result()
    	if err != nil {
        	fmt.Println(err)
   	}
    	fmt.Println(val)
    	
    	
}
