<template>
  <q-page class="screen q-pa-md">
    <h5>Home</h5>
    {{ $route.params.userId }}
  </q-page>
</template>

<script>
export default {
  name: "HomePage",
  data() {
    return {
      userId: "",
    };
  },
  onMounted() {
    if ($route.params.userId == "") {
      this.$router.push({ name: 'Login' });
    }
    this.userId = this.$route.params.userId
  },
};
</script>

<style>
</style>
