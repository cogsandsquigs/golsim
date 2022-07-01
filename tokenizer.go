package main

type Tokenizer struct {
	model map[string]Token // the tokenizer's internal "model" of tokens
}

func New() *Tokenizer {
	return &Tokenizer{}
}

// Trains the tokenizer's model on data, so that it can be improved/updated
func (t *Tokenizer) Train(data string) {
	if t.model == nil {
		t.model = map[string]Token{}
	}

}
