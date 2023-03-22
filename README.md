# powerfactors_go_assignment

go build -o powerfactors

help
./powerfactors --help
-address string
        the address of the http httpServer (default "127.0.0.1")
  -endpoint string
        the exposed endpoint (default "/ptlist")
  -periods string
        the supported time periods (default "1h,1d,1mo,1y")
  -port int
        the port the http httpServer listens on (default 8080)
        
run with defaults
./powerfactors 

run with custom options
./powerfactors --address 127.0.0.2 --port 8000 --endpoint /periodlist --periods 1y,2mo,3d,4h
