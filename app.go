package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

const (
	key  = "6e45924d20fef273429090a206eb737c86ce1c7dccc8275cb313fc8a0d1ddcf38481535ad556c41a566a5c3eebda84c44271a3f2ee1bee48728689a307469c15"
	salt = "182ef2016fc483dcac88666ae9813469b94707f903d8f83efdc2f22ebbf9c5521c20330d9a6c326054c3524d5e08f8ba2e1a957dbdfc7777a4b1843e21fd2d14"
)

func main() {
	var keyBin, saltBin []byte
	var err error

	if keyBin, err = hex.DecodeString(key); err != nil {
		log.Fatal(err)
	}

	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal(err)
	}

	encodedSourceImage := "http://192.168.100.251:39000/digivetmr/0a9349e5-c1b1-4d7e-b964-fc9de0587ac9.JPG?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=EAI9CGRTZWFQCA20PXDZ%2F20221104%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20221104T085921Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiJFQUk5Q0dSVFpXRlFDQTIwUFhEWiIsImV4cCI6MTY2NzU1NTMzOCwicGFyZW50IjoidGFiZWxkYXRhIn0.PLyOaJ0nO6zjS6sFpFU12vwrPxgvnyzuVj1LWNHbikvn1UhkKNUzpfjuwnCkszONtuUx1wbg6V09NbOjS6as6g&X-Amz-SignedHeaders=host&versionId=null&X-Amz-Signature=6e91a0250de99b515ecd1f5aada5df4a3cf3173c8507697954951029b43d90ee"
	path := "/rs:fit:300:300/plain/" + encodedSourceImage

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	fmt.Printf("/%s%s", signature, path)
}
