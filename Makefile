NAME=plansource
DEV_FILE=deploy/dev/docker-compose.yml

BASE_TAG=shaneburkhart/plansource

GOPATH=$(shell pwd)/vendor

IMG?=web
CMD?=/bin/bash

all: run

build:
	sudo docker build -t ${BASE_TAG} .

run:
	sudo docker-compose -f ${DEV_FILE} -p ${NAME} run --rm web goose -path="config/db" up
	sudo docker-compose -f ${DEV_FILE} -p ${NAME} up -d

stop:
	sudo docker-compose -f ${DEV_FILE} -p ${NAME} stop

clean:
	sudo docker rm $$(sudo docker ps -aq)

logs:
	sudo docker-compose -f ${DEV_FILE} -p ${NAME} logs

dependency:
	#sudo docker-compose -f ${DEV_FILE} -p ${NAME} run --rm web go get ${DEP}
	GOPATH=${GOPATH} go get ${DEP}
	sudo chown -R shane:shane vendor
	find vendor -name .git | xargs rm -rf

pg:
	echo "Enter 'postgres'..."
	sudo docker-compose -f ${DEV_FILE} -p ${NAME} run --rm pg psql -h pg -d mydb -U postgres --password

command:
	sudo docker-compose -f ${DEV_FILE} -p ${NAME} run --rm ${IMG} ${CMD}

gulp:
	sudo docker-compose -f ${DEV_FILE} -p ${NAME} run --rm web npm install
	sudo docker-compose -f ${DEV_FILE} -p ${NAME} run --rm web ./node_modules/.bin/gulp

gulp_watch:
	sudo docker-compose -f ${DEV_FILE} -p ${NAME} run --rm web npm install
	sudo docker-compose -f ${DEV_FILE} -p ${NAME} run --rm web ./node_modules/.bin/gulp watch
