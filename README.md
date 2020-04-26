# WhistleNewsBackend
WhistleNews event-driven RESTful API (CRUD) built with GoLang 

# WhistleNews
WhistleNews is a news portal for whistleblowers. Registered authors would be able to post articles without the need to reveal their identity publically. To empower whistleblowers with the power of data driven decision making, WhistleNews wants a analytics system that can be used by their authors to analyse the popularity of the news that they have written. WhistleNews is against sharing their authors’ as well as readers’ data with any third party analytics tools like MixPanel or GoogleAnalytics. The CEO of WhistleNews does not have any specific citizenship and operates from a secretive location to avoid his website from being taken down by any government.

## Requirements

- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/)

## Install
Clone and setup:

```bash
$ git clone https://github.com/GheisMohammadi/WhistleNewsBackend.git
$ cd WhistleNewsBackend/src/app
$ go build
```
## Running
First, setup docker and go to root folder that contains Dockerfile

```bash
$ docker build --tag WhistleNewsBackend .
```

Then start querying at `http://localhost:3000/counter/v1/`

## Shutdown
```bash
$ docker 
```

## APIs
The entity **Article** has the following fields:

- ID (uint)
- Views (array of objects)

Follows the list of article APIs:

|METHOD|URL|REQUEST HEADERS|REQUEST PAYLOAD|RESPONSE HEADERS|RESPONSE PAYLOAD|
|------|---|---------------|---------------|----------------|----------------|
|GET|http://localhost:3000/counter/v1/statistics/ | | | |Article Object|
|POST|http://localhost:3000/counter/v1/statistics/article_id/ |Content-Type: "application/json"|Article Views||Article Views|
