<template>
  <div>
    <h3>{{ word }}</h3>
    <p>=></p>
    <input v-for="(w, i) in word" :key="i">
  </div>
</template>

<script>

export default {
  props: {
    word: String
  }
}

</script>
