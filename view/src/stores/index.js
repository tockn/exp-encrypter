import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

var URL = process.env.API_ENDPOINT

var Store = new Vuex.Store({
  state: {
    text: '',
    key: {'A': 'A', 'B': 'B', 'C': 'C', 'D': 'D', 'E': 'E', 'F': 'F', 'G': 'G', 'H': 'H', 'I': 'I', 'J': 'J', 'K': 'K', 'L': 'L', 'M': 'M', 'N': 'N', 'O': 'O', 'P': 'P', 'Q': 'Q', 'R': 'R', 'S': 'S', 'T': 'T', 'U': 'U', 'V': 'V', 'W': 'W', 'X': 'X', 'Y': 'Y', 'Z': 'Z'},
    count: 0
  },
  mutations: {
    getQuiz (state, response) {
      state.text = response.data.text
    },
    encrypt (state, response) {
      state.text = response.data.text
      state.count = response.data.correct
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
      axios.post(`${URL}/api/quiz/${id}`, {
        key: key
      })
        .then(response => {
          commit('encrypt', response)
        })
    }
  }
})

export default Store
