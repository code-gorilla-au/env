# env
simple env vars, checks if the env var exists or returns the zero value

credit goes to https://github.com/17twenty for initial authoring

[![Go Report Card](https://goreportcard.com/badge/github.com/code-gorilla-au/env)](https://goreportcard.com/report/github.com/code-gorilla-au/env)
[![Go Reference](https://pkg.go.dev/badge/github.com/code-gorilla-au/env.svg)](https://pkg.go.dev/github.com/code-gorilla-au/env)

## Basic use

```go

 env := env.GetAsString("ENV", "dev")

 isFlagEnabled := env.GetAsBool("FEATURE_ONE")

 allowList := env.GetAsSlice("ALLOW_LIST", ",")

```

## Load env file

```go

LoadEnvFile("./.env.local")

foo := env.GetAsString("ENV", "dev")

```

## Strict mode

```go

env.WithStrictMode()

// ENV goes not exist
// panics
foo := env.GetAsString("ENV")

```