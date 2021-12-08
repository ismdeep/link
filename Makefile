SHELL=/bin/bash

help:

build-vendor:
	go mod tidy
	go mod vendor

test-renew:
	-docker stop link-db
	-docker rm   link-db
	docker run --name link-db \
		-p 10003:3306 \
		-e MYSQL_ROOT_PASSWORD=link123456 \
		-e MYSQL_DATABASE=link \
		-d hub.deepin.com/library/mysql:8.0
	waitdb -dialect mysql -dsn 'root:link123456@tcp(127.0.0.1:10003)/link?parseTime=true&loc=Local&charset=utf8mb4,utf8'
