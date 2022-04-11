package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

func anomalyInjectHandler(ctx *gin.Context) {
	coresStr := ctx.DefaultQuery("cores", "-1")
	durationStr := ctx.DefaultQuery("duration", "-1")
	if coresStr == "-1" || durationStr == "-1" {
		fmt.Println("anomaly injector parameters error")
		ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": "anomaly injector parameters error"})
	}

	cores, _ := strconv.Atoi(coresStr)
	args := make([]string, cores+1, cores+1)
	for i, _ := range args {
		if i == 0 {
			args[i] = durationStr
		} else {
			args[i] = strconv.Itoa(i)
		}
	}

	cmd := exec.Command("/tmp/cpu", args...)
	if err := cmd.Start(); err != nil {
		fmt.Println("anomaly injector exec failed")
		ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": "anomaly injector exec failed"})
	}

	fmt.Printf("anomaly injector start, parameters: %v\n", args)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("anomaly injector start, parameters: %v", args),
	})
}
