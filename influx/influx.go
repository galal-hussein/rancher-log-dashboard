package influx

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/galal-hussein/rancher-log-dashboard/container"
	influx "github.com/influxdata/influxdb/client/v2"
)

type InfluxClient struct {
	client influx.Client
	DB     string
}

var influxClient InfluxClient

// Init function
func init() {
	// Parse env variables
	influxHost := os.Getenv("INFLUX_HOST")
	if len(influxHost) == 0 {
		logrus.Errorf("INFLUX_HOST is not set")
	}
	influxPort := os.Getenv("INFLUX_PORT")
	if len(influxPort) == 0 {
		logrus.Errorf("INFLUX_PORT is not set")
	}
	influxClient.DB = os.Getenv("INFLUX_DB")
	if len(influxClient.DB) == 0 {
		logrus.Errorf("INFLUX_DB is not set")
	}
	logrus.Debugf("Initializing Influxdb Connection with host: %s, port: %d, database: %s",
		influxHost, influxPort, influxClient.DB)

	var err error

	influxClient.client, err = influx.NewHTTPClient(influx.HTTPConfig{
		Addr: "http://" + influxHost + ":" + influxPort,
	})
	if err != nil {
		logrus.Fatal(err)
	}
}

func GetContainers(stackname string) container.Containers {
	logrus.Debugf("Querying Influxdb with measurement: %s", stackname)
	containers := make(container.Containers)
	selectCmd := "SELECT * FROM " + stackname
	q := influx.NewQuery(selectCmd, influxClient.DB, "")
	if response, err := influxClient.client.Query(q); err == nil && response.Error() == nil {
		if len(response.Results[0].Series) == 0 {
			logrus.Errorf("No measurements found with the %s", stackname)
			return nil
		}
		res := response.Results[0].Series[0]
		for _, record := range res.Values {
			containerID := record[1].(string)
			containerName := record[2].(string)
			containerLog := record[3].(string)
			var c container.Container
			if _, ok := containers[containerID]; ok {
				logs := append(containers[containerID].Logs, containerLog)
				c.Name = containerName
				c.Logs = logs
				containers[containerID] = c
			} else {
				c.Name = containerName
				c.Logs = append(c.Logs, containerLog)
				containers[containerID] = c
			}
		}
	}
	return containers
}
