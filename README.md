# news-website

Scalable news website on AWS

* Repository - github
* Build - AWS codeBuild

## Executive summary

This news-website has been deployed on Amazon Web Services (AWS) using Lambda,
DynamoDB and API Gateway. The reason for choosing these technologies is that
the developers could completely focus on creating business value instead of
managing build infrastructure.

## Technical documentation

### Build

```
go build && zip -j news-website.zip news-website
```

and upload the zip to AWS Lambda.

### Test

Check the CloudWatch logging. Several debug statements have been added.

#### GET

```
aws apigateway test-invoke-method --rest-api-id REST_API_ID --resource-id RESOURCE_ID --http-method "GET" --path-with-query-string "/news?title=some-title"
```

#### POST

```
aws apigateway test-invoke-method --rest-api-id REST_API_ID --resource-id RESOURCE_ID --http-method "POST" --path-with-query-string "/news" --body file://tests/news1.json
```