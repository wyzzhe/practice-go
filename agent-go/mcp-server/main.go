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

// package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/mark3labs/mcp-go/mcp"
// 	"github.com/mark3labs/mcp-go/server"
// )

// // Define a struct for our typed arguments
// type GreetingArgs struct {
// 	Name      string   `json:"name"`
// 	Age       int      `json:"age"`
// 	IsVIP     bool     `json:"is_vip"`
// 	Languages []string `json:"languages"`
// 	Metadata  struct {
// 		Location string `json:"location"`
// 		Timezone string `json:"timezone"`
// 	} `json:"metadata"`
// }

// func main() {
// 	// Create a new MCP server
// 	s := server.NewMCPServer(
// 		"Typed Tools Demo 🚀",
// 		"1.0.0",
// 		server.WithToolCapabilities(false),
// 	)

// 	// Add tool with complex schema
// 	tool := mcp.NewTool("greeting",
// 		mcp.WithDescription("Generate a personalized greeting"),
// 		mcp.WithString("name",
// 			mcp.Required(),
// 			mcp.Description("Name of the person to greet"),
// 		),
// 		mcp.WithNumber("age",
// 			mcp.Description("Age of the person"),
// 			mcp.Min(0),
// 			mcp.Max(150),
// 		),
// 		mcp.WithBoolean("is_vip",
// 			mcp.Description("Whether the person is a VIP"),
// 			mcp.DefaultBool(false),
// 		),
// 		mcp.WithArray("languages",
// 			mcp.Description("Languages the person speaks"),
// 			mcp.Items(map[string]any{"type": "string"}),
// 		),
// 		mcp.WithObject("metadata",
// 			mcp.Description("Additional information about the person"),
// 			mcp.Properties(map[string]any{
// 				"location": map[string]any{
// 					"type":        "string",
// 					"description": "Current location",
// 				},
// 				"timezone": map[string]any{
// 					"type":        "string",
// 					"description": "Timezone",
// 				},
// 			}),
// 		),
// 	)

// 	// Add tool handler using the typed handler
// 	s.AddTool(tool, mcp.NewTypedToolHandler(typedGreetingHandler))

// 	// Start the stdio server
// 	if err := server.ServeStdio(s); err != nil {
// 		fmt.Printf("Server error: %v\n", err)
// 	}
// }

// // Our typed handler function that receives strongly-typed arguments
// func typedGreetingHandler(ctx context.Context, request mcp.CallToolRequest, args GreetingArgs) (*mcp.CallToolResult, error) {
// 	if args.Name == "" {
// 		return mcp.NewToolResultError("name is required"), nil
// 	}

// 	// Build a personalized greeting based on the complex arguments
// 	greeting := fmt.Sprintf("Hello, %s!", args.Name)

// 	if args.Age > 0 {
// 		greeting += fmt.Sprintf(" You are %d years old.", args.Age)
// 	}

// 	if args.IsVIP {
// 		greeting += " Welcome back, valued VIP customer!"
// 	}

// 	if len(args.Languages) > 0 {
// 		greeting += fmt.Sprintf(" You speak %d languages: %v.", len(args.Languages), args.Languages)
// 	}

// 	if args.Metadata.Location != "" {
// 		greeting += fmt.Sprintf(" I see you're from %s.", args.Metadata.Location)

// 		if args.Metadata.Timezone != "" {
// 			greeting += fmt.Sprintf(" Your timezone is %s.", args.Metadata.Timezone)
// 		}
// 	}

// 	return mcp.NewToolResultText(greeting), nil
// }
