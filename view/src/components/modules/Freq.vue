<template>
  <div>
    <h3>{{ word }}</h3>
    <p>=></p>
    <input v-for="(w, i) in word" :key="i" @input="change" v-model="mykey[w]" v-bind:name="w" maxlength="1">
  </div>
</template>

<script>
import Store from '../../stores/index'

export default {
  props: {
    word: String
  },
  data () {
    return {
      mykey: {'A': 'A', 'B': 'B', 'C': 'C', 'D': 'D', 'E': 'E', 'F': 'F', 'G': 'G', 'H': 'H', 'I': 'I', 'J': 'J', 'K': 'K', 'L': 'L', 'M': 'M', 'N': 'N', 'O': 'O', 'P': 'P', 'Q': 'Q', 'R': 'R', 'S': 'S', 'T': 'T', 'U': 'U', 'V': 'V', 'W': 'W', 'X': 'X', 'Y': 'Y', 'Z': 'Z'},
      keykey: ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
    }
  },
  computed: {
    key () {
      return Store.state.key
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
