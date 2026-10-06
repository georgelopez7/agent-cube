# Seed 10 Rubik's cubes using the configured LLM rotation.

POST {{host}}/api/v1/rubiks-cubes
Content-Type: application/json
{
    "llm": { "provider": "openai", "model": "openai/gpt-5.4-mini" },
    "scramble": 1,
    "max_duration_ms": 30000
}
HTTP 201
[Asserts]
jsonpath "$.cube.id" exists

POST {{host}}/api/v1/rubiks-cubes
Content-Type: application/json
{
    "llm": { "provider": "anthropic", "model": "anthropic/claude-haiku-4.5" },
    "scramble": 2,
    "max_duration_ms": 30000
}
HTTP 201
[Asserts]
jsonpath "$.cube.id" exists

POST {{host}}/api/v1/rubiks-cubes
Content-Type: application/json
{
    "llm": { "provider": "google", "model": "google/gemini-3.1-flash-lite" },
    "scramble": 3,
    "max_duration_ms": 30000
}
HTTP 201
[Asserts]
jsonpath "$.cube.id" exists

POST {{host}}/api/v1/rubiks-cubes
Content-Type: application/json
{
    "llm": { "provider": "moonshot", "model": "moonshotai/kimi-k2.7-code" },
    "scramble": 4,
    "max_duration_ms": 30000
}
HTTP 201
[Asserts]
jsonpath "$.cube.id" exists

POST {{host}}/api/v1/rubiks-cubes
Content-Type: application/json
{
    "llm": { "provider": "x-ai", "model": "x-ai/grok-4.3" },
    "scramble": 5,
    "max_duration_ms": 30000
}
HTTP 201
[Asserts]
jsonpath "$.cube.id" exists

POST {{host}}/api/v1/rubiks-cubes
Content-Type: application/json
{
    "llm": { "provider": "deepseek", "model": "deepseek/deepseek-v4-pro" },
    "scramble": 6,
    "max_duration_ms": 30000
}
HTTP 201
[Asserts]
jsonpath "$.cube.id" exists

POST {{host}}/api/v1/rubiks-cubes
Content-Type: application/json
{
    "llm": { "provider": "openai", "model": "openai/gpt-5.4-mini" },
    "scramble": 7,
    "max_duration_ms": 30000
}
HTTP 201
[Asserts]
jsonpath "$.cube.id" exists

POST {{host}}/api/v1/rubiks-cubes
Content-Type: application/json
{
    "llm": { "provider": "anthropic", "model": "anthropic/claude-haiku-4.5" },
    "scramble": 8,
    "max_duration_ms": 30000
}
HTTP 201
[Asserts]
jsonpath "$.cube.id" exists

POST {{host}}/api/v1/rubiks-cubes
Content-Type: application/json
{
    "llm": { "provider": "google", "model": "google/gemini-3.1-flash-lite" },
    "scramble": 9,
    "max_duration_ms": 30000
}
HTTP 201
[Asserts]
jsonpath "$.cube.id" exists

POST {{host}}/api/v1/rubiks-cubes
Content-Type: application/json
{
    "llm": { "provider": "moonshot", "model": "moonshotai/kimi-k2.7-code" },
    "scramble": 1,
    "max_duration_ms": 30000
}
HTTP 201
[Asserts]
jsonpath "$.cube.id" exists
