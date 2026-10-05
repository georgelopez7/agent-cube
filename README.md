<img src="./docs/assets/agentcube-recording.gif" width="100%" alt="Agent Cube demo recording" />

An application to test the **Rubik's Cube** solving capabilities of **Large Language Models (LLMs).**

## Getting Started

**Copy environment variables:**

```bash
cp .env.sample .env

OPENROUTER_API_KEY=<add-openrouter-api-key>
```

**Spin up services:**

```bash
make run
```

This following services will be available:

| Service | URL |
|---|---|
| Agent Cube App | [http://localhost:3000](http://localhost:3000) |
| Agent Cube API | [http://localhost:8080](http://localhost:8080) |
| Arize Phoenix | [http://localhost:6006](http://localhost:6006) |

Go to [Agent Cube App](http://localhost:3000) - create a **Rubik's Cube** and watch the **agent** attempt to **solve it!**

## Features

#### Stream Agent Reasoning

Stream agent reasoning to see live reasoning from the LLM.

<img src="./docs/assets/agentcube-reasoning.png" width="400" alt="Live agent reasoning stream" />

#### Token and Cost Tracking

Token and cost tracking in both the app and through **Arize**.

<img src="./docs/assets/agentcube-usage.png" width="400" alt="Token and cost tracking" />

#### Decisions API Support

Support for models like Jev through the Decisions API.
