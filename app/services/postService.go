package postService

import (
	"time"

	"github.com/lanjiang/reddit-lite/app/models"
)

var (
	posts  []*models.Post
	cursor int
)

//Add a post to list
func Add(post *models.Post) int {
	post.CreatedOn = time.Now()
	cursor = cursor + 1
	post.PostId = cursor
	posts = append(posts, post)
	sort(posts, len(posts)-1)
	return post.PostId
}

//GetAll return a list of posts ordered by upvote and datetime
func GetAll(count int) []*models.Post {
	if count > len(posts) {
		count = len(posts)
	}
	return posts[0:count]
}

//GetByID returns a post by Id
func GetByID(id int) (int, *models.Post) {
	for i, post := range posts {
		if post.PostId == id {
			return i, post
		}
	}
	return -1, nil
}

//Delete a post
func Delete(post *models.Post) bool {
	if post == nil {
		return false
	}

	for i, p := range posts {
		if p.PostId == post.PostId {
			posts = append(posts[:i], posts[i+1:]...)
			return true
		}
	}
	return false
}

//Vote a post(upvote or downvote)
func Vote(id int, voteType int) {
	index, post := GetByID(id)
	switch voteType {
	case 0:
		post.DownVote = post.DownVote + 1
	case 1:
		post.UpVote = post.UpVote + 1
	}
	sort(posts, index)
}

//sort post by upvote and datetime using insertion sort
func sort(p []*models.Post, index int) {
	for j := index; j > 0; j-- {
		if p[j].UpVote > p[j-1].UpVote {
			p[j-1], p[j] = p[j], p[j-1]
		} else if p[j].UpVote == p[j-1].UpVote && p[j-1].CreatedOn.Before(p[j].CreatedOn) {
			p[j-1], p[j] = p[j], p[j-1]
		} else {
			break
		}
	}
}
