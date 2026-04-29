# Nintendo3DS 妖怪ウォッチバスターズの非公式サーバー

### https://github.com/PretendoNetwork/yo-kai-watch-blasters

## Configuration
All configuration options are handled via environment variables

`.env` files are supported

| Name                                      | Description                                                                                     | Required                                      |
|-------------------------------------------|-------------------------------------------------------------------------------------------------|-----------------------------------------------|
| `PN_YKWB_KERBEROS_PASSWORD`          | Password used as part of the internal server data in Kerberos tickets                           | No (Default password `password` will be used) |
| `PN_YKWB_AUTHENTICATION_SERVER_PORT` | Port for the authentication server                                                              | Yes                                           |
| `PN_YKWB_SECURE_SERVER_HOST`         | Host name for the secure server (should point to the same address as the authentication server) | Yes                                           |
| `PN_YKWB_SECURE_SERVER_PORT`         | Port for the secure server                                                                      | Yes                                           |
| `PN_YKWB_ACCOUNT_GRPC_HOST`          | Host name for your account server gRPC service                                                  | Yes                                           |
| `PN_YKWB_ACCOUNT_GRPC_PORT`          | Port for your account server gRPC service                                                       | Yes                                           |
| `PN_YKWB_ACCOUNT_GRPC_API_KEY`       | API key for your account server gRPC service                                                    | No (Assumed to be an open gRPC API)           |
| `PN_YKWB_FRIENDS_GRPC_HOST`          | Host name for your friends server gRPC service                                                  | Yes                                           |
| `PN_YKWB_FRIENDS_GRPC_PORT`          | Port for your friends server gRPC service                                                       | Yes                                           |
| `PN_YKWB_FRIENDS_GRPC_API_KEY`       | API key for your friends server gRPC service                                                    | No (Assumed to be an open gRPC API)           |
| `PN_YKWB_POSTGRES_URI`               | URI for matchmaking PostgeSQL server                                          
