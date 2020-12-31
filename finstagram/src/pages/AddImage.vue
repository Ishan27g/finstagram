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
      <template>
        <div class="row justify-around">
          <div class="col-9 q-pa-sm">
            <div class="q-gutter-y-md">
              <q-card class="addpost">
                <q-form class="addpost q-pa-md" style="text-align: center">
                  <q-input
                    v-model="caption"
                    label="Caption *"
                    lazy-rules
                    :rules="[
                      (val) =>
                        (val && val.length > 0) || 'Please type something',
                    ]"
                  />
                  <div>
                    <q-input
                      v-model="url"
                      label="url *"
                      lazy-rules
                      :rules="[
                        (val) =>
                          (val && val.length > 0) || 'Please type something',
                      ]"
                    />
                    <q-btn
                      label="Add"
                      class="text-center"
                      style="max-width: 10%; margin: auto"
                      @click="addImage()"
                    />
                  </div>
                </q-form>
              </q-card>
            </div>
          </div>
        </div>
      </template>
    </q-page>
    <q-footer class="bg-white mobile-screen screen" bordered>
      <q-tabs
        class="text-grey-10"
        active-color="primary"
        indicator-color="transparent"
      >
        <q-route-tab name="login" icon="eva-home" :to="this.homepage" />
        <q-route-tab name="signup" icon="eva-plus-square" :to="this.addpage" />
      </q-tabs>
    </q-footer>
  </q-layout>
</template>

<script>
export default {
  name: "AddImage",
  props: ["id"],
  data() {
    return {
      userId: "",
      homepage: "",
      addpage: "",
      text: "",
      newimage: "",
      caption: "",
      url: "",
    };
  },
  beforeMount() {
    if (this.id == "") {
      this.$router.push({ name: "Login" });
    }

    this.userId = this.id;
    console.log(this.userId);

    this.homepage = `/app/home/${this.userId}`;
    this.addpage = `/app/addImage/${this.userId}`;
  },
  methods: {
    addImage() {
      console.log(this.url);
      console.log(this.caption);
      const article = {
        caption: this.caption,
        imageUrl: this.url,
      };
      if(this.caption == "" || this.url == ""){
        return
      }
      this.$axios
        .post(`http://localhost:3000/api/posts/${this.userId}`, article)
        .then((response) => {
          console.log(response.data)
          this.$router.push({ path: `/app/home/${this.userId}` });
        })
        .catch((error) => {
          this.errorMessage = error.message;
          console.error("There was an error!", error);
        });
    },
  },
};
</script>
