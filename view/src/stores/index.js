import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

var URL = process.env.API_ENDPOINT

const keykey = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']

var Store = new Vuex.Store({
  state: {
    text: '',
    key: {'A': 'A', 'B': 'B', 'C': 'C', 'D': 'D', 'E': 'E', 'F': 'F', 'G': 'G', 'H': 'H', 'I': 'I', 'J': 'J', 'K': 'K', 'L': 'L', 'M': 'M', 'N': 'N', 'O': 'O', 'P': 'P', 'Q': 'Q', 'R': 'R', 'S': 'S', 'T': 'T', 'U': 'U', 'V': 'V', 'W': 'W', 'X': 'X', 'Y': 'Y', 'Z': 'Z'},
    count: 0,
    freqWords: []
  },
  mutations: {
    getQuiz (state, response) {
      state.text = response.data.text
    },
    encrypt (state, response) {
      state.text = response.data.text
      state.count = response.data.correct
    },
    getFreq (state, response) {
      state.freqWords = response.data.words
    },
    setKey (state, key) {
      state.key = key
      console.log(state.key)
    }
  },
  actions: {
    getQuiz ({ commit }, id) {
      axios.get(`${URL}/api/quiz/${id}`)
        .then(response => {
          commit('getQuiz', response)
        })
    },
    encrypt ({ commit, state }, {id, key}) {
      state.key = key
      let k = []
      for (let i = 0; i < keykey.length; i++) {
        k.push(state.key[keykey[i]])
      }
      axios.post(`${URL}/api/quiz/${id}`, {
        key: k
      })
        .then(response => {
          commit('encrypt', response)
        })
    },
    getFreq ({ commit }, id) {
      axios.get(`${URL}/api/quiz/${id}/freq`)
        .then(response => {
          commit('getFreq', response)
        })
    }
  }
})

export default Store
