# What the heck is a FTP server

## Table of Contents 
1. [What it does ??](#what-it-does)
2. [Installation](#installation)
3. [Execution](#execution)
4. [Usage](#usage)
5. [Commands](#commands)

## What It does
> I built it from scratch let me tell you !!

File Transfer Protocol (FTP) servers are essential tools for transferring files over the internet. They operate on a client-server model, allowing users to upload and download files between their local devices and remote servers, facilitates file transfer across a TCP/IP network. It uses the File Transfer Protocol to establish secure connections between client devices and the server, enabling efficient data transmission. 

Two Major functions include ->

- Uploading file : Users can send files from their local devices to the server.
- Downloading file : Users can retrieve files from the server to their local devices.

## Installation

1. Cloning the repository ``git clone github.com/mrinalxdev/ftp-go.git``
2. Navigate to the directory ``cd ftp-go``
3. I haven't used any additional packages built on standard libraries :).. if still shows any error go for ``go get tidy``
4. Built the project ``go build ftp_go``

## Execution 

- The binary file ``./ftp_go`` (or ``ftp_go.exe`` for Windows)
- The server will listen to port 2121 by default !

## Usage

- Use the hostname ``localhost`` and port ``2121``
- Authenticate with the username and password (default : ``anonymous``/``anonymous``)

## Commands

The server supports the following FTP commands : 

- ``USER`` : Set Username
- ``PASS`` : Set password
- ``CWD`` : Change working directory
- ``LIST`` : List files in directory
- ``RETR`` : Retreive file
- ``STOR`` : Store File
- ``QUIT`` : Quit Connection


## Troubleshooting

- Problems may occur in Mac for Connection refused: Check server status and port number


Your valuable contributions are appreciated !!

