package main 

import (
    "log"
    "net"
)

func main() {
    cnx, err := net.Dial("unix", "test.sock")
    if err != nil {
        log.Println("Erreur lors de l'appel à net.Dial : ", err)
        return
    }
    defer func(c net.Conn){
        c.Close()
        log.Println("Connexion fermée")
    }(cnx)

    log.Println("Connexion ouverte")

    return
}
