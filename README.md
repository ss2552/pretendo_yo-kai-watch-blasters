# Nintendo3DS 妖怪ウォッチバスターズの非公式サーバー

### https://github.com/PretendoNetwork/yo-kai-watch-blasters

## 構成
全ての構成情報は環境変数で掴む

`.env`ファイルでも可能

| 名前                                      | 概要                                                                                     | 必修                                      |
|-------------------------------------------|-------------------------------------------------------------------------------------------------|-----------------------------------------------|
| `PN_YKWB_KERBEROS_PASSWORD`          | Kerberosチケット内の内部サーバーデータの一部として使用されるパスワード                                    | いいえ (標準のパスワードが使用される) |
| `PN_YKWB_AUTHENTICATION_SERVER_PORT` | 認証サーバー                   ポート                                                              | はい|
| `PN_YKWB_SECURE_SERVER_HOST`         | セキュアサーバー               ホスト名（認証サーバーと同じアドレスを指す必要があります）                   | はい|
| `PN_YKWB_SECURE_SERVER_PORT`         | セキュアサーバー               ポート                                                              | はい|
| `PN_YKWB_ACCOUNT_GRPC_HOST`          | アカウントサーバーのgRPCサービス ホスト名                                                            | はい|
| `PN_YKWB_ACCOUNT_GRPC_PORT`          | アカウントサーバーのgRPCサービス ポート                                                              | はい|
| `PN_YKWB_ACCOUNT_GRPC_API_KEY`       | アカウントサーバーのgRPCサービス APIキー                                                             | いいえ (想定されるgRPC APIを使用する)|
| `PN_YKWB_FRIENDS_GRPC_HOST`          | フレンドサーバーのgRPCサービス   ホスト名                                                               | はい|
| `PN_YKWB_FRIENDS_GRPC_PORT`          | フレンドサーバーのgRPCサービス   ポート                                                                | はい|
| `PN_YKWB_FRIENDS_GRPC_API_KEY`       | フレンドサーバーのgRPCサービス   APIキー                                                               | いいえ (想定されるgRPC APIを使用する)|
| `PN_YKWB_POSTGRES_URI`               | PostgreSQLサーバーのマッチング   URI                                                                 | はい|
