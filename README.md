# sample-golang-grpc-dynamodb


## Start api server

```shell
$ docker compose up -d --build
```

## Call grpc from local PC

```shell
# ls
$ grpc_cli ls localhost:50051 janken.JankenService

# call 
$ grpc_cli call localhost:50051 janken.JankenService.PlayJanken 'hands: 1'
$ grpc_cli call localhost:50051 janken.JankenService.PlayJankenResults ''
```
