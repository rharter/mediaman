package transcoder

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

const (
	TS_STATE_INIT    int = 0
	TS_STATE_RUNNING int = 1
	TS_STATE_EOS     int = 2
	TS_STATE_FAILED  int = 3
)

type TranscodeSession struct {
	ID         string
	InputFile  string
	OutputDir  string
	OutputFile string
	Pipe       *os.File
	Proc       *os.Process

	state int
}

func NewTranscoderSession(id int64, outputDir string, inputFile string) *TranscodeSession {
	t := TranscodeSession{
		InputFile: inputFile,
		OutputDir: outputDir,
	}

	// stringify ID
	t.ID = strconv.FormatInt(id, 10)

	// make sure Close is called
	runtime.SetFinalizer(&t, (*TranscodeSession).Close)
	return &t
}

func (t *TranscodeSession) setState(state int) {
	// EOS and FAILED are final
	if t.state == TS_STATE_EOS || t.state == TS_STATE_FAILED {
		return
	}

	// set state
	t.state = state
}

func (t *TranscodeSession) Open() error {

	if t.IsOpen() {
		return nil
	}

	if err := t.createOutputDirectories(); err != nil {
		log.Printf("Failed to create output directories: %s", err)
		return err
	}

	// create pipe
	pr, pw, err := os.Pipe()
	if err != nil {
		t.setState(TS_STATE_FAILED)
		return err
	}
	t.Pipe = pw

	// Start the transcode process
	attr := os.ProcAttr{}
	attr.Dir = filepath.Join(t.OutputDir, t.ID)
	attr.Files = []*os.File{pr}

	t.OutputFile = filepath.Join(t.OutputDir, t.ID, "index.m3u8")
	t.Proc, err = os.StartProcess("/usr/local/bin/ffmpeg", strings.Fields(t.buildTranscodeCommand()), &attr)
	if err != nil {
		log.Printf("Error starting process: %s", err)
		t.setState(TS_STATE_FAILED)
		pr.Close()
		pw.Close()
		t.Pipe = nil
		return err
	}

	// close the read-end fo the pipe after successful start
	pr.Close()

	// set state
	t.setState(TS_STATE_RUNNING)
	return nil
}

func (t *TranscodeSession) IsOpen() bool {
	return t.state == TS_STATE_RUNNING
}

func (t *TranscodeSession) Close() error {

	t.setState(TS_STATE_EOS)

	t.Pipe.Close()

	// gracefully shut down transcode process
	if err := t.Proc.Signal(syscall.SIGINT); err != nil {
		log.Printf("Sending signal to transcoder failed: %s", err)
		// assume the transcoder process has finished
	}

	log.Printf("Waiting for transcoder to shutdown")
	state, err := t.Proc.Wait()
	if err != nil {
		log.Printf("Transcoder exited with error: %s and state %s", err, state)
		return nil
	}

	log.Printf("Transcoder exit state is %s", state.String())

	return nil

}
