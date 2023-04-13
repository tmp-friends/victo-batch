## victo-batch

- ローカル実行
  - 参照: https://github.com/GoogleCloudPlatform/functions-framework-go

```sh
# 実行対象の関数名を環境変数として与える
$ export FUNCTION_TARGET=InsertFanartTweets
# サーバ起動
$ go run cmd/main.go

# 実行
$ curl localhost:28080
```

- 実行対象の関数
  - https://github.com/tmp-friends/victo-batch/blob/main/functions/function.go#L22
