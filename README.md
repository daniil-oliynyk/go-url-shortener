# Another URL Shortener
Another URL shortener. Written in Go as part of language practice/learning process. Also learned and implemented repository design pattern for storage part.

## Setup
1. Ensure Docker is installed
2. ```docker build -t shortener-service``` to build the Go url shortener service
3.  ``` docker compose up ``` will start up the redis cache and shortener service
4.  You can now send requests to ```localhost:8080```

## API
### POST /api/v1/create-short-url?long_url={longURL}&user_id={id}
Response:
```
201 Created
{
    "Msg": "short url created",
    "ShortUrl": "http://www.localhost:8080/K8v4XzuR"
}
```

###  GET /api/{shortURL}
Response:
```
{"LongUrl":"https://www.google.com/"}
<a href="https://www.google.com/">Found</a>.
```
Or if you put the shortUrl returned from the POST request into your browser youll be redirected accordingly.

## TODO:
Possibly some stuff to add later on if bored
* Add environment variable support (pretty straightforward)
* Add a DB to move stale redis entries
* Redis Sentinel 
