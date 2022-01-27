package ting

import "rsc.io/quote"

func TingQuote() string {
	return quote.Go() + "\n" + quote.Glass()
}

func Ting2Quote() string {
	return quote.Opt() + "\n" + quote.Hello()
}
