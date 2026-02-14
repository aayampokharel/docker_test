package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("=============================")
	fmt.Println("hello this is from docker PC and whats up a to the b")
	fmt.Println("=============================")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	listenPort := viper.GetString("LISTEN_PORT")
	hostPort := viper.GetString("HOST_PORT")
	KEY := viper.GetString("KEY")
	SUBSTITUTE := viper.GetString("SUBSTITUTE")
	fmt.Printf("Server running on: %s\n", listenPort)
	fmt.Printf("DB Connection: %s\n", hostPort)
	fmt.Printf("KEY: %s\n", KEY)
	fmt.Printf("substitute example in .env file'SUBSTITUTE=localhost:${PORT}': %s\n", SUBSTITUTE)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world hello world heyyy whats up"))
	})
	er := http.ListenAndServe(":"+listenPort, nil)
	if er != nil {
		fmt.Println(er)
	}

}
