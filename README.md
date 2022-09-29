# tracking

to use scylladb instead of cassandra
edit the `go.mod` file

```toml
// using scylladb instead of cassandra
replace github.com/gocql/gocql => github.com/scylladb/gocql latest
```
