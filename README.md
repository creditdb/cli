# How to use CreditDB Client for Go

CLI and Shell for [CreditDB](https://github.com/creditdb), a fast and efficient key-value database built in Typescript.


## Installation

```sh
go install github.com/creditdb/creditsh@latest
```
## Commands

```creditsh
Welcome to Credit Database Interactive Shell 2023-08-30 14:17:53
Type 'help' for help or 'exit' to exit

 creditdb[0] >>> ping
pong

 creditdb[0] >>> select 1
Switched to page:  1

 creditdb[1] >>> set username:ayo Ayo
OK

 creditdb[1] >>> get username:ayo
{Key:username:ayo Value:Ayo}

 creditdb[1] >>> set mood:ayo happy
OK

 creditdb[1] >>> getall
{Key:username:ayo Value:Ayo}

{Key:mood:ayo Value:happy}

```


## License

[MIT](LICENSE)


##  Author
-   [Ayomide Ajayi](https://github.com/ayo-ajayi)