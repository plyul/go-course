package parser

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

type TextParser struct {
	r           io.Reader
	textScanner *bufio.Scanner
	lineScanner *bufio.Scanner
	words       []string
}

func New(r io.Reader) *TextParser {
	return &TextParser{
		r:           r,
		textScanner: bufio.NewScanner(r),
		lineScanner: nil,
	}
}

// Scan возвращает следующее допустимое слово из текста или пустую строку, если текст закончился
// Допустимое слово:
// - Длиннее трёх букв
// - Не первое и не последнее в предложении
// - Предложения в тексте разделены символами точки или новой строки.
func (p *TextParser) Scan() string {
	for { // Сканируем слова и строки, подыскивая подходящее слово для возврата
		for p.words == nil { // Буфер слов пуст, нужно сканировать новую строку
			if !p.textScanner.Scan() { // Текст закончен, закончили сканирование
				return ""
			}
			line := p.textScanner.Text()
			p.lineScanner = bufio.NewScanner(strings.NewReader(line))
			p.lineScanner.Split(bufio.ScanWords)
			for p.lineScanner.Scan() { // Сканируем слова из строки в буферный слайс
				p.words = append(p.words, p.lineScanner.Text())
			}
			if len(p.words) > 2 { // Убираем первое и последнее слова строки, т.к. они всегда стоят рядом с '\n'
				p.words = p.words[1 : len(p.words)-1]
			} else {
				p.words = nil
			}
			w := p.getValidWord()
			if w != "" {
				return w
			}
		}
		w := p.getValidWord()
		if w != "" {
			return w
		}
	}
}

// getValidWord возвращает допустимое слово из буфера (p.words) или пустую строку, если буфер пуст
// Если функция возвращает пустую строку, гарантируется, что p.words == nil
func (p *TextParser) getValidWord() string {
	var w string
	for len(p.words) > 0 {
		w = filterPunctuation(p.words[0])
		if strings.Contains(w, ".") { // В слове есть точка, значит оно последнее в предложении. Выбрасываем его и следующее слово
			if len(p.words) > 2 {
				w = ""
				p.words = p.words[2:]
			} else { // В буфере одно/два слова рядом с точкой, они нам не подходят
				p.words = nil
				return ""
			}
		} else {
			p.words = p.words[1:]
		}
		if len(p.words) == 0 {
			p.words = nil
		}
		if len(w) > 3 {
			return w
		}
	}
	return ""
}

// filterPunctuation удаляет все символы с краёв слова, кроме букв и символа точки
// Символ точки нам понадобится, чтобы определять границу предложения
func filterPunctuation(s string) string {
	r := strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && r != '.'
	})
	return r
}
