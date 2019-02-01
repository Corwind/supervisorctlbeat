package beater

import (
    "fmt"
    "time"

    "os/exec"
    "strings"

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


    //var out bytes.Buffer
    logp.Info("Before vars")


    ticker := time.NewTicker(bt.config.Period)
    for {
        select {
        case <-bt.done:
            return nil
        case <-ticker.C:
        }

        logp.Info("Starting command")
        output_, err := exec.Command("supervisorctl", "-c", bt.config.SupervisordConfig, "status").Output()

        if err != nil {
            logp.Err("An error occured while executing command: %v", err)
            return err
        }

        status := common.MapStr {
            "RUNNING": 1,
            "FATAL": -1,
            "STOPPED": -2,
        }

        var services = common.MapStr{}

        output := strings.Split(string(output_), "\n")


        for _, element := range output {

            if len(element) == 0 { break }

            var split_element = strings.Fields(element)
            var state = common.MapStr{}

            state.Put("state", status[split_element[1]])
            state.Put("state_str", split_element[1])

            if (split_element[1] == "STOPPED" ||
                split_element[1] == "FATAL") {

                state.Put("reason", strings.Join(split_element[2:], " "))

            } else if (split_element[1] == "RUNNING") {

                pid := split_element[3]
                state.Put("pid", pid[:len(pid) - 1])
                state.Put("uptime", strings.Join(split_element[5:], " "))

            }

            services.Put(split_element[0], state)
        }

        fields.Put("services", services)

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
