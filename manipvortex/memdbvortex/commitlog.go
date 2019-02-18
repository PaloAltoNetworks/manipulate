package memdbvortex

import (
	"context"
	"encoding/json"
	"os"
)

// newLogWriter creates a new log for transactions. At this point there is no
// file rotation. If it runs out of storage it will die.
func newLogWriter(ctx context.Context, filename string, size int) (chan *Transaction, error) {

	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	events := make(chan *Transaction, size)

	go func(f *os.File) {
		defer f.Close()

		for {

			select {

			case d := <-events:
				data, err := json.Marshal(d)
				if err != nil {
					continue
				}

				_, err = f.Write(data)
				if err != nil {
					return
				}

			case <-ctx.Done():
				close(events)
				return
			}
		}
	}(f)

	return events, nil

}
