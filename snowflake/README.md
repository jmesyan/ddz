Snowflake
===

twitter snowflake UUID generator in golang  

copy from https://github.com/gonet2/snowflake

### OSX

Docker for OSX doesn't support `--net=host`, yet.

change the ip address in etcdclient/client.go to mac's ip address

*for example:*  
```go
const (
	DEFAULT_ETCD = "http://192.168.0.10:2379"
)
```

you can acquire ip address by 
```bash
$ ifconfig en0 | grep "inet " | cut -d " " -f2
```
