package handler

import (
	"os"
	"time"

	"github.com/golang/glog"
)

// UpdateDatabase will update the latest pipelines detail and status
// TODO
func UpdateDatabase() {
	// Read token environment variable
	token, ok := os.LookupEnv(token)
	if !ok {
		glog.Fatalf("TOKEN environment variable required")
	}
	// Update the database, This wil run only first time
	BuildData(token)
	k8sVersion := []string{"v11", "v12", "v13"}
	for _, k8sVersion := range k8sVersion {
		columnName := "packet_" + k8sVersion + "_pid"
		pipelineTable := "packet_pipeline_" + k8sVersion
		jobTable := "packet_jobs_" + k8sVersion
		PacketData(token, columnName, pipelineTable, jobTable)
	}
	// KonvoyData(token, "konvoy", "konvoy_pipeline", "konvoy_jobs")
	// loop will iterate at every 2nd minute and update the database
	tick := time.Tick(2 * time.Minute)
	for range tick {
		BuildData(token)
		for _, k8sVersion := range k8sVersion {
			columnName := "packet_" + k8sVersion + "_pid"
			pipelineTable := "packet_pipeline_" + k8sVersion
			jobTable := "packet_jobs_" + k8sVersion
			PacketData(token, columnName, pipelineTable, jobTable)
		}
		// KonvoyData(token, "konvoy", "konvoy_pipeline", "konvoy_jobs")

	}
}
