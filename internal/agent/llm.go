package agent

import "context"

type LLMClient interface {
	 Chat(ctx context.Context, agentContext Context) (string, error)
	 IsText(s string) bool
	 IsToolCall(s string) bool
	 ToText(s string) (string, error)
	 ToToolCall(s string) (string, map[string]any, error)
}
