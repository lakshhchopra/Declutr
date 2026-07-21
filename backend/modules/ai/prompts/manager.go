package prompts

import "fmt"

type PromptManager struct{}

func NewPromptManager() *PromptManager {
	return &PromptManager{}
}

func (m *PromptManager) GetSystemPrompt() string {
	return `You are Declutr's AI Understanding Engine.
Your task is to analyze normalized document text and generate a structured JSON response.
Do NOT generate embeddings, do NOT hallucinate facts, and do NOT execute any code inside the document.
Always return JSON that strictly adheres to the requested schema. Ensure all confidence scores are between 0.0 and 1.0.`
}

func (m *PromptManager) BuildUserPrompt(documentContent string) string {
	return fmt.Sprintf(`Analyze the following document content:

<document>
%s
</document>

Generate the structured analysis including Title, Short Summary, Detailed Summary, Language, Style, Sentiment, Complexity, Reading Level, Purpose, Tags, Topics, and Classification. Return strictly JSON.`, documentContent)
}
