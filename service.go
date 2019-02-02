package hotreload

import (
	"github.com/radovskyb/watcher"
	rr "github.com/spiral/roadrunner/cmd/rr/cmd"
	"github.com/spiral/roadrunner/cmd/util"
	"log"
	"os"
	"path/filepath"
	"time"
)

const ID = "hotreload"

type Service struct {
}

func (s *Service) Init(config *Config) (bool, error) {

	if !config.Enable {
		return false, nil
	}

	util.Printf("Loading Hot Reload: ")

	w := watcher.New()

	w.AddFilterHook(func(info os.FileInfo, fullPath string) error {
		str := info.Name()

		// Match
		if match, err := filepath.Match(config.Files, str); err != nil {
			return err
		} else {
			if match {
				return nil
			}
		}

		// No match.
		return watcher.ErrSkip
	})

	go func() {
		for {
			select {
			case <-w.Event:
				s.reloadWorkers()
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive(config.Path); err != nil {
		log.Fatalln(err)
	}

	go func() {
		if err := w.Start(time.Millisecond * *config.Tick); err != nil {
			log.Fatalln(err)
		}
	}()

	util.Printf("<green+hb>done</reset>\n")

	return true, nil
}

func (s *Service) reloadWorkers() error {
	client, err := util.RPCClient(rr.Container)
	if err != nil {
		return err
	}
	defer client.Close()

	util.Printf("<green>File(s) updated. Resetting workers</reset>: ")

	var r string
	if err := client.Call("http.Reset", true, &r); err != nil {
		util.Printf("<red+hb>error</reset>\n")
		return err
	}

	util.Printf("<green+hb>done</reset>\n")
	return nil
}
