*** refatorar 
1) propriedades de conexao em arq properties
1) fazer arquivo de migracao para banco de dados
2) criar arquivo de rotas separadas
3) criar pacotes para classes

4) fazer case de microservico para save no kafka que salva no mongo
5) fazer case de microservico para consumer que salva no postgres

6) fazer case de microservico consumer que le o banco e sincroniza com elastic search
7) fazer case de microservico para pesquisas no elastic search

8) fazer case de microserviço no aws
9) fazer case de microserviço no gcp
=======================================================


1 - init mod:    go mod init server
2 - install pkg: go get .
3 - go run .

3 - build files: go build .
4 - run server:  ./server

=======================================================

call endpoints:
http://localhost:8090/api/client/list
http://localhost:8090/api/client/find/1

http://localhost:8090/api/seller/list
http://localhost:8090/api/seller/find/1