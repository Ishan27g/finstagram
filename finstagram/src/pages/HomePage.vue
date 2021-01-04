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
          :to="this.homepage"
        />
        <q-separator vertical inset spaced class="desktop-screen" />
        <q-btn
          flat
          round
          dense
          size="22px"
          color="primary"
          icon="eva-plus-square"
          class="desktop-screen"
          :to="this.addpage"
        />
      </q-toolbar>
    </q-header>
    <q-page class="q-pa-md home">
      <template v-if="!isloading">
        <div class="row justify-around q-gutter-y-lg q-col-gutter-xl">
          <div class="col-12 col-md-8">
            <q-card class="posts" v-for="post in posts" :key="post.Id">
              <q-item>
                <q-item-section avatar>
                  <q-avatar>
                    <img :src="post.User.Avatar" />
                  </q-avatar>
                </q-item-section>
                <q-item-section>
                  <q-item-label>{{ post.User.Username }}</q-item-label>
                  <q-item-label caption>{{ post.Date }}</q-item-label>
                </q-item-section>
              </q-item>
              <q-img :src="post.ImageUrl" class="images full-width"/>
              <q-item>
                <q-item-section>
                  <q-item-label>{{ post.Caption }}</q-item-label>
                </q-item-section>
              </q-item>
            </q-card>
          </div>
          <div class="desktop-screen col-md-3">
            <q-card bordered class="fixed">
              <div class="text-h5 text-center">Users</div>
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
    <q-footer class="bg-white mobile-screen screen" bordered>
      <q-tabs
        class="text-grey-10"
        active-color="primary"
        indicator-color="transparent"
      >
        <q-route-tab name="login" icon="eva-home" :to="this.homepage" />
        <q-route-tab
          name="signup"
          icon="eva-plus-square"
          :to="this.addpage"
        />
      </q-tabs>
    </q-footer>
  </q-layout>
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
      homepage: "",
      addpage: "",
    };
  },
  beforeMount() {
    if (this.id == "") {
      this.$router.push({ name: "Login" });
    }
    console.log(this.userId);

    this.userId = this.id;
    this.homepage =`/app/home/${this.userId}`
    this.addpage =`/app/addImage/${this.userId}`
  },
  mounted() {
    this.isloading = true;
    this.$axios
      .get(`http://localhost:3000/api/posts/${this.userId}`)
      .then((response) => {
        console.log("get posts -> ", response.data.Response)
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
        console.log("get users -> ", response.data)
        this.users = response.data.Response;
        this.isloadingusers = false;
      })
      .catch((error) => {
        this.errorMessage = error.message;
        console.error("There was an error!", error);
        this.isloadingusers = false;
      });
    console.log(this.userId);
  },
  methods: {
    getHomepage() {
      console.log(this.homepage);
      return this.homepage;
    },
  },
};
</script>

<style>
</style>
