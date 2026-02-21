# Go Workshop

This repository contains examples of various Go concepts, with each branch focusing on a specific topic.

## Setup

This repository includes a devcontainer setup to work with Go in vscode. Review the following to launch the environment.

1) First, please clone the repo and open it in visual studio code. 
   
2) Make sure to have the remote development extension pack installed and run `Dev Containers: Rebuild and Reopen in Container` in the visual studio command pallete. 

3) The DevContainer will bind mount your git directory into the container.

4) Open the terminal and switch to a branch of your choice to explore the basics of Go.

For more information on devcontainers please visit: https://code.visualstudio.com/docs/devcontainers/containers.

## Branches

### main
- A clean slate for Go Development. A devcontainer is included so users can hop in to work on Go. 

### hello-world
- A simple branch demonstrating the basic structure of a Go script including package main and function main.

### create-module
- A branch for developing new modules within the project.

### packages
- This branch demonstrates the use of mutliple packages in a go module.

### variables
- A branch dedicated to managing and implementing variables used in the project.

### collections
- Demonstrates the use of collections like arrays, slices, and maps in Go.

### control_flows
- Shows the control flow structures available like if, else, for, defer, panic, and switch statements.

### functions
- Demonstrates the use of multiple functions across packages in a Go module.

### read_write
- Demonstrates basic read and write utilities from the os standard library in Go.

### type_definitions
- Demonstrate creation of type aliases, custom types, structs, type embedding, factories, and interfaces in go

### go_routines
- Demonstrates the use of goroutines for multi-threaded applications and channels for communication and synchronization between go routines.

### api_client
- Demonstrates the creation of an api client to fetch json data over http

### wait_groups
- Expands on the api_client branch to demonstrate the creation of concurrent threads for api requests and synchronization with wait groups.

### testing
- Expands on the wait_groups branch with an example of how to run tests using the built-in go testing library.