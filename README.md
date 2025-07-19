# pacbot!
pacbot is a simple discord bot that provides you with information about packages in the various arch package repositories on command. It's written in Go because that's the language I'm learning right now (also I really like its standard library).

none of it works yet! im mostly just tinkering right now. I'm leaving it public so my friends can see the source code when i stick it in their server.

## Usage:
- `!pacbot "QUERY"`: search the core, extra and multilib repositories for packages that include the string you input.
- `!pac "QUERY"`: search only official repositories
- `!core, !extra, !multilib, !aur "QUERY"`: search the given repository.
