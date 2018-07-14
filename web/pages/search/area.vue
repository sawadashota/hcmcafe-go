<template>
  <section>
    <md-tabs class="md-accent" md-sync-route>
      <md-tab id="tab-home" md-label="Home" to="/components/tabs/home"/>
      <md-tab id="tab-map" md-label="Map" to="/components/tabs/pages"/>
    </md-tabs>
  </section>
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"

@Component({})
export default class extends Vue {
}
</script>

<style lang="scss" scoped>
</style>
