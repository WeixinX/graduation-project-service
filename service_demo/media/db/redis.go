package db

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const (
	// RedisMaxSleepMs 最长睡眠时间 50 ms
	RedisMaxSleepMs = 50

	// RedisMinSleepMs 最短睡眠时间 20 ms
	RedisMinSleepMs = 20
)

func RedisPost(ctx *gin.Context, media Media) {
	// 提取 parent span context, 并创建子client span
	ctxWithSpan, ok := ctx.Get("ctxWithSpan")
	if !ok {
		return
	}
	span, _ := opentracing.StartSpanFromContext(ctxWithSpan.(context.Context), "/redis/post_media")
	defer span.Finish()

	// 该 span 异步(并行)
	span.SetTag("isAsync", "true")

	ext.HTTPUrl.Set(span, "/redis/post_media")
	ext.HTTPMethod.Set(span, http.MethodPost)

	// 模拟实际时延
	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	sleepTime := random.Intn(RedisMaxSleepMs-RedisMinSleepMs) + RedisMinSleepMs
	fmt.Printf("post redis sleep time: %v\n", time.Millisecond*time.Duration(sleepTime))
	fmt.Printf("media info:\n{user_id: %s, time_stamp: %s, content: %s}\n",
		media.UserID, media.TimeStamp, media.MediaContent)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
}
