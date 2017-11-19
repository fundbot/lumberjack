# lumberjack
collects wood from numerous places and brings it to the townhall


## Setting up postgres for development
docker run --name postgresql -itd --restart always --publish 5432:5432 --volume ~/docker/postgresql:/var/lib/postgresql --env 'DB_USER=dbuser' --env 'DB_PASS=dbuserpass' sameersbn/postgresql:9.6-2
docker exec -it postgresql sudo -u postgres psql

## Setting up redis for development
docker run --name redis -d --restart=always --publish 6379:6379 --volume ~/docker/redis:/var/lib/redis sameersbn/redis:latest --logfile /var/log/redis/redis-server.log
docker exec -it redis tail -f /var/log/redis/redis-server.log
