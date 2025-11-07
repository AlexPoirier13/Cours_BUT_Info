/*
package main

import (
    "log"
    "net"
    "time"
)

func main ( ) {
    listener, err := net.Listen("unix", "test.sock")
    if err != nil {
        log.Println("Erreur lors de l'appel à net.Listen : ", err)
        return
    }
    defer func(l net.Listener) {
        l.Close()
        log.Println("Connexion fermée")
    }(listener)
    time.Sleep (10 * time.Second)
}


*/


package main

import (
    "log"
    "net"
    "time"
)

func main ( ) {
    listener, err := net.Listen("unix", "test.sock")
    if err != nil {
        log.Println("Erreur lors de l'appel à net.Listen : ", err)
        return
    }
    defer func(l net.Listener){
        l.Close()
        log.Println("Adresse d'écoute libérée")
    }(listener)

    cnx, err := listener.Accept()
    if err != nil {
        log.Println("Erreur lors de l'appel à listener.Accept : ", err)
        return
    }
    defer func(c net.Conn){
        c.Close()
        log.Println("Connexion fermée")
    }(cnx)

    log.Println("Client connecté")

    time.Sleep (10 * time.Second)

    return
}



//sur silverblue nc -U test.sock
// sur ubutu go run serveur.web


