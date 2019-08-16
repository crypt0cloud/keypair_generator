package main

import (
	"flag"
	"fmt"
	"github.com/crypt0cloud/core/crypto"
	"golang.org/x/crypto/ed25519"
	"log"
	"math/rand"
	"time"
)

var algorithm string

func main() {
	flag.StringVar(&algorithm, "algorithm", "ed25519", "the algorithm used to generate the key")
	flag.Parse()

	fmt.Printf("\nAlgorithm: %s\n\n",algorithm)
	if algorithm == "ed25519" {
		pub, sec := get_key_pair()
		pkey := crypto.Base64_encode(pub)
		fmt.Printf("Pubic Key:\n%s\n\n",pkey)

		skey := crypto.Base64_encode(sec)
		fmt.Printf("Private Key:\n%s\n\n",skey)
	}
}


func get_key_pair()(publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey){
	// Generate Key Pair from random data
	publicKey, privateKey, err := ed25519.GenerateKey(rand.New(rand.NewSource(time.Now().UnixNano())))
	if err != nil{
		log.Panic(err)
	}
	return publicKey, privateKey
}