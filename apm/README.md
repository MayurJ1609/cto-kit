# APM

Go APM allows you to monitor your Go applications currently supported with Elastic APM and a easy of extending to other service. It helps you track transactions, outbound requests, database calls, and other parts of your. Go application's behavior and provides a running overview of garbage collection, goroutine activity, and memory use.

## Installation

To install run the following command

```sh
go get github.com/skortech/st-kit/apm
```

## Elastic

Elastic [APM](https://www.elastic.co/guide/en/apm/agent/go/current/configuration.html#config-service-node-name) is an application performance monitoring system built on the Elastic Stack.

### Required Below environment

- ELASTIC_APM_SERVICE_NAME
- ELASTIC_APM_SERVER_URL AND ELASTIC_APM_SERVER_URLS
- ELASTIC_APM_SECRET_TOKEN: If your server requires a secret token for authentication, you must also set APM_TOKEN.

```go
    apm, err := New()
    if err != nil {
        //Handler error
    }
    txn, _ := apm.StartWebTransaction("/users", nil, httptest.NewRequest(http.MethodGet, "/", nil))
    apm.AddAttribute(txn, "user_id", "SL06202201")
    defer apm.EndTransaction(txn, nil)
    dataSegment, _ := apm.StartDataStoreSegment(txn, "Postgres", "FIND", "SL_USER")
    apm.EndSegment(dataSegment)
    segment, _ := apm.StartSegment(txn, "opt-service")
    apm.EndSegment(segment)
    externalSegment, _ := apm.StartExternalSegment(txn, "https://services.skorlife.com/")
    apm.EndExternalSegment(externalSegment)
```

## TODO

- [ ] Need to do the step by step documentation of APM integration
