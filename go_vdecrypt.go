package main

import(
  "os"
  "fmt"
  "github.com/sosedoff/ansible-vault-go"
)

func main() {

  file := os.Args[1]
  pass := os.Args[2]

  // Encrypt secret data
  //str, err := vault.Encrypt(data, "password")

  str, err := vault.DecryptFile(file, pass)


  fmt.Println(str)

  if err != nil {
    fmt.Println(err)
  }


  // Decrypt secret data
  //str, err := vault.Decrypt("secret", "password")

  // Write secret data to file
  //err := vault.EncryptFile("path/to/secret/file", "secret", "password")

  // Read existing secret

}
