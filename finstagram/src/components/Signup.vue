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
      <q-input
        filled
        v-model="avatar"
        label="Your Avatar url *"
        lazy-rules
        :rules="[(val) => (val && val.length > 0) || 'Please type something']"
      />

      <div>
        <q-btn label="Submit" type="submit" color="primary" />
      </div>
    </q-form>
      <div class="row justify-center">
        <q-btn label="Go to login page" to="/" color="secondary" />
      </div>
  </q-page>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      email: "",
      password: "",
      avatar: "",
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
          email: this.email,
          password: this.password,
          avatar: this.avatar,
        };
        console.log('what')
        this.$axios
          .post("http://localhost:3000/api/signup", article)
          .then((response) => {
            console.log(response.data)
            this.articleId = response.data.id;
            this.$router.push('/')
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