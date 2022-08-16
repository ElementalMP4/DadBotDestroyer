# DadBotDestroyer
A bot designed to cripple Dad Bot

## Example config file

```json
{
    "token": "your-token-here",
    "users": ["503720029456695306", "user-you-wanna-delete"],
    "guilds": ["your-guild-id"]
}
```

## What does DadBotDestroyer do?

Whenever you send a message saying "I'm [something]" DadBot will make a very annoying dad joke. This bot allows you to delete them to make the chat look a bit less hectic.

When DadBot (or any other users in your config) send a message, you can react to that message with a ❌ to delete it.

If you run DadBotDestroyer behind a bot which is already in several guilds, you can restrict the program to certain guilds in the config.

## Running DadBotDestroyer

1) Use the example config to make your own `config.json` file
2) Pull this repo and run `go build` to create an executable (works on windows and linux)
3) Run the executable
