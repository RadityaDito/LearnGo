package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomizer4 = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	randomizer := rand.New(rand.NewSource(10))
	fmt.Println("random ke-1:", randomizer.Int()) // 5221277731205826435
	fmt.Println("random ke-2:", randomizer.Int()) // 3852159813000522384
	fmt.Println("random ke-3:", randomizer.Int()) // 8532807521486154107

	randomizer2 := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random ke-1:", randomizer2.Int())
	fmt.Println("random ke-2:", randomizer2.Int())
	fmt.Println("random ke-3:", randomizer.Int())

	randomizer3 := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random int:", randomizer3.Int())
	fmt.Println("random float32:", randomizer3.Float32())
	fmt.Println("random uint:", randomizer3.Uint32())

}

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer4.Intn(len(letters))]
	}
	return string(b)
}
