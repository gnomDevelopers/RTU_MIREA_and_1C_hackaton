<template>
  <div class="p-2 text-lg bg-color-light rounded-xl cursor-pointer transition-colors border border-solid border-transparent hover:border-blue-900">
    {{ getName() }}
  </div>
</template>
<script lang="ts">
export default {
  props: {
    data: {
      type: Object,
      required: true,
    }
  },
  methods: {
    getName() {
      return `${this.data.surname === undefined ? '' : this.data.surname} ${this.data.name} ${this.data.thirdname === undefined ? '' : this.data.thirdname}`;
    },
  }
};
</script>