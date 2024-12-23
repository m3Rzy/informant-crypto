package interfaces

type ApiFetcher interface {
	FetchData() (string, error)
	GetCurrency() string
	GetRate() string
	GetSource() string
}

type Fetcher interface {
	GetCurrency() string
	GetRate() string
	GetSource() string
}

type Apis interface {
	Fetch(body []byte) string
}