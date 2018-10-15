package crypter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	letter = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

// MyFileは、暗号文とその鍵を持つ構造たい
type MyFile struct {
	Key  []string // 鍵
	Text string   // 暗号文
}

// 指定されたIDのファイルを読んでMyFileを返す関数
// ${id}.keyが鍵ファイル、${id}.txtが暗号文
func ReadByID(id int64) (*MyFile, error) {
	key, err := ioutil.ReadFile(fmt.Sprintf("./quiz/%d.key", id))
	if err != nil {
		return nil, err
	}

	text, err := ioutil.ReadFile(fmt.Sprintf("./quiz/%d.txt", id))
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(strings.NewReader(string(key)))
	keySlice, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(keySlice[0]) != 26 {
		return nil, errors.New("key must has 26 letters")
	}

	f := &MyFile{
		Key:  keySlice[0],
		Text: string(text),
	}

	return f, nil
}

func Encrypt(id int64, key []string) (string, error) {
	pair := make([]string, 52)
	for i := 0; i < 52; i += 2 {
		pair[i] = letter[i/2]
		pair[i+1] = key[i/2]
	}
	r := strings.NewReplacer(pair...)
	f, err := ReadByID(id)
	if err != nil {
		return "", err
	}
	return r.Replace(f.Text), nil
}

func CountCorrect(id int64, key []string) (int64, error) {
	f, err := ReadByID(id)
	if err != nil {
		return 0, err
	}
	m := make(map[string]string, 26)
	for i := 0; i < 26; i++ {
		m[letter[i]] = key[i]
	}
	mk := make(map[string]string, 26)
	for i := 0; i < 26; i++ {
		mk[f.Key[i]] = letter[i]
	}
	count := 0
	for i := 0; i < len(key); i++ {
		if m[letter[i]] == mk[letter[i]] {
			count++
		}
	}
	return int64(count), nil
}
