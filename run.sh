#!/bin/bash

PET_COUNT=100
SPECIALISTS=8
RECEPCIONISTS=1 # limited to 1 due to port, can be fixed if put behind a loadBalancer
API_HOST=localhost
API_PORT=8085

docker-compose stop recepcionist
docker-compose stop specialist
docker-compose up -d --scale specialist=${SPECIALISTS} --scale recepcionist=${RECEPCIONISTS}

until nc -z ${API_HOST} ${API_PORT}
do
    echo "Waiting Recepcionist to be online"
    sleep 1
done

function servicePet() {
    curl ${API_HOST}:${API_PORT}/api/pet/start --data $1 &
    echo ""
}

PET_NAMES=(`cat /usr/share/dict/words | shuf | head -n ${PET_COUNT}`)
for PET_NAME in ${PET_NAMES[@]}
do
    servicePet $PET_NAME
done

docker-compose logs -f | grep -v "rabbitmq_1\|redis_1"