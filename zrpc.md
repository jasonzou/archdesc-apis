# go-zero internal calls

handler/routes.go -> (/login) loginhandler.go -> loginlogic.go
  |---> (/users) getuserhandler.go -> getuserlogic.go
  |---> (/users/id/?id) updateuserhandler.go -> updateuserlogic.go

## rpc

- goctl rpc template -o user.proto to produce user.proto
