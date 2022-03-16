# env
simple env vars, checks if the env var exists or returns the default provided

credit goes to https://github.com/17twenty for initial authoring


## Basic use

```go

 env := env.GetAsString("ENV", "dev")

 isFlagEnabled := env.GetAsBool("FEATURE_ONE", false)

 allowList := env.GetAsSlice("ALLOW_LIST", []string{"one", "two"}, ",")

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
foo := env.GetAsString("ENV", "dev")

```