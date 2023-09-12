package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./jwt-token-generate <app_id> <path_to_private_key>")
		return
	}

	// GitHub App ID and Private Key
	appID := os.Args[1]
	privateKeyPath := os.Args[2]

	// Read the private key file
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		fmt.Printf("Error reading private key file: %v\n", err)
		return
	}

	// Parse the private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		fmt.Printf("Error parsing private key: %v\n", err)
		return
	}

	// Create the JWT token
	// Github's Maximum token limit is 10 minutes
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    appID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign the token with the private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Printf("Error signing token: %v\n", err)
		return
	}

	fmt.Println(tokenString)
}
