# **Oauth2_server**

This service is a incomplete version of an Oauth2 implementation server.
<br>

## Endpoints
- /token -> create token
- /keys -> keys list to the clientId in the request
- /introspection -> check token

## Improvements proposal:
- data management 
- test(Unit test, BDDs)
- more separation on each layer

curl -v -H "Authorization: Basic YTA4OTdlNmQwZWE5NGY1ODljMzgyNzhiY2E0ZTkzNDI6Yzk0ZGJkNTgyZDU5NGU4YWEwNDkzNGY5YzdlZjBmNTI=" -X POST localhost:8080/token