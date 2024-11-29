<template>
  <div class="scrollable flex h-auto">
    <div class="flex justify-center w-full">
      <p class="text-3xl m-8">Моя команда</p>
    </div>

  </div>

</template>

<script lang="ts">
import { useUserFormStore } from '@/stores/userFormStore';

export default {
  setup() {
    const userFormStore = useUserFormStore();
    return {
      formData: userFormStore.formData,
    };
  },
  computed: {
    getDownloadLink() {
      if (this.formData.filesList.length > 0) {
        return URL.createObjectURL(this.formData.filesList[0]);
      }
      return '';
    },
  },
  methods: {
    hideProfile() {
      if (!this.formData.isClosed) {
        this.formData.isClosed = true;
      }
      else this.formData.isClosed = false;
    }
  }
};
</script>