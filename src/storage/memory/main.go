package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Auohtor string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Auohtor] = append(PostsByAuthor[post.Auohtor], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{1, "hello go", "Lee"}
	post2 := Post{2, "Bonjour go", "Fire"}
	post3 := Post{3, "Hola go", "CoCo"}
	post4 := Post{4, "Greetings go", "Lee"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Lee"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["CoCo"] {
		fmt.Println(post)
	}
}
