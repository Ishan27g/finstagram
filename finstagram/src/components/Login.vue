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
    <div class="row justify-center">
        <q-btn label="Go to signup page" to="/signup" color="secondary" />
      </div>
  </q-page>
</template>

<script>
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
        this.$axios
          .post("http://localhost:3000/api/login", article)
          .then((response) => {
            this.userId = response.data.Response;
            if(response.data.Code === 200){
              this.$router.push({ path: `/app/home/${this.userId}` });
            }
            else{
              this.$router.push({ path: '/signup' });  
            }
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