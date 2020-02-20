package main

import (
        "time"
	"fmt"
	"math/rand"
)

func main() {
        rand.Seed(time.Now().UTC().UnixNano())
	
	fmt.Println(fmt.Sprintf("%04d",rand.Intn(9999)))	
}
