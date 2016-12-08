package localhttpserver

import (
    "net"
    "net/http"
    "strconv"
)

func testPort(port int) bool {
    portS := strconv.Itoa(port)
    listenNet, err := net.Listen("tcp", ":" + portS)
    if nil != err {
        return false
    }
    err = listenNet.Close()
    if nil != err {
        return false
    }
    return true
}

func getPort(port int) string {
    for !testPort(port) {
        port += 1
    }
    return strconv.Itoa(port)
}

func main() {
    http.Handle("/", http.FileServer(http.Dir(".")))
    http.ListenAndServe(":"+getPort(8000), nil)
}
