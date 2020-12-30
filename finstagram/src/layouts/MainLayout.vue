<template>
  <q-layout view="lHh Lpr lFf" class="screen">
    <q-header class="bg-white screen" bordered>
      <q-toolbar>
        <q-toolbar-title class="text-primary text-instagram text-bold">
          Finstagram
        </q-toolbar-title>
        <q-btn
          flat
          round
          dense
          size="22px"
          color="primary"
          icon="eva-home"
          class="desktop-screen"
          to="/"
        />
        <q-separator vertical inset spaced class="desktop-screen"/>
        <q-btn
          flat
          round
          dense
          size="22px"
          color="primary"
          icon="eva-plus-square"
          class="desktop-screen"
          to="/addImage"
        />
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>

    <q-footer class="bg-white mobile-screen screen" bordered>
      <q-tabs
        class="text-grey-10"
        active-color="primary"
        indicator-color="transparent"
      >
        <q-route-tab name="home" icon="eva-home" to="/" />
        <q-route-tab name="add" icon="eva-plus-square" to="/addImage" />
      </q-tabs>
    </q-footer>
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