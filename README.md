
# powerfactors go assignment
## Build
go build -o powerfactors

## Help
./powerfactors --help

    -address string
            the address of the http httpServer (default "127.0.0.1")
    -endpoint string
            the exposed endpoint (default "/ptlist")
    -periods string
            the supported time periods (default "1h,1d,1mo,1y")
    -port int
            the port the http httpServer listens on (default 8080)

## Run

### Run with defaults
./powerfactors 

### Run with custom options
./powerfactors --address 127.0.0.2 --port 8000 --endpoint /periodlist --periods 1y,2mo,3d,4h

## Comments
- Time periods beyond 1h,1d,1mo,1y are supported. Examples 3h, 2d etc

- Based on the following requirement
>  The invocation timestamp should be at the start of the  period (e.g. for 1h period a matching timestamp is considered the 20210729T010000Z).

It was my understanding that for a `1d` period the matching timestamp would be like `20210729T000000Z`.
For a `1mo` period like `20210701T000000Z` and so on. However, the given examples did not follow this logic. This implementation follows the logic provided in the examples.
