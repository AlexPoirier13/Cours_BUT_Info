package main

import (
    "log"
    "net"
    //"time"
    "bufio"
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

    cnx, err := listener.Accept() //bloquant tant que pas de tentative de connexion cote client
    if err != nil {
        log.Println("Erreur lors de l'appel à listener.Accept : ", err)
        return
    }
    defer func(c net.Conn){
        c.Close()
        log.Println("Connexion fermée")
    }(cnx)

    log.Println("Client connecté")



    out := bufio.NewWriter(cnx)

    in := bufio.NewReader(cnx)


    err = send_message(out, "hello\n")
    
    if err != nil {
        log.Println("Erreur lors de l'envoi du message")
        return
    }

    message, err := read_message(in)

    if err != nil {
        log.Println("Erreur lors de la reception du message")
        return
    }

    if message != "start"{
        log.Println("le message recu n'est pas start")
        cnx.Close()
        log.Println("Connexion fermée")
    } else {
        send_message(out, "ok\n")
    }


    message, err = read_message(in)

    if err != nil {
        log.Println("Erreur lors de la reception du message")
        return
    }

    if message != "end"{
        log.Println("le message recu n'est pas end")
        return
    } else {
        send_message(out, "end\n")
    }

    cnx, err = listener.Accept() //bloquant tant que pas de tentative de connexion cote client
    if err != nil {
        log.Println("Erreur lors de l'appel à listener.Accept : ", err)
        return
    }
    defer func(c net.Conn){
        c.Close()
        log.Println("Connexion fermée")
    }(cnx)

    log.Println("Client connecté")

    return
}


func send_message(out *bufio.Writer, message string) (error) {
        _, err := out.WriteString(message)
        if err != nil {
            return err
        }
        err = out.Flush()
        if err != nil {
            return err
        }
        log.Println("Envoi du message :", message[0:len(message)-1])
        return nil 
}

func read_message(out *bufio.Reader) (string, error) {

        message, err := out.ReadString('\n')

        if err != nil {
            return "", err
        }
   
       

        log.Println("Message lu :", message[0:len(message)-1])

        return message[0:len(message)-1], err


}

//sur silverblue nc -U test.sock
// sur ubutu go run serveur.web


