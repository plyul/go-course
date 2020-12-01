// Пакет textparser предоставляет функционал извлечения слов из текста по правилам Unicode.
// Разделителем предложения считается символ точки или переноса строки, остальные символы пунктуации игнорируются.
// Возвращаемым словам присваивается признак, в зависимости от положения слова в тексте.
package textparser

import (
	"bufio"
	"go-course/task3/internal/wordsprovider"
	"io"
	"strings"
	"unicode"
)

type TextParser struct {
	textScanner   *bufio.Scanner
	lineScanner   *bufio.Scanner
	words         []string
	sentenceBegin bool
}

// New возвращает экземпляр TextParser, возвращающий слова из переданного интерфейса io.Reader.
func New(r io.Reader) *TextParser {
	return &TextParser{
		textScanner:   bufio.NewScanner(r),
		lineScanner:   nil,
		words:         nil,
		sentenceBegin: true,
	}
}

// GetWord реализует соответствующий метод интерфейса WordsProvider
func (p *TextParser) GetWord() (string, wordsprovider.WordTag) {
	for { // Сканируем слова и строки, подыскивая подходящее слово для возврата
		if p.words == nil { // Буфер слов пуст, нужно сканировать новую строку
			if !p.textScanner.Scan() { // Текст закончен, закончили сканирование
				return "", wordsprovider.EOF
			}
			p.sentenceBegin = true
			line := p.textScanner.Text()
			p.lineScanner = bufio.NewScanner(strings.NewReader(line))
			p.lineScanner.Split(bufio.ScanWords)
			for p.lineScanner.Scan() { // Сканируем слова из строки в буферный слайс
				p.words = append(p.words, p.lineScanner.Text())
			}
		}
		w, s := p.wordFromBuffer()
		if s != wordsprovider.EOF {
			return w, s
		}
	}
}

// wordFromBuffer возвращает слово из буфера p.words.
// Если функция возвращает wordsprovider.EOF, гарантируется, что p.words == nil.
func (p *TextParser) wordFromBuffer() (string, wordsprovider.WordTag) {
	var w string
	if len(p.words) > 0 {
		w = strings.ToLower(distilWord(p.words[0]))
		p.words = p.words[1:]
		if len(p.words) == 0 {
			p.words = nil
		}
		s := wordsprovider.Regular
		if p.sentenceBegin {
			s = wordsprovider.OnEdge
			p.sentenceBegin = false
		}
		if strings.HasSuffix(w, ".") { // слово содержит точку в конце, т.е. оно на границе предложения
			w = strings.Trim(w, ".")
			s = wordsprovider.OnEdge // вернётся для текущего слова
			p.sentenceBegin = true   // условие сыграет при следующем вызове функции
		}
		if len(p.words) == 1 { // осталось одно слово в буфере, т.е. оно на границе предложения
			p.sentenceBegin = true // условие сыграет при следующем вызове функции
		}
		return w, s
	}
	p.words = nil
	return "", wordsprovider.EOF
}

// distilWord удаляет из слова все символы, кроме букв и точки
// Символ точки нам понадобится, чтобы определять границу предложения.
func distilWord(s string) string {
	var result []rune
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) || r == '.' {
			result = append(result, r)
		}
	}
	return string(result)
}
