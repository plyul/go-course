// Пакет chunkreader реализует интерфейс io.Reader для чтения части файла (чанка),
// ограниченного смещениями [startOffset:endOffset]
package chunkreader

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	bufSize = 65536
	chunkOffsetsSliceCapacity = 100
)

type ChunkReader struct {
	f         *os.File
	curOffset int64
	endOffset int64
	buffer    []byte
}

// New возвращает новый объект ChunkReader для заданных
// файла f и чанка [startOffset:endOffset]
func New(f *os.File, startOffset int64, endOffset int64) *ChunkReader {
	cr := ChunkReader{
		f:         f,
		curOffset: startOffset,
		endOffset: endOffset,
		buffer:    make([]byte, bufSize),
	}
	return &cr
}

// Read реализует интерфейс io.Reader, позволяя прочитать чанк от startOffset до endOffset
func (cr *ChunkReader) Read(p []byte) (n int, err error) {
	maxBytesToRead := int64(len(p))
	if maxBytesToRead == 0 {
		return 0, errors.New("read: empty input buffer")
	}
	off, err := cr.f.Seek(cr.curOffset, 0)
	if err != nil || off != cr.curOffset {
		return 0, errors.New(fmt.Sprintf("read: %v", err))
	}
	if cr.curOffset+maxBytesToRead > cr.endOffset {
		maxBytesToRead = cr.endOffset - cr.curOffset
	}
	if maxBytesToRead == 0 {
		return 0, io.EOF
	}
	n, err = cr.f.Read(p[0:maxBytesToRead])
	cr.curOffset += int64(n)
	return n, err
}

// ChunkOffsets возвращает слайс смещений, по которым проходят границы чанков
// Границы чанков всегда находятся на границе строки (см. shiftToLF)
func ChunkOffsets(f *os.File, chunkSize int64) ([]int64, error) {
	result := make([]int64, 0, chunkOffsetsSliceCapacity)
	curOffset := chunkSize
	fs, err := f.Stat()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("chunkOffsets error: %v", err))
	}
	maxOffset := fs.Size()
	result = append(result, 0) // Первый чанк всегда начинается с нулевого смещения
	if curOffset > maxOffset { // Файл мал и может быть обработан одним чанком
		return result, nil
	}
	for {
		offset, err := f.Seek(curOffset, 0)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("chunkOffsets error: %v", err))
		}
		delta, err := shiftToLF(f)
		offset += delta
		if err != nil {
			return nil, errors.New(fmt.Sprintf("chunkOffsets error: %v", err))
		}
		result = append(result, offset)
		curOffset += chunkSize
		if curOffset > maxOffset {
			break
		}
	}
	result = append(result, maxOffset)
	return result, nil
}

// shiftToLF возвращает дельту от текущего смещения в файле f до границы строки (символ \n)
// Функция имеет сайдэффект - она меняет значение текущего смещения для файла f
func shiftToLF(f *os.File) (int64, error) {
	var buffer = make([]byte, 65536)
	n, err := f.Read(buffer)
	if err != nil && err != io.EOF {
		return 0, errors.New(fmt.Sprintf("shiftToLF error: %v", err))
	}
	for i, v := range buffer[:n] {
		if v == '\n' {
			return int64(i), nil
		}
	}
	return 0, nil
}
