package interfaces

type ApiFetcher interface {
	FetchData() (string, error)
	GetCurrency() string
	GetRate() string
	GetSource() string
}
