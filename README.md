# gpsql
gpsql is a personal Golang tool that was created to solve the migration stuff when handling with postgres.


## Installing
this package is built by using the `go build ./cmd/main.go` and then `mv /usr/bin` **for now**.

```bash
mdkir gpsql && cd gpsql
git clone https://github.com/praveenmahasena647/goHelpPsql.git .
make build
```
# Using

After cloning the repo by seeing the examples in the `./gpsql.yaml` you can enter your own config details there in your current project folder and
make sure you provide the relative path to up and down files SQL files in that yaml file.

for more examples there are folers provided in ./cmd/up-files/ && ./cmd/down-files/

```bash
gpsql up
```
```bash
gpsql down
```
