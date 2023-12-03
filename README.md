# Online Library
This repo contains a simple application using Go. This project uses an Online Library implementation as an example.

# Features
- Get book list by genre :
    Getting list of books by it's genre / subject. Book list data retrieved from openlibrary API
- Update pickup time of books :
    To update date and time of when the book will be picked up by user. Once a book has a picked up date then it can't be changed until later it was picked up

# Installation
1. Clone the repository
    ```bash
        git clone https://github.com/rez-a-put/online-library.git
    ```
2. Change into project directory
    ```bash
        cd online-library
    ```
3. Set up your .env file based on .env.example
4. Set up your vendor folder
    ```bash
        go mod vendor
    ```

# Run the project
1. Open terminal
2. Go to project folder
3. Build application
    ```bash
        go build
    ```
4. Run application from terminal or run using go command
    ```bash
        ./online-library
    ```
    ```bash
        go run main.go
    ```

# Testing
1. Open terminal
2. Go to project folder
3. Run Go test
    ```bash
        go test -v ./...
    ```

# Contributing
1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a merge request