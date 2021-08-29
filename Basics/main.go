package main

import "fmt"

func main(){
  /*
     Variables

     The var statement declares the list of variables

     var(
        name string
        age   int
        location string
     )
     or even
     var(
       name,location string
       age int
     )

     Variable can also declred one by one
     var name string
     var location string
     var age int

     A varible declaration can include initializers one per variable

     var name string="zamir"
     var age int=38
     var location string="baku"

     If an initializer is present the type can be ommited the variable will take the type of initializer(inferred typing)

     var(
        name="zamir"
        age=38
        location="baku" 
     )

     You can also initialize variables in the same line
     var(
       name,age,location="zamir",38,"baku"
     )

     Inside a function the := short assigment statement can be used in place of var declaration with implicit types

     A variable can contain any type including functions:

     action:=func(){
       //do something
     }
     action()

     Outside a function every construct begin with var,func and := construct is not available.





  */
    name,location:="zamir","baku"
    age:=30

    fmt.Printf("%s %s %d \n",name,location,age)

    sayHello:=func(word string){
      fmt.Println(word)
    }
    sayHello("Salam")

}