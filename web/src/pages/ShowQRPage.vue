<template>
  <div class="flex flex-col items-center scrollable px-4 lg:px-0">
    <div class="flex flex-col w-full lg:w-10/12 items-center gap-y-4 mb-4">
      <qrcode-vue :size="qrSize" :value="value"></qrcode-vue>
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
      qrSize: 100,
    }
  },
  mounted(){

    setInterval(() => {
      this.index++;
      this.value = `token=AUFHIUAsufH; ${this.index}`;
    }, 5000);

    this.handleWindowSize();
    window.addEventListener('resize', this.handleWindowSize());
  },
  methods: {
    handleWindowSize(){
      this.qrSize = Math.min(window.innerWidth, window.innerHeight);
    }
  }
};
</script>