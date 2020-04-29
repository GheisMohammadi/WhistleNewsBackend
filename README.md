# WhistleNewsBackend
WhistleNews event-driven RESTful API (CRUD) built with GoLang 

# WhistleNews
WhistleNews is a news portal for whistleblowers. Registered authors would be able to post articles without the need to reveal their identity publically. To empower whistleblowers with the power of data driven decision making, WhistleNews wants a analytics system that can be used by their authors to analyse the popularity of the news that they have written. WhistleNews is against sharing their authors’ as well as readers’ data with any third party analytics tools like MixPanel or GoogleAnalytics. The CEO of WhistleNews does not have any specific citizenship and operates from a secretive location to avoid his website from being taken down by any government.

## Requirements

- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/)

## Install
Clone and setup:

```console
$ git clone https://github.com/GheisMohammadi/WhistleNewsBackend.git
$ cd WhistleNewsBackend/src/app
$ go build
```
## Running

1- install mongodb

you can run it using docker by this command:

    ```console
    $ sudo docker run -d -p 27017-27019:27017-27019 --name mongodb mongo:4.0.4
    ```

2- create db name "WhistleNews" and create collection with name "articles"

3- install NSQ and run services

    - download proper binary package from [NSQ Binary Packages](https://nsq.io/deployment/installing.html)

    - run a small NSQ cluster on your local machine by following [NSQ Installtion](https://nsq.io/overview/quick_start.html) and make sure it works correctly!

4- go to root folder that contains Dockerfile

```console
$ sudo docker build --tag "whistlebackend" .
$ sudo docker run --network="host" "whistlebackend"
```

Then start querying at `http://localhost:3085/counter/v1/`

## Tests

For do testing, you can use the following command:

```console
$ go test -v ./...
```

## Test coverage

for testing coverage of the tests, you can use following commands:

```console
$ go test ./... -coverprofile cover.out
$ go tool cover -html=cover.out -o coverage.html
```

and to see total coverage:
```console
$ go tool cover -func cover.out | grep total
```

the output will be like this:

```console
total:                                  (statements)            76.7%
```

## APIs
The entity **Article** has the following fields:

- ID (uint)
- Views (array of objects)

Follows the list of article APIs:

|METHOD|URL|REQUEST HEADERS|REQUEST PAYLOAD|RESPONSE HEADERS|RESPONSE PAYLOAD|
|------|---|---------------|---------------|----------------|----------------|
|POST|http://localhost:3085/counter/v1/statistics/ |Content-Type: "application/json"|Article ID| |Article Object|
|POST|http://localhost:3085/counter/v1/article/add |Content-Type: "application/json"|Article ID| |Article Object|
|GET |http://localhost:3085/counter/v1/statistics/article_id/{id} ||||Article Views|
