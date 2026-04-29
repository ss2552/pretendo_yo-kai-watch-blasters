# Nintendo3DS 妖怪ウォッチバスターズの非公式サーバー

### https://github.com/PretendoNetwork/yo-kai-watch-blasters

## 構成
全ての構成情報は環境変数で掴む

`.env`ファイルでも可能

| 名前                                      | 概要                                                                                     | 必修                                      |
|-------------------------------------------|-------------------------------------------------------------------------------------------------|-----------------------------------------------|
| `PN_YKWB_KERBEROS_PASSWORD`          | Password used as part of the internal server data in Kerberos tickets                           | いいえ (標準のパスワードが使用される) |
| `PN_YKWB_AUTHENTICATION_SERVER_PORT` | Port for the authentication server                                                              | はい                                           |
| `PN_YKWB_SECURE_SERVER_HOST`         | Host name for the secure server (should point to the same address as the authentication server) | はい                                           |
| `PN_YKWB_SECURE_SERVER_PORT`         | Port for the secure server                                                                      | はい                                           |
| `PN_YKWB_ACCOUNT_GRPC_HOST`          | Host name for your account server gRPC service                                                  | はい                                           |
| `PN_YKWB_ACCOUNT_GRPC_PORT`          | Port for your account server gRPC service                                                       | はい                                           |
| `PN_YKWB_ACCOUNT_GRPC_API_KEY`       | API key for your account server gRPC service                                                    | いいえ (Assumed to be an open gRPC API)           |
| `PN_YKWB_FRIENDS_GRPC_HOST`          | Host name for your friends server gRPC service                                                  | はい                                           |
| `PN_YKWB_FRIENDS_GRPC_PORT`          | Port for your friends server gRPC service                                                       | はい                                           |
| `PN_YKWB_FRIENDS_GRPC_API_KEY`       | API key for your friends server gRPC service                                                    | いいえ (Assumed to be an open gRPC API)           |
| `PN_YKWB_POSTGRES_URI`               | URI for matchmaking PostgeSQL server                                          
                | はい                                           |
