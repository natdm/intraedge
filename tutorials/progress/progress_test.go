package progress_test

import (
	"bytes"
	"io"
	"log"
	"testing"
	"time"

	"github.com/natdm/intraedge/tutorials/progress"
	requirer "github.com/stretchr/testify/require"
)

func TestProgress(t *testing.T) {
	require := requirer.New(t)
	bs := []byte{1, 2, 3, 4, 3, 5, 6, 2, 1, 5}
	rdr := bytes.NewReader(bs)
	pc, percentCh := progress.New(uint64(len(bs)))
	go func() {
		for p := range percentCh {
			log.Println(p)
		}
	}()
	_, err := io.Copy(pc, rdr)
	require.NoError(err)
	<-time.After(time.Second * 5)
}
