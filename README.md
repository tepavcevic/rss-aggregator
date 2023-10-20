# Go API with postgres, chi, goose and sqlc.

It is a [RSS](https://en.wikipedia.org/wiki/RSS) feed aggregator in Go! It's a web server that allows clients to:

- Add RSS feeds to be collection
- Follow and unfollow RSS feeds that other users have added
- Fetch all of the latest posts from the RSS feeds they follow

Project has a form of flat architecture since it was one of the first bigger go projects for me, but still small enough to be easy to follow.
Because of that it mostly follows "happy path", so it could be improved in many ways.

Credits: [Boot.Dev](https://github.com/bootdotdev)

### Setup .env file

```js
PORT=8080
DB_URL= <YOUR_DB_URL>
```

### Setup DB

Migrate DB: Run following command inside `sql/schema` directory

```shell
goose postgres <YOUR_DB_URL> up
```

Generate Queries

```shell
sqlc generate
```

### Build & Start the Server

```shell
go build && ./rss-aggregator
```
