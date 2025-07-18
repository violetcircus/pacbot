# pacbot!
pacbot is a simple discord bot that provides you with information about packages in the various arch package repositories on command. It's written in Go because that's the language I'm learning right now (also I really like its standard library).

## Usage:
!pacbot "QUERY": search the core, extra and multilib repositories for packages that include the string you input.
!pac "QUERY" search only official repositories
!core, !extra, !multilib, !aur "QUERY": search the given repository.
