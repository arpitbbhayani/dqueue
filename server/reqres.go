package server

type HTTPError struct {
	Message string
}

type VersionRequest struct{}

type VersionResponse struct {
	Version string
}
