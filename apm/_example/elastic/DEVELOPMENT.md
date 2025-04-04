# DEVELOPMENT

## Elastic

Run the latest version of the Elastic APM stack with Docker and Docker Compose.

It gives you the ability to analyze any data set by using the searching/aggregation capabilities of Elasticsearch and the visualization power of Kibana.

## Requirements

### Host setup

* Docker Engine version 17.05+
* Docker Compose version 1.12.0+
* ~1.5 GB of RAM

By default, the stack exposes the following ports:

* 5000: Log stash TCP input
* 9200: Elasticsearch HTTP
* 9300: Elasticsearch TCP transport
* 5601: Kibana

### Usage

### Bringing up the stack

Clone this repository onto the Docker host that will run the stack, then start services locally using Docker Compose:

```sh
docker-compose up
```

In case any issue, follow the configuration mentioned [here](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html)

You can also run all services in the background (detached mode) by adding the -d flag to the above command.

> ℹ️ You must run docker-compose build first whenever you switch branch or update a base image.

If you are starting the stack for the very first time, please read the section below attentively.

### Cleanup

Elasticsearch data is persisted inside a volume by default.

In order to entirely shutdown the stack and remove all persisted data, use the following Docker Compose command:

```sh
docker-compose down -v
```

Elasticsearch is available at <http://localhost:8100/>
Kibana UI is up at <http://localhost:5601/>

### Local Demo

```sh
go run main.go
```

* User detail - <http://localhost:8000/user>
* Error - <http://localhost:8000/error>
* Kibana APM - <http://localhost:5601/app/apm>
