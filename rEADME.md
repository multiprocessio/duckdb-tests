What's failing:

```
CGO_LDFLAGS="-L$(pwd)/duckdb/src/include" CGO_CFLAGS="-I$(pwd)/duckdb/src/include" go build -ldflags '-extldflags " -lstdc++ -lm -lduckdb -static"'
# github.com/marcboeker/go-duckdb
../../go/pkg/mod/github.com/marcboeker/go-duckdb@v0.0.0-20220427142532-cd9f33e64d9a/connection.go:4:10: fatal error: duckdb.h: No such file or directory
    4 | #include <duckdb.h>
      |          ^~~~~~~~~~
compilation terminated.
```
