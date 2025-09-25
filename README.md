# CLI AI Agent in Go
This project demonstrates how to build a command-line AI agent using Go and the OpenRouter API.

---

## Setup

1. **Initialize Go module**  
```bash
go mod init your_project_name
```

2. **Create `.env` file**  
```bash
vim .env
```  
Add your OpenRouter API key:  
```
OPENROUTER_API_KEY=your_api_key_here
```

3. **Install OpenAI Go SDK**  
```bash
go get github.com/openai/openai-go/v2
```

4. **Create `main.go`**  
```bash
vim main.go
```  
- Define OpenRouter API URL.  
- Load API key from `.env` and check if defined.  
- Create LLM client using OpenAI SDK with base URL and API key.  
- Create a `messages` array and append your query using `openai.UserMessage`.  
- Choose a model and define chat completion parameters.  
- Create a `ctx` and call `client.ChatCompletions.New(ctx, params)`.  
- Print the AI response.

5. **Export environment variables**  
```bash
export $(cat .env | xargs)
```

6. **Run**  
```bash
go run main.go
```

---

## Version 2: Interactive Chat

Enhance the CLI to wait for user input in a loop:

- Create a scanner:  
```go
scanner := bufio.NewScanner(os.Stdin)
```

- Print a prompt symbol:  
```go
fmt.Print(">")
```

- Read input:  
```go
if !scanner.Scan() { break }
input := strings.TrimSpace(scanner.Text())
if input == "" { continue }
```

- Append input to `messages` array.  
- Send chat completion request.  
- Print AI response.  
- Loop back for the next input.

---

## Version 3: Memory and Colors

Enhance further with:

- Append AI response to `messages`.  
- Print message count.  
- Limit `messages` array to last 4 messages if it grows too large.  
- Add colors with `fatih/color`:  
```bash
go get github.com/fatih/color
```

- Use colors to differentiate:  
  - User queries  
  - AI responses  
  - Other info  

This improves readability and visual distinction in the chat.

---

## Run

```bash
go run main.go
```

Enjoy your interactive CLI AI agent!
