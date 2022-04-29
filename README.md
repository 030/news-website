# news-website

Scalable news website on AWS.

## Getting started

```bash
aws s3 mb s3://testi324332524523432 --profile=training --region=eu-central-1
zip bla.zip -r cmd go.mod go.sum buildspec.yml
aws s3 cp bla.zip s3://testi324332524523432/bla.zip --profile=training --region=eu-central-1
```

## Deploy

```bash
sls deploy --verbose --aws-profile training --region eu-central-1 --stage dev
```

## Testing

Post:

```bash
curl --request POST \
  --url https://f2u5ilc7sj.execute-api.eu-central-1.amazonaws.com/dev/newsitem \
  --header 'Content-Type: application/json' \
  --data '{
  "Title": "Hello",
  "Description": "World"
}'
```

Get:

```bash
curl --request GET \
  --url https://f2u5ilc7sj.execute-api.eu-central-1.amazonaws.com/dev/news
```

## Cleanup

Deletion of the CloudFormation stack does not seem to remove the DynamoDB.
Remove it manually to save costs.

## Sources

* <https://github.com/serverless/examples/tree/v3/aws-golang-rest-api-with-dynamodb>
