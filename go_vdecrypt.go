package main

import(
  "os"
  "fmt"
  "github.com/sosedoff/ansible-vault-go"
)

func main() {

  if len(os.Args) < 2 { 
    fmt.Printf("Usage: %s Path Password\n", os.Args[0])
    os.Exit(1) 
  }

  file := os.Args[1]
  pass := os.Args[2]


  _, err := os.Stat(file)
  if err != nil {
     fmt.Printf("File %s does not exist. Exiting\n", file)
     os.Exit(1) 
   }

  str, err := vault.DecryptFile(file, pass)


  fmt.Println(str)

  if err != nil {
    fmt.Println(err)
  }


}
