package greetings

import "fmt"
import "errors"

func Hello(name string) (string,error) {
    if name == "" {
        return "",errors.New("name is []")
    }
    message := fmt.Sprintf("Hi,%v.welcome!", name)
    return message,nil
}


