<template>
  <div>
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
        <td v-for="(k, kk, i) in mykey" :key="i" :kk="kk" :k="k" :i="i"><input @input="change" v-model="mykey[kk]" v-bind:name="k" class="w" maxlength="1"/></td>
      </tr>
    </table>
    <div v-show="score >= 13">
      <p>変更していないワード{{ noChange }}</p>
      <p>復号結果にないワード{{ noDec }}</p>
    </div>
  </div>
</template>

<script>
import Store from '../../stores/index'

export default {
  data () {
    return {
      mykey: {'A': 'A', 'B': 'B', 'C': 'C', 'D': 'D', 'E': 'E', 'F': 'F', 'G': 'G', 'H': 'H', 'I': 'I', 'J': 'J', 'K': 'K', 'L': 'L', 'M': 'M', 'N': 'N', 'O': 'O', 'P': 'P', 'Q': 'Q', 'R': 'R', 'S': 'S', 'T': 'T', 'U': 'U', 'V': 'V', 'W': 'W', 'X': 'X', 'Y': 'Y', 'Z': 'Z'},
      keykey: ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
    }
  },
  computed: {
    score () {
      return Store.state.count
    },
    key () {
      return Store.state.key
    },
    noChange () {
      let k = []
      for (let i = 0; i < this.keykey.length; i++) {
        if (this.mykey[this.keykey[i]] === this.keykey[i]) {
          k.push(this.keykey[i])
        }
      }
      return k
    },
    noDec () {
      let k = []
      let no = []
      Object.keys(this.mykey).forEach(function (key) {
        k.push(this.mykey[key])
      }, this)
      this.keykey.concat(k)
        .forEach(item => {
          if (this.keykey.includes(item) && !k.includes(item)) {
            no.push(item)
          }
        }, this)
      return no
    }
  },
  methods: {
    change () {
      for (let i = 0; i < this.keykey.length; i++) {
        let mk = this.mykey[this.keykey[i]]
        if (mk === '' || mk === undefined) return
      }
      let r = {id: this.$route.params.id, key: this.mykey}
      Store.dispatch('encrypt', r)
    }
  },
  created () {
    this.mykey = Store.state.key
  },
  watch: {
    key (v) {
      this.mykey = v
    }
  }
}

</script>

<style>

</style>
