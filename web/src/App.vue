<template>
  <StatusWindow />
  <Header />
  <RouterView />
</template>

<script lang="ts">
import { mapStores } from "pinia";
import { useUserInfoStore } from "./stores/userInfoStore";
import { useUniversityStore } from "./stores/universityStore";

import StatusWindow from "./entities/statusWindow.vue";
import Header from "./entities/header.vue";

export default {
  components: {
    StatusWindow,
    Header,
  },
  computed:{
    ...mapStores(useUserInfoStore, useUniversityStore),
  },
  async mounted(){
    await this.userInfoStore.Authenticate();
    if(this.userInfoStore.authorized) {
      this.userInfoStore.loadUserData();
      this.universityStore.loadUniversityInfo(); 
    }
  },
};
</script>

<style v-global>

.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease-out;
}
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

.slowRemove-enter-active,
.slowRemove-leave-active{
  transition: all 100ms ease-out;
}

.slowRemove-enter-from,
.slowRemove-leave-to {
  opacity: 0;
}
</style>
