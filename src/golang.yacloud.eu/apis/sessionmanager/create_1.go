// client create: SessionManagerClient
/*
  Created by /home/cnw/devel/go/yatools/src/golang.yacloud.eu/yatools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.yacloud.eu/apis/sessionmanager/sessionmanager.proto
   gopackage : golang.yacloud.eu/apis/sessionmanager
   importname: ai_0
   clientfunc: GetSessionManager
   serverfunc: NewSessionManager
   lookupfunc: SessionManagerLookupID
   varname   : client_SessionManagerClient_0
   clientname: SessionManagerClient
   servername: SessionManagerServer
   gsvcname  : sessionmanager.SessionManager
   lockname  : lock_SessionManagerClient_0
   activename: active_SessionManagerClient_0
*/

package sessionmanager

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_SessionManagerClient_0 sync.Mutex
  client_SessionManagerClient_0 SessionManagerClient
)

func GetSessionManagerClient() SessionManagerClient { 
    if client_SessionManagerClient_0 != nil {
        return client_SessionManagerClient_0
    }

    lock_SessionManagerClient_0.Lock() 
    if client_SessionManagerClient_0 != nil {
       lock_SessionManagerClient_0.Unlock()
       return client_SessionManagerClient_0
    }

    client_SessionManagerClient_0 = NewSessionManagerClient(client.Connect(SessionManagerLookupID()))
    lock_SessionManagerClient_0.Unlock()
    return client_SessionManagerClient_0
}

func SessionManagerLookupID() string { return "sessionmanager.SessionManager" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.

func init() {
   client.RegisterDependency("sessionmanager.SessionManager")
   AddService("sessionmanager.SessionManager")
}

