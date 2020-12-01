// В пакете определяется интерфейс WordProvider
package wordsprovider

// Признаки:
// * EOF: слова отсутствуют, входные данные закончились;
// * OnEdge: слово находится на границе предложения;
// * Regular: слово без особых признаков.
type WordTag int

const (
	Regular WordTag = iota
	EOF
	OnEdge
)

type WordsProvider interface {
	// GetWord возвращает слово из текста и его признак
	GetWord() (string, WordTag)
}
