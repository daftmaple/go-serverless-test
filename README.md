# Vercel Serverless Functions in Go

## Serve locally

Make the `test.go` file which contains `main()` that runs an http server in port 8080 by default or any port set with environment variable `APP_PORT`.

```sh
make
# Run the app in port 8090
APP_PORT=8090 ./a.out
```
