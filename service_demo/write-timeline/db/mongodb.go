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
	// MongoMaxSleepMs 最长睡眠时间 100 ms
	MongoMaxSleepMs = 100

	// MongoMinSleepMs 最短睡眠时间 30 ms
	MongoMinSleepMs = 30
)

func MongoDBPost(ctx *gin.Context, text Text) {
	// 提取 parent span context, 并创建子client span
	ctxWithSpan, ok := ctx.Get("ctxWithSpan")
	if !ok {
		return
	}
	span, _ := opentracing.StartSpanFromContext(ctxWithSpan.(context.Context), "/mongo/post_text")
	defer span.Finish()

	// 该 span 异步(并行)
	span.SetTag("is-async", "true")

	ext.HTTPUrl.Set(span, "/mongo/post_text")
	ext.HTTPMethod.Set(span, http.MethodPost)

	// 模拟实际时延
	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	sleepTime := random.Intn(MongoMaxSleepMs-MongoMinSleepMs) + MongoMinSleepMs
	fmt.Printf("post mogodb sleep time: %v\n", time.Millisecond*time.Duration(sleepTime))
	fmt.Printf("text info:\n{user_id: %s, time_stamp: %s, content: %s}\n",
		text.UserID, text.TimeStamp, text.Content)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
}
