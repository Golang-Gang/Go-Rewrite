# Go-Rewrite

## Background

For this Alchemy Code Lab assignment, we were tasked with rewritting any previous lab in a new http framework. Our team wanted a bigger challenge, so we rewrote the assignment in Go. We got permission to do this, but only if we met the other requirements, which included making the existing JS tests pass and structuring our code to have multiple controllers and db models. 

## Setup Steps

- Clone repo to local system.
- Copy and update `.env` file.
- In the Command Line:
  - `go get github.com/gorilla/mux`
  - `go get github.com/joho/godotenv`
  - `go get github.com/lib/pq`

## Execution and Testing

- Execute: `go run main.go`
- Test: `go test -v`


## Resources

- [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL - Kulshekhar Kabra](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)
- Languages/Packages:
  - [mux Github](https://github.com/gorilla/mux) / [mux Package](https://pkg.go.dev/github.com/gorilla/mux#section-readme)
  - [godotenv Github](https://github.com/joho/godotenv) / [godotenv Package](https://pkg.go.dev/github.com/joho/godotenv)
  - [pq Package](https://pkg.go.dev/github.com/lib/pq)
  - [PostgreSQL](https://www.postgresql.org/)
  - [Golang](https://go.dev/)
- Helpful Links:
  - [Learn GO in 12 Minutes](https://youtu.be/C8LgvuEBraI)
  - [Importing Local Files in Golang](https://www.youtube.com/watch?v=Nv8J_Ruc280)
  - [Golang Error Handling](https://youtu.be/VMveb4GqRck)
  - [Golang Middleware](https://youtu.be/HOlklLaFgfM)
