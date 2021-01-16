package rest

type Request struct {
	Headers []string
	Uri       string
	PerPage   int
}

type Flag struct {
	User        string
	PerPage     int
}