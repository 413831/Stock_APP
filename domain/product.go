package domain

type Product interface {
	ShowDetail() string
	SetPrice(price int)
}
