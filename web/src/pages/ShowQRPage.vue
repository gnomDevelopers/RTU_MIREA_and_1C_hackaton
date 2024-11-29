<template>
  <div class="flex flex-col items-center scrollable px-4 lg:px-0">
    <div class="flex flex-col w-full h-full lg:w-10/12 items-center gap-y-4 mb-4" id="qrContainer">
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

      container: null as null | HTMLElement,
    }
  },
  mounted(){

    setInterval(() => {
      this.index++;
      this.value = `token=AUFHIUAsufH; ${this.index}`;
    }, 5000);

    this.container = document.getElementById('qrContainer');
    this.handleWindowSize();
    window.addEventListener('resize', this.handleWindowSize);
  },
  methods: {
    handleWindowSize(){
      const height = this.container ? this.container.offsetHeight : 300;
      const widht = this.container ? this.container.offsetWidth : 300;
      this.qrSize = Math.min(height, widht) - 20;
    }
  },
  unmounted(){
    window.removeEventListener('resize', this.handleWindowSize);
  }
};
</script>