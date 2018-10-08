<template>
  <div class="body">
    <h1>換字暗号を解読してみよう!!</h1>
    <mytable />
    <p>正解数：{{ correct }}</p>
    <pre style="text-align: left">{{ text }}</pre>
  </div>
</template>

<script>
import Vue from 'vue'
import Table from '../modules/Table'
import Store from '../../stores/index'

export default {
  computed: {
    text () {
      return Store.state.text
    },
    correct () {
      return Store.state.count
    }
  },
  created () {
    Store.dispatch('getQuiz', this.$route.params.id)
  }
}

Vue.component('mytable', Table)

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
