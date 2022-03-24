package db

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// RedisMaxSleepMs 最长睡眠时间 50 ms
	RedisMaxSleepMs = 50

	// RedisMinSleepMs 最短睡眠时间 20 ms
	RedisMinSleepMs = 20
)

func RedisPost(media Media) {
	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	sleepTime := random.Intn(RedisMaxSleepMs-RedisMinSleepMs) + RedisMinSleepMs
	fmt.Printf("post redis sleep time: %v\n", time.Millisecond*time.Duration(sleepTime))
	fmt.Printf("media info:\n{user_id: %s, time_stamp: %s, content: %s}\n",
		media.UserID, media.TimeStamp, media.MediaContent)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
}
