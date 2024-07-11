### echo-server

A very simple echo server which can be used for local-testing or webhook related testing
It can return the following information at the moment

```go
type RequestInfo struct {
	Method      string              `json:"method"`
	URL         string              `json:"url"`
	Headers     map[string][]string `json:"headers"`
	QueryParams map[string][]string `json:"query_params"`
	Body        string              `json:"body"`
}
```

### Usage

1. Build the image

    ```bash
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o echo_server .

    ECHO_SERVER_PORT=18080 go run main.go
    ```

2. Build the container

    ```bash
    docker build -t echo-server:latest -f Dockerfile .
    docker run --rm -it -e ECHO_SERVER_PORT=18080 echo-server:latest
    ```


3. Test with curl or browser or your favorite http client
   

   ```bash
   ❯ curl -s http://localhost:18080 | jq -r
    {
    "method": "GET",
    "url": "/",
    "headers": {
        "Accept": [
        "*/*"
        ],
        "User-Agent": [
        "curl/8.6.0"
        ]
    },
    "query_params": {},
    "body": ""
    }
    ```