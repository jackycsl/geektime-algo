package main

import "container/heap"

type Twitter struct {
	time       int
	tweets     map[int][]Tweet
	subscribes map[int]*Set
}

func Constructor() Twitter {
	return Twitter{
		tweets:     make(map[int][]Tweet),
		subscribes: make(map[int]*Set),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	var (
		tweets []Tweet
		ok     bool
	)
	if tweets, ok = this.tweets[userId]; ok {
		tweets = append(tweets, Tweet{
			id:        tweetId,
			timestamp: this.time,
		})
	} else {
		tweets = make([]Tweet, 0)
		tweets = append(tweets, Tweet{
			id:        tweetId,
			timestamp: this.time,
		})
	}
	this.tweets[userId] = tweets
	this.time++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	minHeap := &TweetMinHeap{}
	set, ok := this.subscribes[userId]
	if !ok {
		set = &Set{m: make(map[int]int)}
	}
	set.Add(userId)
	for uid, _ := range set.m {
		for _, tweet := range this.tweets[uid] {
			heap.Push(minHeap, tweet)
			if minHeap.Len() > 10 {
				heap.Pop(minHeap)
			}
		}
	}
	res := make([]int, minHeap.Len())
	i := minHeap.Len() - 1
	for minHeap.Len() > 0 {
		res[i] = heap.Pop(minHeap).(Tweet).id
		i--
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	var (
		subs *Set
		ok   bool
	)
	if subs, ok = this.subscribes[followerId]; ok {
		subs.Add(followeeId)
	} else {
		subs = &Set{m: map[int]int{}}
		subs.Add(followeeId)
	}
	this.subscribes[followerId] = subs
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if set, ok := this.subscribes[followerId]; ok {
		set.Remove(followeeId)
		this.subscribes[followerId] = set
	}
}

/**
* Your Twitter object will be instantiated and called as such:
* obj := Constructor();
* obj.PostTweet(userId,tweetId);
* param_2 := obj.GetNewsFeed(userId);
* obj.Follow(followerId,followeeId);
* obj.Unfollow(followerId,followeeId);
 */

type Tweet struct {
	id        int
	timestamp int
}

type Set struct {
	m map[int]int
}

func (set *Set) Add(v int) {
	set.m[v] = 1
}

func (set *Set) Remove(v int) {
	delete(set.m, v)
}

type TweetMinHeap []Tweet

func (minHeap *TweetMinHeap) Len() int {
	return len(*minHeap)
}

func (minHeap *TweetMinHeap) Less(i, j int) bool {
	return (*minHeap)[i].timestamp < (*minHeap)[j].timestamp
}

func (minHeap *TweetMinHeap) Swap(i, j int) {
	(*minHeap)[i], (*minHeap)[j] = (*minHeap)[j], (*minHeap)[i]
}

func (minHeap *TweetMinHeap) Push(x interface{}) {
	*minHeap = append(*minHeap, x.(Tweet))
}

func (minHeap *TweetMinHeap) Pop() interface{} {
	n := len(*minHeap) - 1
	x := (*minHeap)[n]
	*minHeap = (*minHeap)[:n]
	return x
}
