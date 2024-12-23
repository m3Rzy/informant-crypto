package interfaces

type Getter interface {
	GetCurrency() string
	GetRate() string
	GetSource() string
}