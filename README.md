# Crucible

Crucible is an open-source coding contest and practice platform by the LOL Coding Club. It provides a space for students to hone their coding skills through various challenges and compete in contests. Contributions and feature suggestions are welcome!

There are two parts to this project: the client side where code will be written, and the server side which will compile this code and return the results.

## Table of Contents

1. [Client Side](#client-side)
    - [Web Application](#web-application)
    - [Technology Used for Web](#technology-used-for-web)
    - [Local Client Setup](#local-client-setup)
2. [Server Side](#server-side)
    - [Technology Used](#technology-used)
    - [Local Server Setup](#local-server-setup)

---

# Client Side

## Web Application

For now, we have gone with [this method](https://medium.com/front-end-weekly/how-to-build-your-own-codepen-app-a8a7140d52d7), but eventually, the plan is to make our own homebrewed solution.

Other possible client applications can include a **VSCode extension** or **nvim plugin**.

## Technology Used for Web

1. Vite + React
2. Tailwind CSS

## Local Client Setup

**Requirements:** Install [Node](https://nodejs.org/en/download/package-manager/current).

1. Start Client Server:

    ```bash
    npm run dev
    ```

---

# Server Side

This part will take the code from the client, run all the tests, and return the results.

## Technology Used

1. Golang
2. Docker

## Local Server Setup

**Requirements:** Install [Golang compiler](https://go.dev/doc/install), [Docker Desktop](https://www.docker.com/products/docker-desktop/), and [Make](https://gnuwin32.sourceforge.net/packages/make.htm).

1. Clone the repository:

    ```bash
    git clone https://github.com/lolclubwit/Crucible.git
    ```

2. Docker Setup:

    ```bash
    make build
    make create
    ```

    For testing, use:

    ```bash
    make test
    ```

3. Run the server:

    ```bash
    go run server.go
    ```
