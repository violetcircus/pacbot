# pacbot!
pacbot is a simple discord bot that provides you with information about packages in the various arch package repositories on command. It's written in Go because that's the language I'm learning right now (also I really like its standard library).

none of it works yet! im mostly just tinkering right now. I'm leaving it public so my friends can see the source code when i stick it in their server. this is 3 am experiment code! gaze upon it at your own peril.

The only external library I'm using is gorilla/websockets. Ideally I'd have written my own websocket implementation, but that felt a little excessive given I've never worked with them before and this was supposed to be a fun little quick project. I'm not  using one of the various discord API libraries because I want to learn about interacting with websockets directly.

## Usage:
- `!pacbot "QUERY"`: search the core, extra and multilib repositories for packages that include the string you input.
- `!pac "QUERY"`: search only official repositories
- `!core, !extra, !multilib, !aur "QUERY"`: search the given repository.
