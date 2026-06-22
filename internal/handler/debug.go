package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
	"runtime/debug"
	"time"
)

var startedAt = time.Now()

type Info struct {
	GoVersion string `json:"go_version"`
	Module    string `json:"module"`
	Version   string `json:"version"`

	CommitHash string `json:"commit_hash"`
	Revision   string `json:"revision"`
	Modified   string `json:"modified"`

	StartedAt string `json:"started_at"`
	Uptime    string `json:"uptime"`
}

func HandleDebugInfo(w http.ResponseWriter, r *http.Request) {
	info := Info{
		GoVersion: runtime.Version(),
		StartedAt: startedAt.Format(time.RFC3339),
		Uptime:    time.Since(startedAt).String(),
	}

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		info.Module = buildInfo.Main.Path
		info.Version = buildInfo.Main.Version

		for _, setting := range buildInfo.Settings {
			switch setting.Key {
			case "vcs.revision":
				info.Revision = setting.Value
				info.CommitHash = setting.Value
			case "vcs.modified":
				info.Modified = setting.Value
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(info)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
