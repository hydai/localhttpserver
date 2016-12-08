package main

import (
    "net"
    "net/http"
    "strconv"
    "fmt"
)

func testPort(port string) bool {
    listenNet, err := net.Listen("tcp", ":" + port)
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
    portS := strconv.Itoa(port)
    for !testPort(portS) {
        port += 1
        portS = strconv.Itoa(port)
    }

    fmt.Printf("Listen on http://127.0.0.1:" + portS + "\n")
    return portS
}

func main() {
    http.Handle("/", http.FileServer(http.Dir(".")))
    http.ListenAndServe(":"+getPort(8000), nil)
}
