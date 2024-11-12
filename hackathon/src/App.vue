<template>
  <StatusWindow />
  <Header />
  <SignUpPage />
</template>
<script lang="ts">
import { mapStores } from "pinia";
import { useStatusWindowStore, statusCodes } from "./stores/statusWindowStore";

import LoginPage from "./pages/loginPage.vue";
import SignUpPage from "./pages/SignUpPage.vue";
import StatusWindow from "./entities/statusWindow.vue";
import Header from "./entities/header.vue";

export default {
  components: {
    LoginPage,
    SignUpPage,
    StatusWindow,
    Header,
  },
  data() {
    return {

    };
  },
  computed: {
    ...mapStores(useStatusWindowStore),
  },
  methods: {

  },
};
</script>
