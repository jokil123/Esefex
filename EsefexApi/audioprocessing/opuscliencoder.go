package audioprocessing

import (
	"io"
	"log"
	"os/exec"
)

type OpusCliEncoder struct {
	source *io.Reader
	cmd    *exec.Cmd
	stdin  *io.WriteCloser
	stdout *io.ReadCloser
}

func NewOpusCliEncoder(s io.Reader) (*OpusCliEncoder, error) {
	cmd := exec.Command("opusenc", "--raw", "--raw-bits=16", "--raw-rate=48000", "--raw-chan=2", "--quiet", "-", "-")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return &OpusCliEncoder{
		source: &s,
		cmd:    cmd,
		stdin:  &stdin,
		stdout: &stdout,
	}, nil
}

// Returns the next encoded opus frame (20ms)
// this function is called 50 times per second
// and therefore needs to be fast
func (e *OpusCliEncoder) EncodeNext() ([]byte, error) {
	buf := make([]byte, 960*2*2)

	// Read from the source
	log.Println("Reading from source")
	n, err := (*e.source).Read(buf)
	if err != nil {
		log.Println("Error reading from source")
		return nil, err
	}

	//Write to the encoder
	log.Println("Writing to encoder stdin")
	n, err = (*e.stdin).Write(buf[:n])
	if err != nil {
		log.Println("Error writing to encoder stdin")
		log.Println(n)
		return nil, err
	}

	bytes := make([]byte, 960*2*2)

	// Read from the encoder
	log.Println("Reading from encoder stdout")
	n, err = (*e.stdout).Read(bytes)
	if err != nil {
		log.Println("Error reading from encoder stdout")
		return nil, err
	}

	// Read from the encoder
	// log.Println("Reading from encoder stdout")
	// bytes, err := io.ReadAll(*e.stdout)
	// if err != nil {
	// 	log.Println("Error reading from encoder stdout")
	// 	return nil, err
	// }

	// return bytes, nil
	return bytes[:n], nil
}

func (e *OpusCliEncoder) Close() error {
	(*e.stdin).Close()
	(*e.stdout).Close()
	return e.cmd.Wait()
}