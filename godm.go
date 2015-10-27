package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gopkg.in/qml.v1"
)

var Receiver = make(chan int64)

type Downloader struct {
	URLField    qml.Object
	Download    qml.Object
	Pause       qml.Object
	Resume      qml.Object
	Cancel      qml.Object
	ProgressBar qml.Object
	Percent     qml.Object
}

type Pipe struct {
	io.ReadCloser
	total int64 // Total no. of bytes transferred
}

func (pt *Pipe) Read(p []byte) (int, error) {

	n, err := pt.ReadCloser.Read(p)
	pt.total += int64(n)
	Receiver <- pt.total

	return n, err
}

func (d *Downloader) StartDownload(urlField, progress, percent, download, pause, resume, cancel qml.Object) {

	d.ProgressBar = progress
	d.Percent = percent
	d.Download = download
	d.Pause = pause
	d.Resume = resume
	d.Cancel = cancel

	d.ProgressBar.Set("value", 0)
	d.Percent.Set("text", "")

	url := urlField.String("text")

	if url == "" {
		fmt.Println("Please specify a url to download a file from")
		return
	}

	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("Error While Getting file from URL: ", err)
		return
	}

	destination, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error While Creating File: ", err)
		return
	}

	fmt.Println("Download Started for : ", url)

	contentLength := resp.ContentLength

	resp.Body = &Pipe{ReadCloser: resp.Body}

	d.ButtonClicked(0, 0, 0, 0)

	go d.Copier(destination, resp, int(contentLength))
}

func (d *Downloader) Copier(destination *os.File, resp *http.Response, contentLength int) {

	defer resp.Body.Close()
	defer destination.Close()

	go d.Progresser(contentLength)

	noOfBytesWritten, err := io.Copy(destination, resp.Body)
	if err != nil {
		fmt.Println("Error While Downloading file from URL: ", err)
		os.Exit(1)
	}

	fmt.Println("Transferred", noOfBytesWritten, "bytes")
}

func (d *Downloader) Progresser(contentLength int) {

	for {
		select {
		case downloaded := <-Receiver:
			percent := int(float64(downloaded) / (float64(contentLength) / float64(100)))
			d.Percent.Set("text", strconv.Itoa(percent)+"%")
			d.ProgressBar.Set("value", float64(downloaded)/float64(contentLength))
			if downloaded >= int64(contentLength) {
				fmt.Println("Done...")
				d.ButtonClicked(1, 0, 0, 0)
				return
			}
		}
	}
}

func (d *Downloader) ButtonClicked(download, pause, resume, cancel int) {
	d.Download.Set("enabled", download)
	d.Pause.Set("enabled", pause)
	d.Resume.Set("enabled", resume)
	d.Cancel.Set("enabled", cancel)
}

func main() {
	if err := qml.Run(godmUI); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func godmUI() error {
	engine := qml.NewEngine()

	controls, err := engine.LoadFile("main.qml")
	if err != nil {
		fmt.Println("Error While Loading QML File", err)
	}

	downloader := Downloader{}

	context := engine.Context()
	context.SetVar("download", &downloader)

	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()
	return err
}
