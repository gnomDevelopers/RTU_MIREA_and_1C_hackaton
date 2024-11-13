<template>
  <StatusWindow />
  
  <Header />

  <RouterView />
</template>

<script lang="ts">
import { mapStores } from "pinia";
import { useStatusWindowStore, statusCodes } from "./stores/statusWindowStore";

import StatusWindow from "./entities/statusWindow.vue";
import Header from "./entities/header.vue";

export default {
  components: {
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
