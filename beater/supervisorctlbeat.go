package beater

import (
	"fmt"
	"time"

        "os/exec"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/Corwind/supervisorctlbeat/config"
)

// Supervisorctlbeat configuration.
type Supervisorctlbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of supervisorctlbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Supervisorctlbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts supervisorctlbeat.
func (bt *Supervisorctlbeat) Run(b *beat.Beat) error {
	logp.Info("supervisorctlbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

        var fields = common.MapStr{
                "type":    b.Info.Name,
        }

        var output []byte

	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

                output, err = exec.Command("supervisorctl", "-c", 
                                           bt.config.SupervisorConf, "status").Output()

		event := beat.Event{
			Timestamp: time.Now(),
                        Fields: fields,
		}
		bt.client.Publish(event)
		logp.Info("Event sent")
	}
}

// Stop stops supervisorctlbeat.
func (bt *Supervisorctlbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
