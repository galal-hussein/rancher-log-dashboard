package container

// Container Struct
type Container struct {
	Host string
	Name string
	Logs []string
}

// Container list
type Containers map[string]Container
