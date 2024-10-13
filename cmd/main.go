package main

import (
	"aesnet256"
	"fmt"
	"log"
	"flag"
)

const version="0.1"

func main() {
	sizes := map[string]int {
		"128" : aesnet256.Aes_128_length,
		"192" : aesnet256.Aes_192_length,
		"256" : aesnet256.Aes_256_length,
	}
	ver := flag.Bool("v",false,"display that the version is " + fmt.Sprintf("%s",version) + " and exit.")
	key := flag.String("key","","decryption key, must be non-empty")
	text := flag.String("text","","test to encrypt or decrypt, must be non-empty")
	iv := flag.String("iv",aesnet256.Default_aesnet_iv,"Initialization Vector, usese aesencryption.net default if not specified")
	aes_type := flag.String("aes_type","256","AES type, default is 256, options are 128, 192 and 256.")
	mode := flag.String("mode","","mode type, use 'encrypt' or 'decrypt'")
	flag.Parse()
	if *ver {
		log.Printf("Version %s",version)
		return
	}
	if *key == "" || *text == "" {
		log.Printf("Key and text must have a value. They are blank. use -key and -text for a command line parameter.")
		return
	}
	if *mode != "encrypt"  && *mode != "decrypt" {
		log.Printf("Invalid mode, must be encrypt or decrypt")
		return
	}
	aes_size, ok := sizes[*aes_type]
	if !ok {
		log.Printf("AES size is invalid, must be 128, 192 or 256.")
		return
	}	
	log.Printf("Key is: %s",*key)
	log.Printf("Text is: %s", *text)
	log.Printf("IV is: %s", *iv)
	log.Printf("Aes_type is: %s", *aes_type)
	a := &aesnet256.Aes256{}
	a.Source = []byte(*text)
	a.IV = []byte(*iv)
	e := fmt.Errorf("Invalid mode, must be Encrypt or Decrypt")
	if *mode == "encrypt" {
		e = a.Encrypt(aesnet256.PadKey([]byte(*key), aes_size))
	}
	if *mode == "decrypt" {
		e = a.Decrypt(aesnet256.PadKey([]byte(*key), aes_size))
	}	
	if e != nil {
		log.Printf("Your result was malformed, perhaps the key is incorrect? try again.")
		return
	}
	if fmt.Sprintf("%s",a) == "" {
		log.Printf("The result was empty, your key was incorrect, try again!")
		return
	}
	log.Printf("Result is [%s]\n", a)
	return
}
