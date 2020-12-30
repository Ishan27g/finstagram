<template>
  <q-page class="screen q-pa-md" style="max-width: 400px">
    <h5 class="text-center">Login</h5>
    <q-form @submit="handleSubmit" class="q-gutter-md">
      <q-input
        filled
        v-model="username"
        label="Your username *"
        lazy-rules
        :rules="[(val) => (val && val.length > 0) || 'Please type something']"
      />
      <q-input
        filled
        type="password"
        v-model="password"
        label="Your password *"
        lazy-rules
        :rules="[(val) => (val && val.length > 0) || 'Please type something']"
      />

      <div>
        <q-btn label="Submit" type="submit" color="primary" />
      </div>
    </q-form>
  </q-page>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      username: "",
      password: "",
      userId: "",
    };
  },
  methods: {
    handleSubmit() {
      // validate password
      const passwordError =
        this.password.length > 5
          ? ""
          : "Password must be at least 6 characters long";
      if (!passwordError) {
        // make request to server to sign up user

        const article = {
          username: this.username,
          password: this.password,
        };
        axios
          .post("http://localhost:3000/api/login", article)
          .then((response) => {
            console.log(response.data)
            this.userId = response.data.Response;
            this.$router.push({path: `/app/home/${this.userId}`})
          })
          .catch((error) => {
            this.errorMessage = error.message;
            console.error("There was an error!", error);
          });
      }
    },
  },
};
</script>
<style lang="sass">
.q-form
  margin-top: 25%
  font-size: 34px
  text-align: center
</style> 