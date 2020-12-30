<template>
  <q-page class="q-pa-md">
    <q-card class="posts my-card">
      <q-item>
        <q-item-section avatar>
          <q-avatar>
            <img src="https://cdn.quasar.dev/img/avatar2.jpg" />
          </q-avatar>
        </q-item-section>

        <q-item-section>
          <q-item-label>Title</q-item-label>
          <q-item-label caption>Subhead</q-item-label>
        </q-item-section>
      </q-item>

      <img src="https://cdn.quasar.dev/img/parallax2.jpg" />
    </q-card>
    <q-card class="posts my-card">
      <q-item>
        <q-item-section avatar>
          <q-avatar>
            <img src="https://cdn.quasar.dev/img/avatar2.jpg" />
          </q-avatar>
        </q-item-section>

        <q-item-section>
          <q-item-label>Title</q-item-label>
          <q-item-label caption>Subhead</q-item-label>
        </q-item-section>
      </q-item>

      <img src="https://cdn.quasar.dev/img/parallax2.jpg" />
    </q-card>
  </q-page>
</template>

<script>
export default {
  name: "HomePage",
  props: ["id"],
  data() {
    return {
      userId: "",
    };
  },
  beforeMount() {
    if (this.id == "") {
      this.$router.push({ name: "Login" });
    }
    this.userId = this.id;
    this.$axios
      .get(`http://localhost:3000/api/posts/${this.userId}`)
      .then((response) => {
        console.log(response.data);
      })
      .catch((error) => {
        this.errorMessage = error.message;
        console.error("There was an error!", error);
      });
  },
};
</script>

<style>
</style>
