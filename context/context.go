package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 1*time.Millisecond)
	defer cancel()

	select {
	// 通道阻塞直到读取到数据
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	// 通道被关闭读取操作返回nil或0
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
