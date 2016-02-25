package main

type Word struct {
	Text    []byte
	Meaning []byte
}

func (word *Word) addBytes(textByte, meaningByte byte) {
	word.Text = append(word.Text, textByte)
	word.Meaning = append(word.Meaning, meaningByte)
}

func (word *Word) equalizeBytesLengths() *Word {
	textLength := len(word.Text)
	meaningLength := len(word.Meaning)
	if textLength > meaningLength {
		word.Meaning = append(word.Meaning, 0)
		word.equalizeBytesLengths()
	} else if meaningLength > textLength {
		word.Text = append(word.Text, 0)
		word.equalizeBytesLengths()
	}
	return word
}

func (word *Word) textByteAt(index int) byte {
	return word.Text[index]
}
