
# build: libfoo-android.so libfoo-linux-amd64.so libfoo-linux-i386.so libfoo-windows-amd64.dll libfoo-windows-i386.dll libfoo-darwin-amd64.so libfoo-darwin-i386.so libfoo-linux-arm.so
build: libfoo-linux-amd64.so

libfoo-linux-amd64.so: *.go
	GOOS=linux   GOARCH=amd64 go build -buildmode=c-shared -o libfoo-linux-amd64.so       github.com/1681-dodo-bird/libfoo 

libfoo-linux-i386.so: *.go
	GOOS=linux   GOARCH=386   go build -buildmode=c-shared -o libfoo-linux-i386.so        github.com/1681-dodo-bird/libfoo 

libfoo-windows-amd64.dll: *.go
	GOOS=windows GOARCH=amd64 go build -buildmode=c-shared -o libfoo-windows-amd64.dll    github.com/1681-dodo-bird/libfoo 

libfoo-windows-i386.dll: *.go
	GOOS=windows GOARCH=386   go build -buildmode=c-shared -o libfoo-windows-i386.dll     github.com/1681-dodo-bird/libfoo 

libfoo-darwin-amd64.so: *.go
	GOOS=darwin  GOARCH=amd64 go build -buildmode=c-shared -o libfoo-darwin-amd64.so      github.com/1681-dodo-bird/libfoo 

libfoo-darwin-i386.so: *.go
	GOOS=darwin  GOARCH=386   go build -buildmode=c-shared -o libfoo-darwin-i386.so       github.com/1681-dodo-bird/libfoo 

libfoo-linux-arm.so: *.go
	GOOS=linux   GOARCH=arm   go build -buildmode=c-shared -o libfoo-linux-arm.so         github.com/1681-dodo-bird/libfoo 

libfoo-android.so: *.go
	GOOS=android              go build -buildmode=c-shared -o libfoo-android.so           github.com/1681-dodo-bird/libfoo 

*.go:
