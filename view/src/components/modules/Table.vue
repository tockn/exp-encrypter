<template>
  <table border="1" align="center">
    <tr>
      <td>暗号文</td>
      <td>A</td>
      <td>B</td>
      <td>C</td>
      <td>D</td>
      <td>E</td>
      <td>F</td>
      <td>G</td>
      <td>H</td>
      <td>I</td>
      <td>J</td>
      <td>K</td>
      <td>L</td>
      <td>M</td>
      <td>N</td>
      <td>O</td>
      <td>P</td>
      <td>Q</td>
      <td>R</td>
      <td>S</td>
      <td>T</td>
      <td>U</td>
      <td>V</td>
      <td>W</td>
      <td>X</td>
      <td>Y</td>
      <td>Z</td>
    </tr>
    <tr>
      <td>復号結果</td>
      <td v-for="(k, kk, i) in key" :key="i" :kk="kk" :k="k" :i="i"><input @input="change" v-model="key[kk]" v-bind:name="k" class="w" maxlength="1"/></td>
    </tr>
  </table>
</template>

<script>
import Store from '../../stores/index'

export default {
  data () {
    return {
      key: {'A': 'A', 'B': 'B', 'C': 'C', 'D': 'D', 'E': 'E', 'F': 'F', 'G': 'G', 'H': 'H', 'I': 'I', 'J': 'J', 'K': 'K', 'L': 'L', 'M': 'M', 'N': 'N', 'O': 'O', 'P': 'P', 'Q': 'Q', 'R': 'R', 'S': 'S', 'T': 'T', 'U': 'U', 'V': 'V', 'W': 'W', 'X': 'X', 'Y': 'Y', 'Z': 'Z'},
      keykey: ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
    }
  },
  methods: {
    change () {
      let k = []
      for (let i = 0; i < this.keykey.length; i++) {
        if (this.key[this.keykey[i]] === '') return
        k.push(this.key[this.keykey[i]])
      }
      let r = {id: this.$route.params.id, key: k}
      Store.dispatch('encrypt', r)
    }
  },
  created () {
    this.key = Store.state.key
  }
}

</script>

<style>

</style>
