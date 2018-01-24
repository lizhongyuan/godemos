package main

import "fmt"

type book struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var book1 book
	var book2 book

	/* book 1 描述 */
	book1.title = "Go 语言"
	book1.author = "www.runoob.com"
	book1.subject = "Go 语言教程"
	book1.book_id = 6495407

	/* book 2 描述 */
	book2.title = "Python 教程"
	book2.author = "www.runoob.com"
	book2.subject = "Python 语言教程"
	book2.book_id = 6495700

	/* 打印 book1 信息 */
	fmt.Printf( "book 1 title : %s\n", book1.title)
	fmt.Printf( "book 1 author : %s\n", book1.author)
	fmt.Printf( "book 1 subject : %s\n", book1.subject)
	fmt.Printf( "book 1 book_id : %d\n", book1.book_id)

	/* 打印 book2 信息 */
	fmt.Printf( "book 2 title : %s\n", book2.title)
	fmt.Printf( "book 2 author : %s\n", book2.author)
	fmt.Printf( "book 2 subject : %s\n", book2.subject)
	fmt.Printf( "book 2 book_id : %d\n", book2.book_id)
}
