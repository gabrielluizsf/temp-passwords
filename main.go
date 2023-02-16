package main

import (
  "github.com/theGOURL/encrypting"
  "fmt"
)

func main(){
var  password string;
  fmt.Printf("You password: ");
  fmt.Scanf("%s",&password);
  generatePasswordEncrypted :=  encrypting.Encrypt(password,10);
  fmt.Println("Strong Password: ", generatePasswordEncrypted);
}
