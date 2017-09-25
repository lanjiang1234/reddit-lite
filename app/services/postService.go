package services

import (
	"github.com/lanjiang/reddit-lite/app/models"
	"time"
	"sort"
)

var (
	postList []*models.Post
	cursor int
)

func AddPost(p *models.Post) {
	p.PostId = buildPostId()
	p.CreatedOn = time.Now()
	postList = append(postList, p)
}

func GetAllPosts(count int) []*models.Post {
	if count > len(postList) {
		count = len(postList)
	}
	return postList[0:count]
}

func GetPostById(id int) *models.Post {
	for _, p := range postList {
		if (p.PostId == id) {
			return p
		}
	} 
	return nil
}

func VotePost(id int, voteType int) {
	p := GetPostById(id)
	switch(voteType) {
	case 0:
		p.DownVote = p.DownVote + 1;
	case 1:
		p.UpVote = p.UpVote + 1;
	}

	sortPostsByVote()
}

func sortPostsByVote() {
	sort.Sort(topicArray(postList))
}

func buildPostId() int {
	cursor++;
	return cursor;
}

type topicArray []*models.Post

func (c topicArray) Len() int {  
	return len(c)  
}  

func (c topicArray) Swap(i, j int) {  
	c[i], c[j] = c[j], c[i]  
}  

func (c topicArray) Less(i, j int) bool {  
	if c[i].UpVote > c[j].UpVote {
		return true
	} else if c[i].UpVote == c[j].UpVote {
		return c[i].CreatedOn.Before(c[j].CreatedOn)
	}
	return false;
}  