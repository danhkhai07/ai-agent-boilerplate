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
	AgentRole Role = "Assistant"
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

func NewAgent(llm LLMClient, mcpClient *mcp.MCPClient) *Agent {
	return &Agent{
		LLM: llm,
		MCPClient: mcpClient,
	}
}

func (a *Agent) Call(ctx context.Context, input string, agentContext *Context) (string, error) {
	tools, err := a.MCPClient.Tools(ctx)
	if err != nil {
		agentContext.Tools = tools
	} else {
		agentContext.Tools = nil
	}
	agentContext.Tools = tools
	// System prompt
	if len(agentContext.Messages) == 0 {
		agentContext.Messages = append(agentContext.Messages, Message{
			Role: SystemRole,	
			Content: INITIAL_SYSTEM_PROMPT,
		})
		
	}
	// User initial input
	agentContext.Messages = append(agentContext.Messages, Message{
		Role: UserRole,
		Content: input,
	})

	for {
		chatOutput, err := a.LLM.Chat(ctx, *agentContext)
		if err != nil {
			return "", err
		}

		if IsToolCall(chatOutput) {
			toolOutput, err := a.MCPClient.CallTool(ctx, chatOutput.ToolName, chatOutput.Args)
			toolMessage := Message{
				Role: ToolRole,
			}
			if err != nil {
				toolMessage.Content = err.Error()
			} else {
				toolMessage.Content = toolOutput
			}
			agentContext.Messages = append(agentContext.Messages, toolMessage)
			continue
		}

		if IsText(chatOutput) {
			return chatOutput.Text, nil
		}
	}
}

