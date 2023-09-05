package main

import (
	"io"
	"log"
	"os"
	"pack.ag/tftp"
)

func main() {
	s, err := tftp.NewServer(":69", tftp.ServerSinglePort(true))
	if err != nil {
		panic(err)
	}
	readHandler := tftp.ReadHandlerFunc(proxyTFTP)
	writeHandler := tftp.WriteHandlerFunc(ReceiveTFTP)
	s.ReadHandler(readHandler)
	s.WriteHandler(writeHandler)
	s.ListenAndServe()
	select {}

}

func proxyTFTP(w tftp.ReadRequest) {
	log.Printf("[%s] GET %s\n", w.Addr().IP.String(), w.Name())
	file, err := os.Open("/tftpboot/" + w.Name()) // For read access.
	if err != nil {
		log.Println(err)
		w.WriteError(tftp.ErrCodeFileNotFound, err.Error())
		return
	}
	defer file.Close()

	if _, err := io.Copy(w, file); err != nil {
		log.Println(err)
	}
}

func ReceiveTFTP(w tftp.WriteRequest) {
	log.Printf("[%s] PUT %s\n", w.Addr().IP.String(), w.Name())
	data, err := io.ReadAll(w)
	if err != nil {
		log.Println(err)
		return
	}

	if err := os.WriteFile("/tftpboot/"+w.Name(), data, 0644); err != nil {
		log.Println(err)
		return
	}
}
