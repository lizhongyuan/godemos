package main

import (
    "./visibility"
    "fmt"
    "./sorts"
)

func main() {
    visibility.Public_fuction() //used in anywhere!
    //visibility.private_function() //不能访问私有函数，无法通过编译
    fmt.Println(visibility.P) //1
    //fmt.Println(visibility.p) //不能访问私有变量，无法通过编译
    fmt.Println(visibility.PI) //3.14
    //fmt.Println(visibility.pi) //不能访问私有常量，无法通过编译
    //fmt.Println(visibility._PI) //不能访问私有常量，无法通过编译

    var arr []int = []int{1,2,3,2,1}
    sorts.Pop_sort(&arr)
    fmt.Print(arr)
}
