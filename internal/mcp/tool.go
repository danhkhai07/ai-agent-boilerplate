package mcp

import (
	mcp_sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

type Tool struct {
	SDKTool mcp_sdk.Tool
	Handler mcp_sdk.ToolHandlerFor[any, any]
}

func bindTools(server *mcp_sdk.Server) {
	timeTool := mcp_sdk.Tool{
		Name: "checkTime",
		Description: "get the time at the current moment in hanoi, nyc, beijing",
	}

	mcp_sdk.AddTool(server, &timeTool, checkTime)
}
