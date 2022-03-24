package db

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// MongoMaxSleepMs 最长睡眠时间 100 ms
	MongoMaxSleepMs = 100

	// MongoMinSleepMs 最短睡眠时间 30 ms
	MongoMinSleepMs = 30
)

func MongoDBPost(media Media) {
	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	sleepTime := random.Intn(MongoMaxSleepMs-MongoMinSleepMs) + MongoMinSleepMs
	fmt.Printf("post mogodb sleep time: %v\n", time.Millisecond*time.Duration(sleepTime))
	fmt.Printf("media info:\n{user_id: %s, time_stamp: %s, content: %s}\n",
		media.UserID, media.TimeStamp, media.MediaContent)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
}
