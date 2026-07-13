package agent

import (
	"agent/internal/domain"
	"agent/internal/mcp"
	"context"
)

type Agent struct {
	LLM 		LLMClient
	MCPClient 	*mcp.MCPClient
}

func NewAgent(llm LLMClient, mcpClient *mcp.MCPClient) *Agent {
	return &Agent{
		LLM: llm,
		MCPClient: mcpClient,
	}
}

func (a *Agent) Call(ctx context.Context, input string, agentContext *domain.Context) (string, error) {
	tools, _ := a.MCPClient.Tools(ctx)
	agentContext.Tools = tools
	// System prompt
	if len(agentContext.Messages) == 0 {
		agentContext.Messages = append(agentContext.Messages, domain.Message{
			Role: domain.SystemRole,	
			Content: INITIAL_SYSTEM_PROMPT,
		})
		
	}
	// User initial input
	agentContext.Messages = append(agentContext.Messages, domain.Message{
		Role: domain.UserRole,
		Content: input,
	})

	for {
		chatOutput, err := a.LLM.Chat(ctx, *agentContext)
		if err != nil {
			return "", err
		}

		if IsToolCall(chatOutput) {
			toolOutput, err := a.MCPClient.CallTool(ctx, chatOutput.ToolName, chatOutput.Args)
			toolMessage := domain.Message{
				Role: domain.ToolRole,
			}
			if err != nil { toolMessage.Content = err.Error() } else { toolMessage.Content = toolOutput }
			agentContext.Messages = append(agentContext.Messages, toolMessage)
			continue
		}

		if IsText(chatOutput) {
			return chatOutput.Text, nil
		}
	}
}

