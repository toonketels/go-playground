package main

import . "fmt"

const Hello = "Hello"
var world string

func init() {
    world = "world"
}

func main() {
    Println(Hello, world)
}
