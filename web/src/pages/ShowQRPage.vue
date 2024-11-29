<template>
  <div class="flex flex-col items-center scrollable px-4 lg:px-0">
    <div class="flex flex-col w-full lg:w-10/12 items-center gap-y-4 mb-4">
      <qrcode-vue :size="800" :value="value"></qrcode-vue>
    </div>
  </div>
</template>
<script lang="ts">
import QrcodeVue from 'qrcode.vue'
export default {
  components: {
    QrcodeVue,
  },
  data() {
    return {
      value: 'Test qr code',
      index: 1,
    }
  },
  mounted(){

    setInterval(() => {
      this.index++;
      this.value = `token=${this.index}`;
    }, 5000);
  },
};
</script>