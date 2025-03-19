# Telegram Bot with Go

This repository contains a Telegram bot built with Go that interacts with an AI engine to generate Twitter posts. The bot allows users to invoke it with a command, generate a post using an AI engine, and post it to Twitter upon user confirmation.

## Features

1. User invokes the Telegram bot with `/post <content>`.
2. The bot sends the content to an AI engine (ChatGPT, Ollama, or a local LLM) to generate a Twitter post.
3. The bot sends the generated post back to the user via Telegram.
4. The user confirms whether to post it or not.
5. Upon confirmation, the post is made to Twitter.

## Project Structure
```
.
├── config
│   └── config.go
├── config.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── llm
│   └── llm.go
├── main.go
├── readme.md
├── telegram
│   └── bot.go
└── x
    └── x.go
```

## Dependencies

The project relies on the following dependencies:

- Go 1.16+
- Viper for configuration management
- Telegram Bot API
- Twitter API
- AI engine (ChatGPT, Ollama, or a local LLM)

## Configuration

The configuration is managed using `viper` and is stored in `config.yaml`. The configuration file includes the following fields:

```yaml
tgToken: "your-telegram-bot-token"
llm: "ollama" # or "chatgpt"
model: "llama3:latest"
llmToken: "your-llm-token"
preamble: "Rewrite the following content into a concise, engaging Twitter post. Keep it within 280 characters, maintain the original intent and emotion, and make it attention-grabbing. Use a clear, conversational tone, and include relevant emojis if appropriate and exclude hashtags. Use simple indian english\n\nUser content: "
xAPIKey: "your-twitter-api-key"
xAPISecret: "your-twitter-api-secret"
```

## Installation

1. Clone the repository:

```sh
git clone https://github.com/svnaditya/telegram-x-bot.git
cd telegram-x-bot
```

2. Install dependencies:

```sh
go mod download
```

3. Build the application:

```sh
go build -o telegram-bot-with-go main.go
```

4. Set up environment variables:

```sh
export TG_TOKEN="your-telegram-bot-token"
export LLM="ollama" # or "chatgpt"
export MODEL="llama3:latest"
export LLM_TOKEN="your-llm-token"
export X_API_KEY="your-twitter-api-key"
export X_API_SECRET="your-twitter-api-secret"
```

5. Run the application:

```sh
./telegram-bot-with-go
```

## Usage

1. Start a chat with your Telegram bot.
2. Invoke the bot with the command `/post <content>`.
3. The bot will send the content to the AI engine to generate a Twitter post.
4. The bot will send the generated post back to you via Telegram.
5. Confirm whether to post it or not.
6. Upon confirmation, the post will be made to Twitter.

## Docker

The project includes a Dockerfile to containerize the application. Follow these steps to build and run the application using Docker:

1. Build the Docker image:

```sh
docker build -t telegram-bot-with-go .
```

2. Run the Docker container:

```sh
docker run --network host -it telegram-bot-with-go
```

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

Please make sure your code adheres to the project's coding standards and includes appropriate tests.