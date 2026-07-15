# Agent Cube

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

`LANGFUSE_INIT_USER_NAME` and `LANGFUSE_INIT_USER_PASSWORD` determine the Langfuse login credentials and are already set in `.env.sample`. Use the defaults (`admin` / `123123123`) to log in.

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
