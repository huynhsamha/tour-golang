package main

import "fmt"

// Post : post of an user
type Post struct {
	ID      int64
	Title   string
	Content string
	Owner   *User
}

// User : user profile
type User struct {
	ID       int64
	Username string
	Password string
	Email    string
	Posts    []*Post
}

// Create method for an instance User
// we can call user.printFormat()
func (user User) printFormat() {
	fmt.Println("****** User Profile ******")
	fmt.Println("ID       : ", user.ID)
	fmt.Println("Username : ", user.Username)
	fmt.Println("Email    : ", user.Email)
	fmt.Println("Posts    : ")
	for _, val := range user.Posts {
		fmt.Println(val)
	}
	fmt.Println("**************************")
}

// Create method for an instance Post
// we can call post.printFormat()
func (post Post) printFormat() {
	fmt.Println("****** Post Info ******")
	fmt.Println("ID       : ", post.ID)
	fmt.Println("Title    : ", post.Title)
	fmt.Println("Content  : ", post.Content)
	if post.Owner != nil {
		fmt.Println("OwnerID  : ", post.Owner.ID)
	}
	fmt.Println("**************************")
}

func (user *User) addPost(post *Post) {
	post.Owner = user
	user.Posts = append(user.Posts, post)
}

func main() {
	posts := []Post{
		Post{101, "Post 1", "Lorem Isum 1", nil},
		Post{102, "Post 2", "Lorem Isum 2", nil},
		Post{103, "Post 3", "Lorem Isum 3", nil},
	}

	user := User{508, "Alice", "qwiyv9w43", "alice@bob.trudy", nil}

	user.addPost(&posts[0])
	user.addPost(&posts[2])
	user.printFormat()

	posts[0].printFormat()
	posts[1].printFormat()
	posts[2].printFormat()
}
