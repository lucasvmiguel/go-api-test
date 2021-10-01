# go-api-test

REST API in Golang to practice.

## Install

```
git clone git@github.com:lucasvmiguel/go-api-test.git
```

*You must have Golang installed and configured to work with this API.*

## Getting Started

First, generate the database schema and Golang ORM code:
```
make db-migrate
make db-generate
```

Then, run the development server:

```bash
make run-api
```

## Stack

* Language: `Golang`
* API/REST framework: `chi`
* Database ORM: `Prisma`
* Config reader: `godotenv`


## Testing

### How to run unit
```
make test-unit
```

## Configuration

All config are passed using environment variables, see all them below:
Env Var | Default value |
--- | --- |
`PORT` | 8080 |

## Folder/File struct

* `/db`: All ORM entities are auto generated using prisma.
* `/cmd`: Main applications for this project.
* `/internal`: Private application and library code.
* `/pkg`: Library code that's ok to use by external applications.
* `/.github`: CI/CD from Github.
* `schema.prisma`: Database schema.
* `dev.db`: Database used for development.
* `Makefile`: Project's executable tasks.

*Reference: https://github.com/golang-standards/project-layout*

## CI/CD

Check the Github actions for this repository, but in a nutshell:
1. Set up Go
2. Build
3. Test
4. Log in to the Container registry (Github)
5. Build and push Docker images