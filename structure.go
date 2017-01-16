package main

import "fmt"

type Book struct{
	title string
	author string
	subject string
	book_id int
}

func main(){
	var book1 Book
	var book2 Book
	
	book1.title = "Harry Porter and the Philosophers Stone"
	book1.author = "J.K.Rowling"
	book1.subject = "Fantacy"
	book1.book_id = 1000
	
	book2.title = "Harry Potter and Prisoner of Azkaban"
	book2.author = "J.K.Rowling"
	book2.subject = "Fantacy"
	book2.book_id = 1001

	printBook(book1)
	fmt.Println()
	printBook(book2)
}

func printBook(book Book){
	fmt.Println(book.title)
	fmt.Println(book.author)
	fmt.Println(book.subject)
	fmt.Println(book.book_id)
}