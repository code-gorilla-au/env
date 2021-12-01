# env
simple env vars, checks if the env var exists or returns the default provided


## Basic use

```go

 env := env.GetAsString("ENV", "dev")

 isFlagEnabled := env.GetAsBool("FEATURE_ONE", false)

 allowList := env.GetAsSlice("ALLOW_LIST", []string{"one", "two"}, ",")

```

## Load env file

```go

LoadEnvFile("./.env.local")

env := env.GetAsString("ENV", "dev")

```