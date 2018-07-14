<template>
  <div class="card">
    <h2>title: {{ title }}</h2>
  </div>
</template>
<script lang="ts">
  import {
    Component,
    Prop,
    Vue
  } from "nuxt-property-decorator"

  @Component({})
  export default class Hoge extends Vue {
    @Prop() title
  }
</script>
<style scoped>
  .card {
    font-family: "Segoe UI", Tahoma, Geneva, Verdana,
    sans-serif;
    padding: 1rem;
    margin: 0.25rem;
    border: 0.25rem solid gainsboro;
  }
</style>
