package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// 创建 MCP 服务器实例
	s := server.NewMCPServer(
		"Calculator Demo", // 服务名称
		"1.0.0",           // 协议版本
		server.WithResourceCapabilities(true, true), // 支持资源和工具
		server.WithLogging(),                        // 启用日志输出
	)

	// 定义一个计算器工具
	calculatorTool := mcp.NewTool("calculate",
		mcp.WithDescription("执行基本算术运算"),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description("运算类型"),
			mcp.Enum("add", "subtract", "multiply", "divide"),
		),
		mcp.WithNumber("x",
			mcp.Description("第一个操作数"),
		),
		mcp.WithNumber("y",
			mcp.Description("第二个操作数"),
		),
	)

	// 注册工具及其处理逻辑
	s.AddTool(calculatorTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// 将 Arguments any 断言为 map[string]any
		op := request.GetArguments()["operation"].(string)
		x := request.GetArguments()["x"].(float64)
		y := request.GetArguments()["y"].(float64)
		var result float64

		switch op {
		case "add":
			result = x + y
		case "subtract":
			result = x - y
		case "multiply":
			result = x * y
		case "divide":
			if y == 0 {
				return mcp.NewToolResultError("除数不能为 0"), nil
			}
			result = x / y
		default:
			return mcp.NewToolResultError("无效的运算类型"), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("%.2f", result)), nil
	})

	// 启动 MCP 服务器 (采用标准输入/输出流方式)
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("服务端启动出错: %v\n", err)
	}
}
