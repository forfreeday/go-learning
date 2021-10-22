package P03struct

import "fmt"

type Books struct {
	Title   string
	Author  string
	Subject string
	BookId  int
}

//方式一
//var book1 Books
func StructTest1() {
	var book1 Books
	var book2 Books

	book1.Author = "嗯哼"
	book1.Title = "test"
	book1.Subject = "subjecttest1"
	book1.BookId = 1234

	book2.Author = "嗯哼2"
	book2.Title = "test2"
	book2.Subject = "subject2222"
	book2.BookId = 123123123

	printBook(&book1)
	printBook(&book2)
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.Title)
	fmt.Printf("Book author : %s\n", book.Author)
	fmt.Printf("Book subject : %s\n", book.Subject)
	fmt.Printf("Book book_id : %d\n", book.BookId)
}

type Person struct {
	Name string
	Age  int
}

//方式二 new
func StructTest2() {
	var person *Person = new(Person)
	person.Name = "结构体测试"
	(*person).Age = 111
	fmt.Println(*person)
}

//方式三 &
func StructTest3() {
	var person *Person = &Person{
		Name: "测试&方式",
	}

	(*person).Age = 1234
	fmt.Println(*person)
}

//方式三 &
func StructTest31() {
	var person1 *Person = &Person{
		Name: "测试&方式111",
	}
	(*person1).Age = 1234

	var person2 *Person = &Person{
		Name: "测试&方式222",
	}
	person2.Age = 444444

	fmt.Printf("Persion1 Address: %p, 内容: %v \n", &person1, person1)
	fmt.Printf("Persion2 Address: %p, 内容: %v \n", &person2, person2)
}
