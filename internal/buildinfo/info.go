package buildinfo

import (
	"os"
	"runtime"
	"time"
)

var (
	_version     = "unknown-_version"  //nolint:gochecknoglobals // nothing
	_gitCommit   = "unknown-commit"    //nolint:gochecknoglobals // nothing
	_buildTime   = "unknown-buildtime" //nolint:gochecknoglobals // nothing
	_APIVersion  = "v0.1.0"            //nolint:gochecknoglobals // nothing
	_buildBranch = "unknown-branch"    //nolint:gochecknoglobals // nothing
	_buildUser   = "unknown-user"      //nolint:gochecknoglobals // nothing
	_startTime   = time.Now()          //nolint:gochecknoglobals // nothing
	// _dependencies = "unknown-deps"      //nolint:gochecknoglobals // nothing.
)

type Info struct {
	Version     string `json:"_version"`
	APIVersion  string `json:"api_version"`
	GitCommit   string `json:"git_commit"`
	BuildBranch string `json:"build_branch"`
	BuildUser   string `json:"build_user"`
	BuildTime   string `json:"build_time"`

	GoVersion  string `json:"go_version"`
	GoCompiler string `json:"go_compiler"`
	OS         string `json:"os"`
	Arch       string `json:"arch"`
	Hostname   string `json:"hostname"`
	PID        int    `json:"pid"`
	CPUCount   int    `json:"cpu_count"`

	Uptime      string `json:"uptime"`
	CurrentTime string `json:"current_time"`

	// _dependencies string `json:"dependencies"`
}

func Get() Info {
	hostname, _ := os.Hostname()

	return Info{
		Version:     _version,
		APIVersion:  _APIVersion,
		GitCommit:   _gitCommit,
		BuildBranch: _buildBranch,
		BuildUser:   _buildUser,
		BuildTime:   _buildTime,
		GoVersion:   runtime.Version(),
		GoCompiler:  runtime.Compiler,
		OS:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		Hostname:    hostname,
		PID:         os.Getpid(),
		CPUCount:    runtime.NumCPU(),
		Uptime:      time.Since(_startTime).Round(time.Second).String(),
		CurrentTime: time.Now().Format(time.RFC3339),
		// _dependencies: _dependencies,
	}
}

func (i Info) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"_version":     i.Version,
		"api_version":  i.APIVersion,
		"git_commit":   i.GitCommit,
		"build_branch": i.BuildBranch,
		"build_user":   i.BuildUser,
		"build_time":   i.BuildTime,
		"go_version":   i.GoVersion,
		"go_compiler":  i.GoCompiler,
		"os":           i.OS,
		"arch":         i.Arch,
		"hostname":     i.Hostname,
		"pid":          i.PID,
		"cpu_count":    i.CPUCount,
		"uptime":       i.Uptime,
		"current_time": i.CurrentTime,
		// "dependencies": i._dependencies,
	}
}
