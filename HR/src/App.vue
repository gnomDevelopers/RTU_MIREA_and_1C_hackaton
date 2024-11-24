<template>
  <StatusWindow />
  <RouterView />
</template>
<script lang="ts">
import StatusWindow from './entities/statusWindow.vue';
export default{
  components: {
    StatusWindow,
  }
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