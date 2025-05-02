package upload

import (
	"io"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func Upload() {
	config := &ssh.ClientConfig{
		User:            os.Getenv("SFTP_USER"),
		Auth:            []ssh.AuthMethod{ssh.Password(os.Getenv("SFTP_PASSWORD"))},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, _ := ssh.Dial("tcp", os.Getenv("SFTP_HOST"), config)
	client, _ := sftp.NewClient(conn)
	defer client.Close()

	localFile, _ := os.Open("news.mp3")
	defer localFile.Close()

	remoteFile, _ := client.Create("/news.mp3")
	defer remoteFile.Close()

	io.Copy(remoteFile, localFile)
	log.Println("News uploaded succcessfully on Azuracast Server")
}
