package mcp

import (
	"context"
	"errors"
	"strings"

	mcp_sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

type MCPClient struct {
	SDKClient *mcp_sdk.Client
	URL string
	Transport mcp_sdk.Transport
	tools []mcp_sdk.Tool
	session *mcp_sdk.ClientSession
}

func (client *MCPClient) CreateClient(ctx context.Context) error {
	if client.SDKClient != nil {
		return nil
	}
	impl := mcp_sdk.Implementation{
		Name: "agent-client",
		Version: "v1.0.0",
	}
	client.SDKClient = mcp_sdk.NewClient(&impl, nil)
	client.Transport = &mcp_sdk.StreamableClientTransport{
		Endpoint: client.URL,
	}

	session, err := client.SDKClient.Connect(ctx, client.Transport, nil)
	if err != nil {
		return err
	}
	client.session = session
	return nil
}

func (client *MCPClient) Tools(ctx context.Context) ([]mcp_sdk.Tool, error) {
	if client.tools != nil {
		return client.tools, nil
	}

	tools := client.session.Tools(ctx, nil)
	result := make([]mcp_sdk.Tool, 0)
	for tool, err := range tools {
		if err == nil {
			result = append(result, *tool)
		}
	}
	client.tools = result
	return result, nil
}

func (client *MCPClient) CallTool(ctx context.Context, toolName string, args map[string]any) (string, error) {
	params := mcp_sdk.CallToolParams{
		Name: toolName,
		Arguments: args,
	}
	callResult, err := client.session.CallTool(ctx, &params)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	for _, content := range callResult.Content {
		if textContent, ok := content.(*mcp_sdk.TextContent); ok {
			b.WriteString(textContent.Text + "\n")
		} else {
			return "", errors.New("error marshalling text content")
		}
	}
	return b.String(), nil
}
