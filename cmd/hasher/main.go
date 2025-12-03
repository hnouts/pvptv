package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/argon2"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password to hash: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if password == "" {
		fmt.Println("Password cannot be empty")
		return
	}

	// Parameters MUST match internal/http/admin.go
	salt := []byte("pvptv-admin")
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	fmt.Printf("\nAdd this to your .env file:\n")
	fmt.Printf("ADMIN_USERNAME=admin\n")
	fmt.Printf("ADMIN_PASSWORD_HASH=%x\n", hash)
}

