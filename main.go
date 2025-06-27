package main

import "fmt"

func main(){

        Rohit := Details{
                Name: "Rohit Shankar",
                Email: "rohit123@gmail.com",
                Address: "Heaven",
                Age: 20,
        }
        fmt.Println("Rohit's Details")
        fmt.Println("Name:", Rohit.Name)
        fmt.Println("Email:", Rohit.Email)
        fmt.Println("Address:", Rohit.Address)
        fmt.Println("Age:", Rohit.Age)
        fmt.Printf("%+v",Rohit)

}
type Details struct{
        Name string
        Email string
        Address string
        Age int
}
