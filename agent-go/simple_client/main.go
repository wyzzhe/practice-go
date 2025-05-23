package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	// Define command line flags
	stdioCmd := flag.String("stdio", "", "Command to execute for stdio transport (e.g. 'python server.py')")
	httpURL := flag.String("http", "", "URL for HTTP transport (e.g. 'http://localhost:8080/mcp')")
	flag.Parse()

	// Validate flags
	if (*stdioCmd == "" && *httpURL == "") || (*stdioCmd != "" && *httpURL != "") {
		fmt.Println("Error: You must specify exactly one of --stdio or --http")
		flag.Usage()
		os.Exit(1)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create client based on transport type
	var c *client.Client
	var err error

	if *stdioCmd != "" {
		fmt.Println("Initializing stdio client...")
		// Parse command and arguments
		args := parseCommand(*stdioCmd)
		if len(args) == 0 {
			fmt.Println("Error: Invalid stdio command")
			os.Exit(1)
		}

		// Create command and stdio transport
		command := args[0]
		cmdArgs := args[1:]

		// Create stdio transport with verbose logging
		stdioTransport := transport.NewStdio(command, nil, cmdArgs...)

		// Create client with the transport
		c = client.NewClient(stdioTransport)

		// // 读取子进程的标准错误输出
		// // Set up logging for stderr if available
		// if stderr, ok := client.GetStderr(c); ok {
		// 	go func() {
		// 		buf := make([]byte, 4096)
		// 		for {
		// 			n, err := stderr.Read(buf)
		// 			if err != nil {
		// 				if err != io.EOF {
		// 					log.Printf("Error reading stderr: %v", err)
		// 				}
		// 				return
		// 			}
		// 			if n > 0 {
		// 				fmt.Fprintf(os.Stderr, "[Server] %s", buf[:n])
		// 			}
		// 		}
		// 	}()
		// }
	} else {
		fmt.Println("Initializing HTTP client...")
		// Create HTTP transport
		httpTransport, err := transport.NewStreamableHTTP(*httpURL)
		if err != nil {
			log.Fatalf("Failed to create HTTP transport: %v", err)
		}

		// Create client with the transport
		c = client.NewClient(httpTransport)
	}

	// Start the client
	if err := c.Start(ctx); err != nil {
		log.Fatalf("Failed to start client: %v", err)
	}

	// Set up notification handler
	c.OnNotification(func(notification mcp.JSONRPCNotification) {
		fmt.Printf("Received notification: %s\n", notification.Method)
	})

	// Initialize the client
	fmt.Println("Initializing client...")
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "MCP-Go Simple Client Example",
		Version: "1.0.0",
	}
	initRequest.Params.Capabilities = mcp.ClientCapabilities{}

	serverInfo, err := c.Initialize(ctx, initRequest)
	if err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}

	// Display server information
	fmt.Printf("Connected to server: %s (version %s)\n",
		serverInfo.ServerInfo.Name,
		serverInfo.ServerInfo.Version)
	fmt.Printf("Server capabilities: %+v\n", serverInfo.Capabilities)

	// List available tools if the server supports them
	if serverInfo.Capabilities.Tools != nil {
		fmt.Println("Fetching available tools...")
		toolsRequest := mcp.ListToolsRequest{}
		toolsResult, err := c.ListTools(ctx, toolsRequest)
		if err != nil {
			log.Printf("Failed to list tools: %v", err)
		} else {
			fmt.Printf("Server has %d tools available\n", len(toolsResult.Tools))
			for i, tool := range toolsResult.Tools {
				fmt.Printf("  %d. %s - %s\n", i+1, tool.Name, tool.Description)
			}
		}
	}

	// List available resources if the server supports them
	if serverInfo.Capabilities.Resources != nil {
		fmt.Println("Fetching available resources...")
		resourcesRequest := mcp.ListResourcesRequest{}
		resourcesResult, err := c.ListResources(ctx, resourcesRequest)
		if err != nil {
			log.Printf("Failed to list resources: %v", err)
		} else {
			fmt.Printf("Server has %d resources available\n", len(resourcesResult.Resources))
			for i, resource := range resourcesResult.Resources {
				fmt.Printf("  %d. %s - %s\n", i+1, resource.URI, resource.Name)
			}
		}
	}

	fmt.Println("Client initialized successfully.")

	var operation string
	var x, y int

	fmt.Print("请输入运算符[add subtract multiply divide]:")
	fmt.Scanf("%s\n", &operation)
	fmt.Print("请输入两个整数（用空格分隔）:")
	fmt.Scanf("%d %d\n", &x, &y)

	toolReq := mcp.CallToolRequest{
		Params: struct {
			Name      string    `json:"name"`
			Arguments any       `json:"arguments,omitempty"`
			Meta      *mcp.Meta `json:"_meta,omitempty"`
		}{
			Name: "calculate",
			Arguments: map[string]any{
				"operation": operation,
				"x":         x,
				"y":         y,
			},
		},
	}

	toolResult, err := c.CallTool(ctx, toolReq)
	if err != nil {
		log.Fatalf("调用工具失败: %v", err)
	}
	result, err := json.Marshal(toolResult.Content)
	if err != nil {
		fmt.Printf("json.Marshal err=%s", err)
	}
	fmt.Printf("工具调用返回: %s\n", string(result))

	c.Close()
}

// parseCommand splits a command string into command and arguments
func parseCommand(cmd string) []string {
	// This is a simple implementation that doesn't handle quotes or escapes
	// For a more robust solution, consider using a shell parser library
	var result []string
	var current string
	var inQuote bool
	var quoteChar rune

	for _, r := range cmd {
		switch {
		case r == ' ' && !inQuote:
			if current != "" {
				result = append(result, current)
				current = ""
			}
		case (r == '"' || r == '\''):
			if inQuote && r == quoteChar {
				inQuote = false
				quoteChar = 0
			} else if !inQuote {
				inQuote = true
				quoteChar = r
			} else {
				current += string(r)
			}
		default:
			current += string(r)
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}
