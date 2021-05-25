
# Prerequisites

## Terminal

* Most useful developper tool
* Any number of customizations
* On Windows: Linux Bash Shell, Powershell, Git Bash
* On macOS / Linux : Terminal (default)

## Vi(m)

* Bash text editor
* Use `:` to enter command mode
  * `w` to write file
  * `q` to quit
  * `q!` to quit without saving
  * `x` to write & quit
* Use `/` to search for text
* Use `i` to enter edit mode and `Esc` to exit it
* `vimtutor` is the best tutorial to learn

## Client VS Server

* Two parts of a distributed computing model:
  * Client requests the info and displays it
  * Server processes the request and services the result

## The IP protocol

* Send data from one computer to another over a network (ex: client/server)
* Use of IPV4 addresses (ex: 172.16.254.1), IPV6 also available but not much used
* Data packaged in IP packets with 2 sections
  * Header: IP version, addresses, TTL, ...
  * Data: the packet's content

## The HTTP protocol

* Application protocol for transmitting hypermedia documents (HTML)
* Two types of messages: *requests* & *responses*
* HTTP message split between *headers* & *body*
* HTTP response always contains
  * the *protocol* (HTTP/1.1)
  * a *status code* (200, 404, ...)
  * a *status text* (page not found)

## SSL/TLS & HTTPS

(Secure Sockets Layer / Transport Layer Security)
* Establish an encrypted link over a network
* Exchange of public & private keys to secure the exchange
  * Server sends SSL certificate + public key
  * Client checks the certificate & answers with an encrypted session key
  * Client & server exchange messages encrypted with the keys to authenticate
* SSL certificate has been certified by a renowned authority
* HTTPS: HTTP secured with SSL/TLS

## SSH - Secure SHell

* Cryptographic network protocol to operate network services securely over an unsecured network
* Exchange of public & private keys to secure the exchange
  * Client has the private key
  * Server needs to have the associated public key
  * Client & server exchange messages encrypted with the keys to authenticate

## The SFTP protocol

* Send files over SSH
* ex: deploy website to a server
* SFTP apps: FileZilla, Cyberduck, WinSCP, ...

## Git

* Distributed version control
* Users keep entire code & history locally, can make any change without internet
* Users create snapshots of current code (`commit`) associated to a hash code
* Users `push` commited code to the remote git server
* Multiple users can work on the same git project
* When two users modify the same code they have to `merge` the two codes

## Git commands

* `git init` : initialize a git repository
* `git status` : show the current status of the local git repo
* `git clone`: download a repository locally
* `git add [files]` : add the files to the git index
* `git commit -m "[message]"` : create a commit
* `git push -u origin master` : push commits to the distant repo
* `git pull` : pull changes from the distant repo

[Git cheatsheet](https://git-tower.com/blog/git-cheat-sheet/)

## Git platforms / tools

* Platforms (hosting)
  * Github.com: free public repository hosting
  * Bitbucket.com: free public / private repository hosting
  * GitLab: installs your own Git server anywhere, free public / private repository hosting
* Git UI tools
  * GitX (Mac) / GitG (Linux)
  * GitHub (Mac/Win/Linux)
  * SourceTree (Mac/Win/Linux)
  * Fork (Mac/Win)
  * GitKraken (Mac/Win/Linux)
  * Your terminal!

## Editors

* As a developer, your editor shall be your best friend
* Vim
* VsCode editor
* Atom
* Sublime Text
* TextWrangler 
* Notepad++
* WebStorm
* ...
