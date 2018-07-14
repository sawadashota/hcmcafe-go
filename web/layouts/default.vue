<template>
  <div class="page-container">
    <md-app md-waterfall md-mode="fixed-last">
      <md-app-toolbar class="md-dense md-primary">
        <div class="md-toolbar-row">
          <div class="md-toolbar-section-start">
            <nuxt-link :to="{name: 'index'}" class="md-title logo">{{ $t('title') }}</nuxt-link>
          </div>

          <div class="md-toolbar-section-end sp">
            <md-button class="md-icon-button" @click="switchMenuVisible">
              <md-icon>menu</md-icon>
            </md-button>
          </div>
        </div>
      </md-app-toolbar>

      <md-app-drawer :md-active.sync="menuVisible" md-permanent="clipped">
        <md-list>
          <md-list-item :to="{name: 'index'}">
            <md-icon>home</md-icon>
            <span class="md-list-item-text">{{ $t('nav.home') }}</span>
          </md-list-item>

          <md-list-item :to="{name: 'search-gps'}">
            <md-icon>near_me</md-icon>
            <span class="md-list-item-text">{{ $t('nav.near_me') }}</span>
          </md-list-item>

          <md-list-item :to="{name: 'search-feeling'}">
            <md-icon>favorite</md-icon>
            <span class="md-list-item-text">{{ $t('nav.feeling') }}</span>
          </md-list-item>

          <md-list-item :to="{name: 'search-area'}">
            <md-icon>trip_origin</md-icon>
            <span class="md-list-item-text">{{ $t('nav.area_search') }}</span>
          </md-list-item>

          <md-list-item :to="{name: 'features'}">
            <md-icon>star</md-icon>
            <span class="md-list-item-text">{{ $t('nav.features') }}</span>
          </md-list-item>

          <md-divider/>

          <md-list-item :href="$t('facebook_link')" target='_blank' rel='noopener'>
            <md-icon>thumb_up_alt</md-icon>
            <span class="md-list-item-text">{{ $t('facebook') }}</span>
          </md-list-item>
        </md-list>
      </md-app-drawer>

      <md-app-content>
        <nuxt/>
      </md-app-content>
    </md-app>
  </div>
</template>

<script lang="ts">
  import {
    Component,
    Vue
  } from "nuxt-property-decorator"

  @Component({})
  export default class extends Vue {
    menuVisible: boolean = false

    /**
     * Switch Menu Visible
     */
    switchMenuVisible(): void {
      this.menuVisible = !this.menuVisible
    }
  }
</script>

<style lang="scss" scoped>
  .md-app {
    border: 1px solid rgba(#000, .12);
  }

  .logo {
    &:hover {
      text-decoration: none;
    }
  }

  .md-drawer {
    background-color: #fff;
    width: 230px;
    max-width: calc(100vw - 125px);
  }
</style>
