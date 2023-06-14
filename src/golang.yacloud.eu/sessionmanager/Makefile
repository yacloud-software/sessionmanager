.PHONY:	client server
all:	client server
client:
	cd client && go install ${LDFLAGS} sessionmanager-client.go `ls -1 *.go|grep -v sessionmanager-client.go`
server:
	cd server && go install ${LDFLAGS} sessionmanager-server.go `ls -1 *.go|grep -v sessionmanager-server.go`
