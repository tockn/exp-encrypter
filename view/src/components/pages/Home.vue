<template>
  <div class="body">
    <h1>換字暗号を解読してみよう!!</h1>
    <mytable />
    <p>正解数：{{ correct }}</p>
    <freq v-for="(word, index) in freqWords" :key="index" :word="word" />
    <pre style="text-align: left">{{ text }}</pre>
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
    }
  },
  created () {
    Store.dispatch('getQuiz', this.$route.params.id)
    Store.dispatch('getFreq', this.$route.params.id)
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
  text-align: center;
}

</style>
