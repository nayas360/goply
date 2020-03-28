package goply

type TokenStream struct {
	tokens *[]*Token // the slice
	p      int       // pointer
}

// returns the token at the nth position
// does not affect
func (t *TokenStream) GetAt(n int) *Token {
	if n >= len(*t.tokens) || n < 0 {
		return nil
	}
	return (*t.tokens)[n]
}

func (t *TokenStream) Get() *Token {
	if t.EOS() || t.p < 0 {
		return nil
	}
	return (*t.tokens)[t.p]
}

func (t *TokenStream) Iter() chan *Token {
	ch := make(chan *Token, t.Len())
	for i := 0; i < t.Len(); i++ {
		ch <- (*t.tokens)[i]
	}
	close(ch)
	return ch
}

// This sets internal pointer to move by k
// positive integers move forward and negative move backward
// values are clipped between [0, #Tokesn]
func (t *TokenStream) PosDelta(k int) *TokenStream {
	t.p += k
	if t.p >= len(*t.tokens) {
		t.p = len(*t.tokens)
	} else if t.p < 0 {
		t.p = 0
	}
	return t
}

func (t *TokenStream) Pos() int {
	return t.p
}

func (t *TokenStream) ResetPos() *TokenStream {
	t.p = 0
	return t
}

func (t *TokenStream) Len() int {
	return len(*t.tokens)
}

// Returns true when the token stream has finished
func (t *TokenStream) EOS() bool {
	return t.p >= t.Len()
}
