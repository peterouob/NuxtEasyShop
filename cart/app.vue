<template>
  <v-app id="inspire" :theme="useCartStore().getTheme">
<!--    Header-->
    <Header />
    <v-main>
      <v-container>
        <RouterView />
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>

</script>