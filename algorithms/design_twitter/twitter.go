// Package design_twitter
// @file: twitter.go
// @date: 2021/2/7
package design_twitter

import (
	"container/heap"
	"time"
)

// An Item is something we manage in a priority queue.
type Item struct {
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
	value *Tweet
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value.time > pq[j].value.time
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	tweet := x.(*Tweet)
	item := &Item{value: tweet}

	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]

	return item.value
}

type Tweet struct {
	id   int
	time int
	next *Tweet
}

type User struct {
	id       int
	followed map[int]struct{}
	head     *Tweet
}

func NewUser(id int) User {
	followed := make(map[int]struct{}, 0)
	followed[id] = struct{}{}
	return User{
		id:       id,
		followed: followed,
		head:     nil,
	}
}

func (u *User) follow(userID int) {
	if _, ok := u.followed[userID]; !ok {
		u.followed[userID] = struct{}{}
	}
}

func (u *User) unfollow(userID int) {
	if userID == u.id {
		return
	}
	if _, ok := u.followed[userID]; ok {
		delete(u.followed, userID)
	}
}

func (u *User) post(tweetID int) {
	tweet := &Tweet{id: tweetID, time: int(time.Now().Unix())}
	tweet.next = u.head
	u.head = tweet
}

type Twitter struct {
	idToUser map[int]*User
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{idToUser: make(map[int]*User, 0)}
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := t.idToUser[userId]; !ok {
		user := NewUser(userId)
		t.idToUser[userId] = &user
	}
	user := t.idToUser[userId]
	user.post(tweetId)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	if _, ok := t.idToUser[userId]; !ok {
		return []int{}
	}
	for id := range t.idToUser[userId].followed {
		tweet := t.idToUser[id].head
		if tweet == nil {
			continue
		}
		heap.Push(&pq, tweet)
	}
	for len(pq) != 0 {
		if len(res) == 10 {
			break
		}
		tweet := heap.Pop(&pq).(*Tweet)
		res = append(res, tweet.id)
		if tweet.next != nil {
			heap.Push(&pq, tweet.next)
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	// 若 follower 不存在则创建
	if _, ok := t.idToUser[followerId]; !ok {
		user := NewUser(followerId)
		t.idToUser[followerId] = &user
	}
	// 若 followee 不存在则创建
	if _, ok := t.idToUser[followeeId]; !ok {
		user := NewUser(followeeId)
		t.idToUser[followeeId] = &user
	}
	t.idToUser[followerId].follow(followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := t.idToUser[followerId]; !ok {
		return
	}
	t.idToUser[followerId].unfollow(followeeId)
}

func PrintRes() {
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	twitter.GetNewsFeed(1)
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	twitter.GetNewsFeed(1)
	twitter.Unfollow(1, 2)
	twitter.GetNewsFeed(1)
}
