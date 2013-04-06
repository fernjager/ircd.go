package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "net"
    "time"
)


const homePage = `<!DOCTYPE html>
<html>
<head>
    <script src="//ajax.googleapis.com/ajax/libs/jquery/1.8.3/jquery.min.js"></script>
</head>
<body>
    <form action="/target" id="postToGoHandler">
    <input type="submit" value="Post" />
    </form>
    <div id="result"></div>
<script>
$("#postToGoHandler").submit(function(event) {
    event.preventDefault();
    $.post("/target", JSON.stringify({"Param1": "Value1"}),
        function(data) {
            $("#result").empty().append(data);
        }
    );
});
</script>
</body>
</html>`

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "{\"leads\":\"1\",\"success\":true}")
}


func target(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    if body, err := ioutil.ReadAll(r.Body); err != nil {
        fmt.Fprintf(w, "Couldn't read request body: %s", err)
    } else {
        dec := json.NewDecoder(strings.NewReader(string(body)))
        var m Message
        if err := dec.Decode(&m); err != nil {
            fmt.Fprintf(w, "Couldn't decode JSON: %s", err)
        } else {
            
        }
    }
}

func handleConnection( c net.Conn ){
        user := InitUser(c)
        
        daytime := time.Now().String()
        msg := user.Receive()
        //&Message{ "PRIVMSG", daytime }
        user.Send( msg )

        //user.Disconnect();
}

func startWeb(){
    http.HandleFunc("/", home)
    http.HandleFunc("/target", target)

    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }

}

func startLinks(){

}
func startIRCd(){
    ln, err := net.Listen( "tcp", ":6667" )
    if err != nil {
        // handle error
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            // handle error
            continue
        }
        go handleConnection(conn)
    }
}

func main() {
    go startWeb();
    go startLinks();
    startIRCd();
}



