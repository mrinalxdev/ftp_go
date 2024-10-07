# What the heck is a FTP server

> I built it from scratch let me tell you !!

## Table of Contents 
1. [What it does ??](#what-it-does)





## What It does

This FTP server is a basic implementation of the File Transfer Protocol (FTP) in golang from scratch. It allows users to upload, download, and manage files on the server

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

