# Redis Subscription Test

Using redis pub-sub, this Go code subscribes to two local channels: `nuts` and `butts`.

![squirrel girl](assets/squirrel-girl.gif)


## Usage

In one terminal, run `go run main.go`. Then, in another, fire up your `redis-cli` and publish some
messages:

```
PUBLISH nuts "are so good"
PUBLISH butts "need to be kicked"
PUBLISH nuts "a squirrel favorite"
```

You should see output in the first terminal

```
Message received on 'nuts': are so good
Message received on 'butts': need to be kicked
Message received on 'nuts': a squirrel favorite
```

## Reconnecting

Per [the Redis PubSub docs](https://godoc.org/github.com/go-redis/redis#PubSub), reconnection should be automatic.
