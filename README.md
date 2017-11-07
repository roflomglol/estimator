# Estimator

Estimator is a service written on Golang for getting time and distance data for a route between two points on a map. It uses gRPC for communications with clients.

## Getting Started

To start this service localy you can use docker compose but read Prerequisites section before you do.

```docker-compose up```

### Prerequisites

Estimator uses OSM and Google API for map info. For optimized perfomance you should have your own OSM server up and running. Check below on tips how to do that.

#### OSM(Docker)
Use this image to build OSM container https://hub.docker.com/r/osrm/osrm-backend/
Download map data, for example http://download.geofabrik.de/russia.html
Unpack data, like so:
```
docker run -t -v $(pwd)/../osm/data:/data osrm/osrm-backend osrm-extract -p /opt/car.lua /data/central-fed-district-latest.osm.pbf
docker run -t -v $(pwd)/../osm/data:/data osrm/osrm-backend osrm-partition /data/central-fed-district-latest.osm.pbf
docker run -t -v $(pwd)/../osm/data:/data osrm/osrm-backend osrm-customize /data/central-fed-district-latest.osm.pbf
```

### Installing

If you have OSM server then all you need to do to start this servis is this:

```
docker-compose build
docker-compose up
```

This should be enough to get the server started.

### Notes

Для реализации данного микросервиса выбрал Golang, т.к. он отлично подходит для таких задач. Для общения с клиентами выбрал gRPC и protobuf. protobuf помогает не тратить время на сериализацию/десериализацию JSON. protobuf не всегда является  очевидным выбором, особенно если уже есть реализация с помощью JSON.

Если сервис OSM не доступен, то приложение попробует запросить данные от Google API(надо заметить, что это займет на порядок больше времени, т.к. запрос не локальный).

Я только недавно стал использовать Golang, так что не все концепции знаю или до конца понимаю. Что-то я делаю так же, как это принято делать в Ruby. Но я быстро учусь. :)
