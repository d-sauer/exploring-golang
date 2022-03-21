# Hello Lambda


## Local Docker Lambda
https://docs.aws.amazon.com/lambda/latest/dg/go-image.html

Running Lambda function locally in Docker container with RIE (aws-lambda-runtime-interface-emulator)

1. Build Docker container `make build-docker-lambda`
2. Start container `docker run -p 9090:8080 hello-lambda /main Handler`
    - attribute `/main` name of the binary
    - attribute `Handler` name of Lambda handler function
3. Trigger lambda function with `curl`
```bash
curl -v -XPOST "http://localhost:9090/2015-03-31/functions/function/invocations" -d '{"id": 1234,"value": "value1"}'
```

