package main 

import (
    "log"
    "net"
    "bufio"
    //"time"
    
)

func main() {
    cnx, err := net.Dial("unix", "test.sock") //demande de connexion qui sera valider par l'accept du serveur, bloquant aussi
    if err != nil {
        log.Println("Erreur lors de l'appel à net.Dial : ", err)
        return
    }
    defer func(c net.Conn){
        c.Close()
        log.Println("Connexion fermée")
    }(cnx)

    log.Println("Connexion ouverte")

    

    out := bufio.NewWriter(cnx)
    in := bufio.NewReader(cnx)


    message, err := read_message(in)

    if err != nil {
        log.Println("Erreur message non recu")
        return
    }

    if message != "hello"{
        log.Println("Le message recu n'est pas hello c'est ", message)
        return
    } else {
        send_message(out, "start\n")
    }

    message, err = read_message(in)

    if err != nil {
        log.Println("Erreur message non recu ou autre probleme")
        return
    }

    if message != "ok"{
        log.Println("message recu n'est pas ok")
        cnx.Close()
        log.Println("Connexion fermée")
    } else {
        send_message(out, "end\n")
    }


    message, err = read_message(in)

    if err != nil {
        log.Println("Erreur message non recu")
        return
    }

    if message != "end"{
        log.Println("message recu n'est pas end")
        return
    } else {
        read_message(in)
    }


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
        log.Println("Envoi du message :", message[0:len(message)])
        return nil 
}

func read_message(in *bufio.Reader) (string, error) {

        message, err := in.ReadString('\n')

        if err != nil {
            return "", err
        }
   
       

        log.Println("Message lu :", message[0:len(message)-1])

        return message[0:len(message)-1], nil

}
