package jpkana

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/ledongthuc/goterators"
)

const (
	HIRAGANA = "hiragana"
	KATAKANA = "katakana"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Kana struct {
	Consonant  string `json:"consonant"`
	Vocal      string `json:"vocal"`
	Romanji    string `json:"romanji"`
	Kana       string `json:"kana"`
	Difficulty uint   `json:"difficulty"`
}

func load(raw []byte) ([]Kana, error) {
	var kanas []Kana

	if err := json.Unmarshal(raw, &kanas); err != nil {
		return nil, err
	}
	return kanas, nil
}

func getKanaSequence(kanas []Kana, length uint, difficulty uint) (string, string) {
	kanaByDifficulty := goterators.Filter(kanas, func(item Kana) bool {
		return item.Difficulty <= difficulty
	})

	if len(kanaByDifficulty) == 0 {
		return "", ""
	}

	kana := ""
	romanji := ""

	for i := 0; i < int(length); i++ {
		idx := rand.Intn(len(kanaByDifficulty))
		kana += kanaByDifficulty[idx].Kana
		romanji += kanaByDifficulty[idx].Romanji
	}
	return kana, romanji
}

type KanaGenerator struct {
	kanas map[string][]Kana
}

func New(hbytes []byte, kbytes []byte) (*KanaGenerator, error) {
	hiragana, err := load(hbytes)
	if err != nil {
		return nil, err
	}
	katakana, err := load(kbytes)
	if err != nil {
		return nil, err
	}
	kG := &KanaGenerator{
		kanas: map[string][]Kana{
			HIRAGANA: hiragana,
			KATAKANA: katakana,
		},
	}
	return kG, nil
}

func (kG KanaGenerator) generate(kanaName string, length uint, difficulty uint) (string, string) {
	kanaList, ok := kG.kanas[kanaName]
	if !ok {
		return "", ""
	}
	return getKanaSequence(kanaList, length, difficulty)
}
