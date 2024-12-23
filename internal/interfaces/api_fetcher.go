package interfaces

type ApiFetcher interface {
	FetchData() (string, error)
}