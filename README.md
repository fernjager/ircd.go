ircd.go
=======

A self-contained, minimal IRC server written in Google Go.

# Purpose
The purpose of the project is twofold:

1. Serve as a learning experience for the author.
2. Innovation through iteration, by reinventing the wheel using the latest and greatest technologies.
 
The end goal is to create an easily deployed IRCd requiring minimal configuration.

# Objectives
1. Implement minimum viable product
2. Abstract IRC command functionality such that it can be an independent library
3. Complete command implementation

# Progress
##Command Implementation Progress:

###Connection Registration
- PASS
- ~~NICK~~ partial: works, error handling
- ~~USER~~ partial: works, error handling
- OPER
- MODE
- SERVICE
- QUIT
- ~~PING~~
- ~~PONG~~

###Channel Operations
- JOIN
- PART
- MODE
- TOPIC
- NAMES
- LIST
- INVITE
- KICK

###Messaging
- PRIVMSG
- NOTICE

###Server queries and commands
- SQUIT
- MOTD
- LUSERS
- VERSION
- STATS
- LINKS
- TIME
- CONNECT
- TRACE
- ADMIN
- INFO

###Service Query and Commands
- SERVLIST
- SQUERY
- WHO
- ERROR
- WHOIS
- WHOWAS
- KILL

###Optional Features
- AWAY
- REHASH
- DIE
- RESTART
- WALLOPS
- USERHOST
- ISON
