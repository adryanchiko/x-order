<template>
  <v-app-bar>
    <v-app-bar-nav-icon @click="toHome">
      <v-icon icon="mdi-basket" />
    </v-app-bar-nav-icon>
    <v-app-bar-title>
      X-Order
    </v-app-bar-title>
    <v-spacer></v-spacer>
    <v-btn @click="toOrders">
      Order
    </v-btn>
    <v-divider vertical class="mx-2 py-2" />
    <v-btn icon @click="onClick">
      <v-icon>
        {{appStore.theme === 'light' ? 'mdi-weather-sunny' : 'mdi-weather-night'}}
      </v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script lang="ts" setup>
import router from '@/router';
import { useAppStore } from '../../store/app'

const appStore = useAppStore()

function onClick() {
  let toggleTheme = appStore.theme === 'light' ? 'dark' : 'light'

  appStore.theme = toggleTheme
}

function toOrders() {
  router.push('/orders')
}
function toHome() {
  router.push('/')
}
</script>
