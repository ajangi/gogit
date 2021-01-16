package rest

type Request struct {
	Headers []string
	Uri       string
	PerPage   int
}