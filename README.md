# golang-repo

gvm setup
- Installation
```
curl -sSL https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash
source ~/.gvm/scripts/gvm
```

- Project setup
```
gvm install go1.16
gvm use go1.16 --default
```

basic cmd

- Run a go program
```
go run exercise001.go
```

- Test a go program :- note: the program files should be included along with the test files.
```
go test exercise001.go exercise001_test.go      
```