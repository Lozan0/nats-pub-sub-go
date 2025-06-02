package main

import "time"

func main() {

	ready := make(chan bool)

	go subscribe(ready)

	// Esperamos a que subscribe() esté listo
	<-ready

	go publish()

	// Le damos tiempo para recibir mensajes
	time.Sleep(3 * time.Second)

}
