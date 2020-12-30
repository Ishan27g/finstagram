<template>
  <q-page class="screen q-pa-md" style="max-width: 400px">
    <h5 class="text-center">Register</h5>

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
        type="email"
        v-model="email"
        label="Your email *"
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
      email: "",
      password: "",
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
        console.log("username: ", this.username);
        console.log("email: ", this.email);
        console.log("password: ", this.password);
        const article = {
          username: this.username,
          email: this.email,
          password: this.password,
        };
        axios
          .post("http://localhost:3000/api/app/signup", article)
          .then((response) => (this.articleId = response.data.id))
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