<template>
  <div class="body">
    <div class="header">
      <h1>換字暗号を解読してみよう!!</h1>
      <mytable />
      <div v-if="clear">
        <h3>全問正解！おめでとう！</h3>
        <button @click="next" >次の問題へ</button>
      </div>
      <p v-else>正解数：{{ correct }}</p>
    </div>
    <div class="content">
      <div class="freq box">
        <freq v-for="(word, index) in freqWords" :key="index" :word="word" />
      </div>
      <div class="text box">
        <p class="textp">{{ text }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import Table from '../modules/Table'
import Freq from '../modules/Freq'
import Store from '../../stores/index'

export default {
  computed: {
    text () {
      return Store.state.text
    },
    correct () {
      return Store.state.count
    },
    freqWords () {
      return Store.state.freqWords
    },
    clear () {
      return this.correct === 26
    }
  },
  created () {
    Store.dispatch('getQuiz', this.$route.params.id)
    Store.dispatch('getFreq', this.$route.params.id)
  },
  methods: {
    next () {
      let n = Number(this.$route.params.id) + 1
      location.href = `/${n}`
    }
  },
  components: {
    'mytable': Table,
    'freq': Freq
  }
}

</script>

<style>
input {
  width: 20px;
  height: 30px;
  font-size: 20px;
  text-align: center;
}

.body {
  height: 100%;
  text-align: center;
}

.header {
}

.content {
  display: flex;
  height: 60vh;
}

.freq {
  height: 100%;
  text-align: center;
  width: 50%;
  overflow: scroll;
}

.text {
  height: 100%;
  width: 50%;
}

.textp {
  white-space: pre-wrap;
  text-align: left;
  width: 100%;
  height: 100%;
  overflow: scroll;
}

.box {
  border: solid 3px #000000;
}

</style>
