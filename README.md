# news-website

Scalable news website on AWS.

## Build

```bash
for b in get post; do env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/${b} cmd/${b}/main.go; done
```

## Deploy

```bash
sls deploy --verbose --aws-profile training --region eu-central-1 --stage dev
```

## Testing

Post:

```bash
curl --request POST \
  --url https://brzqxi4pp2.execute-api.eu-central-1.amazonaws.com/dev/newsitem \
  --header 'Content-Type: application/json' \
  --data '{
  "Title": "Hello",
  "Description": "World"
}'
```

Get:

```bash
curl --request GET \
  --url some-url/dev/news
```

## Sources

* <https://github.com/serverless/examples/tree/v3/aws-golang-rest-api-with-dynamodb>
