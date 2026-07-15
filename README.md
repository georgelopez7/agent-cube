<img src="./docs/assets/hero-section.png" width="100%" alt="Agent Cube hero section" />

An application to test the **Rubik's Cube** solving capabilities of **Large Language Models (LLMs).**

## Technologies

- [TanStack](https://tanstack.com)
- [Go](https://go.dev)
- [Google ADK](https://google.github.io/adk/)
- [OpenRouter](https://openrouter.ai)

## Getting Started

**1. Copy the environment variables from `.env.sample`:**

```bash
cp .env.sample .env
```

**2. Open `.env` and fill in the required values:**

| Variable | Action |
|---|---|
| `OPENROUTER_API_KEY` | Visit [OpenRouter](https://openrouter.ai), generate an API key, and set this value. |
| `LANGFUSE_PUBLIC_KEY` | Generate a Langfuse public key and set this value. |
| `LANGFUSE_SECRET_KEY` | Generate a Langfuse secret key and set this value. |
| `LANGFUSE_INIT_PROJECT_PUBLIC_KEY` | Copy the same value as `LANGFUSE_PUBLIC_KEY`. |
| `LANGFUSE_INIT_PROJECT_SECRET_KEY` | Copy the same value as `LANGFUSE_SECRET_KEY`. |

`LANGFUSE_INIT_USER_NAME` and `LANGFUSE_INIT_USER_PASSWORD` determine the Langfuse login credentials and are already set in `.env.sample`. 

Use the defaults (`admin` / `123123123`) to log in.

**3. Spin up the services:**

```bash
docker compose up --build -d
```

This following services will be available:

| Service | URL |
|---|---|
| Agent Cube | [http://localhost:3001](http://localhost:3001) |
| Langfuse | [http://localhost:3000](http://localhost:3000) |

Go to [http://localhost:3001](http://localhost:3001) - create a **Rubik's Cube** and watch the **agent** attempt to solve it!

**4. (Optional) Extend the suite of large language models:**

The list of supported LLMs is defined in [`backend/internal/domain/ai.go`](./backend/internal/domain/ai.go).

Add a new entry to the `LLMs` slice to include another model:

```go
var LLMs = []LLM{
	NewLLM("openai", "openai/gpt-5.4-mini"),
	NewLLM("anthropic", "anthropic/claude-haiku-4.5"),
	NewLLM("google", "google/gemini-3.1-flash-lite"),
	NewLLM("moonshot", "moonshotai/kimi-k2.7-code"),
	NewLLM("x-ai", "x-ai/grok-4.3"),
	NewLLM("deepseek", "deepseek/deepseek-v4-pro"),
}
```

Then run the **build** command again:

```bash
docker compose up --build -d
```
