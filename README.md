# GoApi

A simple API written in Go.

## Prerequisites

 * (https://git-scm.com/download/win) You will need Git to clone this, unless you just want to grab the zip file from Github :)
 * (https://golang.org/dl) Go is a little necessary :P

## Getting Started

run "git clone https://github.com/AaronCTech/GoApi.git" in your favorite project directory on your computer to download the project. GitHub will also allow a direct download via a zip file from the project.

## Running the application

To run the application, open your favorite command line tool (CMD/Powershell/Bash), enter the GoApi directory, and run "go run main.go". This will invoke Go's built in run tool. Visit localhost:8280 in your browser, and behold!

## Building an executable

To build a binary, run "go build ." from the GoApi directory, and you will see a new file "main.exe" appear if you are on Windows, or "main" appear on Linux. These are standalone applications, and can be executed on ANY system without the need to install any dependencies.

It should be noted that this service uses an absolutely mind-blowingly low amount of memory (2mb) due to its being compiled down to a system-level executable.