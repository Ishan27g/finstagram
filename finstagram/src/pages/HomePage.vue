<template>
  <q-page class="q-pa-md">
    <template v-if="!isloading">
      <div class="row justify-around q-col-gutter-xl">
        <div class="col-12 col-md-8">
          <q-card class="posts" v-for="post in posts" :key="post.Id">
            <q-item>
              <q-item-section avatar>
                <q-avatar>
                  <img src="https://cdn.quasar.dev/img/avatar2.jpg" />
                </q-avatar>
              </q-item-section>
              <q-item-section>
                <q-item-label>Username</q-item-label>
                <q-item-label caption>{{ post.Date }}</q-item-label>
              </q-item-section>
            </q-item>
            <q-img :src="post.ImageUrl" class="images full-width" />
            <q-item>
              <q-item-section>
                <q-item-label>{{ post.Caption }}</q-item-label>
              </q-item-section>
            </q-item>
          </q-card>
        </div>
        <div class="desktop-screen col-md-4">
          <q-card bordered class="fixed">
            <q-card-section
              horizontal
              v-for="user in users"
              :key="user.Username"
            >
              <q-card-section class="q-pt-xs">
                <q-item>
                  <q-item-section avatar>
                    <q-avatar size="60px">
                      <q-img :src="user.Avatar" />
                    </q-avatar>
                  </q-item-section>
                  <q-item-section>
                    <q-item-label>{{ user.Username }}</q-item-label>
                    <q-item-label caption>Last post at</q-item-label>
                    <q-item-label caption>31 Dec 20 01:52 AEST</q-item-label>
                  </q-item-section>
                </q-item>
              </q-card-section>
            </q-card-section>
          </q-card>
        </div>
      </div>
    </template>
    <template v-else>
      <q-card class="postsghost">
        <q-item>
          <q-item-section avatar>
            <q-skeleton type="QAvatar" animation="fade" />
          </q-item-section>
          <q-item-section>
            <q-item-label>
              <q-skeleton type="text" animation="fade" />
            </q-item-label>
            <q-item-label caption>
              <q-skeleton type="text" animation="fade" />
            </q-item-label>
          </q-item-section>
        </q-item>
        <q-skeleton min-height="300px" square animation="fade" />
        <q-card-section>
          <q-skeleton type="text" class="text-subtitle2" animation="fade" />
        </q-card-section>
      </q-card>
    </template>
  </q-page>
</template>

<script>
export default {
  name: "HomePage",
  props: ["id"],
  data() {
    return {
      userId: "",
      posts: [],
      isloading: false,
      isloadingusers: false,
      users: [],
    };
  },
  beforeMount() {
    if (this.id == "") {
      this.$router.push({ name: "Login" });
    }
    this.userId = this.id;
  },
  mounted() {
    this.isloading = true;
    this.$axios
      .get(`http://localhost:3000/api/posts/${this.userId}`)
      .then((response) => {
        this.posts = response.data.Response;
        this.isloading = false;
      })
      .catch((error) => {
        this.errorMessage = error.message;
        console.error("There was an error!", error);
        this.isloading = false;
      });
    this.$axios
      .get(`http://localhost:3000/api/users`)
      .then((response) => {
        this.users = response.data.Response;
        this.isloadingusers = false;
      })
      .catch((error) => {
        this.errorMessage = error.message;
        console.error("There was an error!", error);
        this.isloadingusers = false;
      });
  },
};
</script>

<style>
</style>
