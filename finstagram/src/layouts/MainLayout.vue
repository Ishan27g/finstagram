<template>
  <q-layout >
    
   
    <q-page-container class="bg-grey-1">
      <router-view />
    </q-page-container>


  </q-layout>
</template>

<script>
export default {
  name: "MainLayout",
  data() {
    return {};
  },
};
</script>

<style lang="sass">
.q-toolbar__title
  font-size: 34px
  @media (max-width: $breakpoint-xs-max)
    text-align: center
</style> 