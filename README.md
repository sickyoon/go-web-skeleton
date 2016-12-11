# go-web-skeleton

Lightweight Go Web Skeleton with focus on performance.
Use this as a starting point for your new Go project!

# Dependencies

* github.com/julienschmidt/httprouter
* github.com/golang/groupcache
* github.com/gorilla/handlers
* gopkg.in/mgo.v2

# Mongo Deployment through Docker
```bash
docker pull mongo
docker run --name some-mongo -d mongo
docker run -it --link some-mongo:mongo --rm mongo sh -c 'exec mongo "$MONGO_PORT_27017_TCP_ADDR:$MONGO_PORT_27017_TCP_PORT/test"'
```

# License

MIT

