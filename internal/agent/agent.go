package agent

import (
	"agent/internal/mcp"
	"context"

	mcp_sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

type Role string

const (
	UserRole Role = "User"
	SystemRole Role = "System"
	AgentRole Role = "Agent"
	ToolRole Role = "Tool"
)

type Agent struct {
	LLM 		LLMClient
	MCPClient 	*mcp.MCPClient
}

type Message struct {
	Role Role
	Content string
}

type Context struct {
	Messages []Message
	Tools []mcp_sdk.Tool
}

func (a *Agent) Call(ctx context.Context, input string) (string, error) {
	tools, err := a.MCPClient.Tools(ctx)
	if err != nil {
		return "", err
	}
	agentContext := Context{
		Tools: tools,
	}
	// System prompt
	agentContext.Messages = append(agentContext.Messages, Message{
		Role: SystemRole,	
		Content: INITIAL_SYSTEM_PROMPT,
	})
	// User initial input
	agentContext.Messages = append(agentContext.Messages, Message{
		Role: UserRole,
		Content: input,
	})

	for {
		chatOutput, err := a.LLM.Chat(ctx, agentContext)
		if err != nil {
			return "", err
		}

		if IsToolCall(chatOutput) {
			toolName, params, err := ToToolCall(chatOutput)
			if err != nil {
				return "", err
			}

			toolOutput, err := a.MCPClient.CallTool(ctx, toolName, params)
			if err != nil {
				return "", err
			}
			toolMessage := Message{
				Role: ToolRole,
				Content: toolOutput,
			}
			agentContext.Messages = append(agentContext.Messages, toolMessage)
			continue
		}

		if IsText(chatOutput) {
			return chatOutput, nil
		}
	}
}

