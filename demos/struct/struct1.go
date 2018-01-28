package main


import "fmt"

type User struct {
	 name	string
	 email	string
	 password	string
}

func main() {
	// 声明
	var user User
	var userref *User
	fmt.Println(user)
	fmt.Println(userref)

	userref = new(User)
	fmt.Println(userref)

	user1 := User {
		name:"lzy",
		email:"fenixlee@163.com",
		password:"1234"}

	userref1 := &User{name:"lzy", email:"fenixlee@163.com", password:"1234"}
	fmt.Println(user1)
	fmt.Println(userref1)

	type LinkedNode struct {
		data string
		su *LinkedNode
	}

	type DoubleLinkedNode struct {
		pr *DoubleLinkedNode
		data string
		su *DoubleLinkedNode
	}

	type TreeNode struct {
		left *TreeNode
		data string
		right *TreeNode
	}
}
