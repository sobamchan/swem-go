package w2v

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	// "github.com/gonum/blas"
	// cblas "github.com/gonum/blas/cgo"
	"strings"
)

type Embeddings struct {
	// blas      blas.Float32Level2
	matrix    []float32
	embedSize int
	indices   map[string]int
	words     []string
}

func NewEmbeddings(embedSize int) *Embeddings {
	return &Embeddings{
		// blas:      cblas.Implementation{},
		matrix:    make([]float32, 0),
		embedSize: embedSize,
		indices:   make(map[string]int),
		words:     make([]string, 0),
	}
}

func ReadWord2VecBinary(r *bufio.Reader) (*Embeddings, error) {
	var nWords uint64
	if _, err := fmt.Fscanf(r, "%d", &nWords); err != nil {
		return nil, err
	}

	var vSize uint64
	if _, err := fmt.Fscanf(r, "%d", &vSize); err != nil {
		return nil, err
	}

	matrix := make([]float32, nWords*vSize)
	indices := make(map[string]int)
	words := make([]string, nWords)

	for idx := 0; idx < int(nWords); idx++ {
		word, err := r.ReadString(' ')
		word = strings.TrimSpace(word)
		indices[word] = idx
		words[idx] = word

		start := idx * int(vSize)
		if err = binary.Read(r, binary.LittleEndian, matrix[start:start+int(vSize)]); err != nil {
			return nil, err
		}
	}

	return &Embeddings{
		// blas:      cblas.Implementation{},
		matrix:    matrix,
		embedSize: int(vSize),
		indices:   indices,
		words:     words,
	}, nil
}

// func ReadWord2VecText(r *bufio.Reader) (*Embeddings, error) {
func ReadWord2VecText(filepath string) (*string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {

			params := strings.Split(line, " ")
			nWordsStr := params[0]
			vSizeStr := params[1]

			nWords, err := strconv.ParseInt(nWordsStr, 10, 64)
			if err != nil {
				return nil, err
			}

			vSize, err := strconv.ParseInt(vSizeStr, 10, 64)
			if err != nil {
				return nil, err
			}

			matrix := make([][]float32, nWords)
			indices := make(map[string]int)
			words := make([]string, nWords)

			fmt.Println(nWords)
			fmt.Println(vSize)

		} else {
			instanceStr := strings.Split(line, " ")
			word := instanceStr[0]
			indices[word] = i + 1
			words[i+1] = instanceStr[0]
		}

		i++
	}

	return nil, nil
}
