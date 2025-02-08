## Appointment System ##

https://github.com/golang/mock
go install github.com/golang/mock/mockgen@v1.6.0

ls -l ~/go/bin

vi ~/.bashrc

export PATH=$PATH:~/go/bin

mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/jpmoraess/appointment-api/db/sqlc Store
