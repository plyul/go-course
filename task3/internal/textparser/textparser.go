// Пакет textparser предоставляет функционал извлечения слов из текста по правилам Unicode.
// Разделителем предложения считается символ точки или переноса строки, остальные символы пунктуации игнорируются.
// Возвращаемым словам присваивается признак, в зависимости от положения слова в тексте.
package textparser

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

type TextParser struct {
	textScanner   *bufio.Scanner
	lineScanner   *bufio.Scanner
	words         []string
	wordIndex     int
	sentenceBegin bool
}

// Признаки:
// * EOF: слова отсутствуют, входные данные закончились;
// * OnEdge: слово находится на границе предложения;
// * Regular: слово без особых признаков.
type WordTag int

const (
	Regular WordTag = iota
	OnEdge
	EOF
)

type WordData struct {
	Word        string
	Tag         WordTag
	ChunkOffset int64
	ChunkIdx    int
}

// New возвращает экземпляр TextParser, возвращающий слова из переданного интерфейса io.Reader.
func New(r io.Reader) *TextParser {
	return &TextParser{
		textScanner:   bufio.NewScanner(r),
		lineScanner:   nil,
		words:         nil,
		wordIndex:     0,
		sentenceBegin: true,
	}
}

// GetWord возвращает структуру WordData, описывающую слово, полученное из io.Reader
func (p *TextParser) GetWord() WordData {
	for { // Сканируем слова и строки, подыскивая подходящее слово для возврата
		if p.words == nil { // Буфер слов пуст, нужно сканировать новую строку
			if !p.textScanner.Scan() { // Текст закончен, закончили сканирование
				return WordData{
					Word:     "",
					Tag:      EOF,
					ChunkIdx: p.wordIndex,
				}
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
		p.wordIndex++
		if s != EOF {
			return WordData{
				Word:     w,
				Tag:      s,
				ChunkIdx: p.wordIndex,
			}
		}
	}
}

// wordFromBuffer возвращает слово из буфера p.words.
// Если функция возвращает wordsprovider.EOF, гарантируется, что p.words == nil.
func (p *TextParser) wordFromBuffer() (string, WordTag) {
	var w string
	if len(p.words) > 0 {
		w = strings.ToLower(distilWord(p.words[0]))
		p.words = p.words[1:]
		if len(p.words) == 0 {
			p.words = nil
		}
		s := Regular
		if p.sentenceBegin {
			s = OnEdge
			p.sentenceBegin = false
		}
		if strings.HasSuffix(w, ".") { // слово содержит точку в конце, т.е. оно на границе предложения
			w = strings.Trim(w, ".")
			s = OnEdge             // вернётся для текущего слова
			p.sentenceBegin = true // условие сыграет при следующем вызове функции
		}
		if len(p.words) == 1 { // осталось одно слово в буфере, т.е. оно на границе предложения
			p.sentenceBegin = true // условие сыграет при следующем вызове функции
		}
		return w, s
	}
	p.words = nil
	return "", EOF
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
