package main

import (
	"fmt"

	"github.com/skast96/pushshift/pushshift"
)

func main() {
	client := pushshift.NewClient("testClient/0.1.0")
	comments := client.StreamComments(nil)
	for comment := range comments {
		fmt.Println(comment.Subreddit)
		fmt.Println(comment.Author)
		fmt.Println(comment.Body)
	}
}
