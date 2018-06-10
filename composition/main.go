package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

type post struct {
	title   string
	content string
	author
}

type website struct {
	posts []post
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.title)
	fmt.Println("Author: ", p.author.fullName())
	fmt.Println("Bio: ", p.author.bio)
}

func (w website) contents() {
	fmt.Println("Contents of website")
	fmt.Println("--------------------")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Sri Harsha",
		"Kappala",
		"Avid Programmer & Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post2 := post{
		"Struct instead of class",
		"Go doesn't support classes, instead use structs",
		author1,
	}
	w := website{
		posts: []post{post1, post2},
	}
	w.contents()
}
