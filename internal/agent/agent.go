package agent

import (
	"agent/internal/mcp"
)

type Agent struct {
	LLM 		LLMClient
	MCPClient 	*mcp.MCPClient
}

type Message struct {
	Role string
	Content string
}

func (a *Agent) Call(input string) (error) {
	// messages := []Message{
	// 	{Role: "User", Content: input },
	// }

	for {
				
	}
}
