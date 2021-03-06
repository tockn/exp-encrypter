package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tockn/exp-encrypter/analyzer"
	"github.com/tockn/exp-encrypter/crypter"
)

func main() {
	var (
		addr = flag.String("addr", ":8080", "addr to bind")
	)
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/api/quiz/{quizID}", GetQuiz).Methods("GET")
	router.HandleFunc("/api/quiz/{quizID}", Encrypt).Methods("POST")
	router.HandleFunc("/api/quiz/{quizID}", Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/quiz/{quizID}/freq", GetFreqWords).Methods("GET")

	router.PathPrefix("/static/img/").Handler(
		http.StripPrefix("/static/img/", http.FileServer(http.Dir("static/img"))))
	router.PathPrefix("/static/css/").Handler(
		http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css"))))
	router.PathPrefix("/static/js/").Handler(
		http.StripPrefix("/static/js/", http.FileServer(http.Dir("static/js"))))

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	log.Printf("start listening on %s", *addr)
	http.ListenAndServe(*addr, router)
}

type RespText struct {
	Text string `json:"text"`
}

type ReqKey struct {
	Key []string `json:"key"`
}

type RespEncrypt struct {
	Correct int64  `json:"correct"`
	Text    string `json:"text"`
}

type RespFreqWords struct {
	Words []string `json:"words"`
}

func JSON(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(data)
}

func param(r *http.Request, key string) (int64, error) {
	return strconv.ParseInt(mux.Vars(r)[key], 10, 64)
}

func Preflight(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, "ok")
}

func GetQuiz(w http.ResponseWriter, r *http.Request) {
	idNum, err := param(r, "quizID")
	if err != nil {
		log.Println(err)
		return
	}
	f, err := crypter.ReadByID(idNum)
	if err != nil {
		log.Println(err)
		return
	}
	res := &RespText{
		Text: f.Text,
	}
	JSON(w, http.StatusOK, res)
}

func Encrypt(w http.ResponseWriter, r *http.Request) {
	quizID, err := param(r, "quizID")
	if err != nil {
		return
	}
	var k ReqKey
	if err := json.NewDecoder(r.Body).Decode(&k); err != nil {
		return
	}
	key := k.Key
	text, err := crypter.Encrypt(quizID, key)
	if err != nil {
		return
	}
	c, err := crypter.CountCorrect(quizID, key)
	if err != nil {
		return
	}
	res := RespEncrypt{
		Correct: c,
		Text:    text,
	}
	JSON(w, http.StatusOK, res)
}

func GetFreqWords(w http.ResponseWriter, r *http.Request) {
	quizID, err := param(r, "quizID")
	if err != nil {
		log.Println(err)
		return
	}
	f, err := crypter.ReadByID(quizID)
	if err != nil {
		log.Println(err)
		return
	}
	counts := analyzer.Analyze(f.Text)
	result := sortWord(counts)
	var words []string
	for _, res := range result {
		words = append(words, res.Word)
	}
	res := RespFreqWords{Words: words}
	JSON(w, http.StatusOK, res)
}

func sortWord(data []*analyzer.Count) []*analyzer.Count {
	sort.Slice(data, func(i, j int) bool { return data[i].Count > data[j].Count })
	return data
}
