Kantine bot
-----------

An experiment to build a bot that would retrieve a menu of nearby eating places.


### Building

```bash
export PATH=/usr/local/go/bin:$PATH
go mod vendor
make build -j4
```

### Running

1. Get a token
```bash
bin/new_auth-linux-amd64
```

2. Get your slack api key

3. Run the bot

```bash
AUTHORIZATION=YOUR_AUTHORIZATION SLACK_KEY=YOUR_SLACK_KEY bin/kantinebot-linux-amd64
```