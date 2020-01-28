package w2v

import (
	"testing"
)

// func TestReadWord2VecBinary_run(t *testing.T) {
// 	f, err := os.Open("./analogy.bin")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer f.Close()
//
// 	embeds, err := ReadWord2VecBinary(bufio.NewReader(f))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	fmt.Println(embeds)
// }

func TestReadWord2VecText_run(t *testing.T) {
	ReadWord2VecText("./enwiki_20180420_100d.txt")
}
