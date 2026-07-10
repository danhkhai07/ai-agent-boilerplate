package agent

type LLMClient interface {
	 Chat(input string) (string, error)
}
