package jpkana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	input := []byte(
		"[{\"consonant\": \"\", \"vocal\": \"A\", \"romanji\": \"a\", \"kana\": \"あ\", \"difficulty\" : 1}]",
	)
	kanas, err := load(input)

	assert.Nil(t, err)
	assert.NotEmpty(t, kanas)
	assert.Len(t, kanas, 1)
	assert.Zero(t, kanas[0].Consonant)
	assert.Equal(t, "A", kanas[0].Vocal)
	assert.Equal(t, "a", kanas[0].Romanji)
	assert.Equal(t, "あ", kanas[0].Kana)
	assert.Equal(t, uint(1), kanas[0].Difficulty)
}

func TestKanaSequence1(t *testing.T) {
	input := []byte(
		"[{\"consonant\": \"\", \"vocal\": \"A\", \"romanji\": \"a\", \"kana\": \"あ\", \"difficulty\" : 1}]",
	)
	kanas, _ := load(input)

	kana, romanji := getKanaSequence(kanas, 1, 1)

	assert.NotEmpty(t, kana)
	assert.NotEmpty(t, romanji)
	assert.Equal(t, "あ", kana)
	assert.Equal(t, "a", romanji)
}

func TestKanaSequenceN(t *testing.T) {
	input := []byte(
		"[{\"consonant\": \"\", \"vocal\": \"A\", \"romanji\": \"a\", \"kana\": \"あ\", \"difficulty\" : 1},{\"consonant\": \"\", \"vocal\": \"A\", \"romanji\": \"a\", \"kana\": \"あ\", \"difficulty\" : 1}]",
	)
	kanas, _ := load(input)

	kana, romanji := getKanaSequence(kanas, 5, 1)

	assert.NotEmpty(t, kana)
	assert.NotEmpty(t, romanji)
	assert.Equal(t, "あああああ", kana)
	assert.Equal(t, "aaaaa", romanji)
}

func TestKanaSequenceDifficultyTooHigh(t *testing.T) {
	input := []byte(
		"[{\"consonant\": \"\", \"vocal\": \"A\", \"romanji\": \"a\", \"kana\": \"あ\", \"difficulty\" : 1},{\"consonant\": \"\", \"vocal\": \"A\", \"romanji\": \"a\", \"kana\": \"あ\", \"difficulty\" : 1}]",
	)
	kanas, _ := load(input)

	kana, romanji := getKanaSequence(kanas, 5, 0)

	assert.Empty(t, kana)
	assert.Empty(t, romanji)
}

func TestGeneratorError(t *testing.T) {
	hiragana := []byte(
		"[{\"consonant\": \"\", \"vocal\": \"A\", \"romanji\": \"a\", \"kana\": \"あ\", \"difficulty\" : 1}]",
	)
	katakana := []byte(
		"[{\"consonant\": \"\", \"vocal\": \"A\", \"romanji\": \"a\", \"kana\": \"ア\", \"difficulty\" : 1}]",
	)
	kG, err := New(hiragana, katakana)

	assert.NotNil(t, kG)
	assert.Nil(t, err)

	kana, romanji := kG.Generate("BOH", 5, 1)

	assert.Empty(t, kana)
	assert.Empty(t, romanji)
}
func TestGenerator(t *testing.T) {
	hiragana := []byte(
		"[{\"consonant\": \"\", \"vocal\": \"A\", \"romanji\": \"a\", \"kana\": \"あ\", \"difficulty\" : 1}]",
	)
	katakana := []byte(
		"[{\"consonant\": \"\", \"vocal\": \"A\", \"romanji\": \"a\", \"kana\": \"ア\", \"difficulty\" : 1}]",
	)
	kG, err := New(hiragana, katakana)

	assert.NotNil(t, kG)
	assert.Nil(t, err)

	kana, romanji := kG.Generate(HIRAGANA, 5, 1)
	assert.NotEmpty(t, kana)
	assert.NotEmpty(t, romanji)
	assert.Equal(t, "あああああ", kana)
	assert.Equal(t, "aaaaa", romanji)

	kana, romanji = kG.Generate(KATAKANA, 5, 1)
	assert.NotEmpty(t, kana)
	assert.NotEmpty(t, romanji)
	assert.Equal(t, "アアアアア", kana)
	assert.Equal(t, "aaaaa", romanji)
}
