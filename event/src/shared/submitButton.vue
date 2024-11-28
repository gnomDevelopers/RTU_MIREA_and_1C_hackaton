<template>
  <input
    class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-xl font-semibold btn w-full text-slate-100"
    type="button"
    :value
  />
</template>
<script lang="ts">
export default {
  props: {
    value: {
      type: String,
      required: true,
    },
  },
};
</script>