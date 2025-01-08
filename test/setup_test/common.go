package setuptest

import "fmt"

type Exposed struct {
	Protocol string
	Port     string
}

type PortBinding struct {
	HostIP   string
	HostPort string
}

type DBEnv map[string]string

func (d DBEnv) ToSlice() []string {
	var s []string

	for k, v := range d {
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}

	return s
}
