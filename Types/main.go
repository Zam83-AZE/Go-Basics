package main

import "fmt"

func main(){
  /*
    INTEGER TYPES

    type   range    size(bytes)
    -----------------------------
    int8  -128 127 1b
    int16 -32768 32767 2b
    int32 -2147483648 2147483647 4b
    int64 -9223372036854755808 9223372036854775807 8b
    uint8 0 255 1b
    uint16 0 65535 4b
    uint32 0 4294967295 4b
    uint64 0 18446744073709551615 8b
    byte 0 255 1b
    rune -2147483648 2147483647 4b
    int 4b or 8b
    uint 4b or 8b

    FLOATING-POINT TYPES

    type   range    size(bytes)
    -----------------------------
    float32 1.4*10(-45) 3.4*10(38) 4b
    float64 4.9*10(-324) 1.8*10(308)

    COMPLEX NUMBERS TYPES

    type   range    size(bytes)
    -----------------------------
    complex64 
    complex128


    
  */

  var i =1
  const f="%T(%v)\n"
  fmt.Printf(f,i,i)

  /*
    TYPE CONVERSION
    The expression T(v) converts the value v to the type T.Some numeric conversions.

  */

  var i1=1
  var j1=float32(i1)
  println(j1)


  /*
    TYPE ASSERTION
     If you have value and you want to convert it to another or specific type(in case of interface{}) you can use type assertion. A type assertion takes a value and tries create another version  in the specified explicity type

     https://stackoverflow.com/questions/38816843/explain-type-assertions-in-go

     https://marcofranssen.nl/go-interfaces-and-type-assertions
  */

  type Person struct{
    name string
    age int
    address string
  }
    
    x:= Person{}

    println(x.name)


    var k int
    var h int
    _ =h
    fmt.Print("Enter key: ")
    fmt.Scanf("%d",&k)

    if d:=1;k==1{
    
      fmt.Println("ok")
    } else {
      fmt.Println(d)
      
    }
    fmt.Println(k)


    for i:=1;i<5;i++{
      println(i)
    }
    var g,y=1,2
    swap(&g,&y)
    fmt.Println(g,y)

    odd,even:=oddeven("13578")
    fmt.Println(odd,even)  


    fmt.Println(ssum(3,2,3))
   
}

func swap(a,b *int){
  *a,*b=*b,*a
}

func oddeven(value string) (int,int){
  odd,even:=0,0
  for _,e:=range value{
    if e%2==0{
      odd++
    } else {
      even++
    }
  }
  return odd,even
}


func ssum(nums...int) int {
  b:=0
  for _,a:=range nums{
    b=b+a
  }
  return b
}