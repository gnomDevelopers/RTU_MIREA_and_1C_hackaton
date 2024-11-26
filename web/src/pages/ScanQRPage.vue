<template>
    <div class="flex flex-col items-center scrollable px-4 lg:px-0">
    <div class="flex flex-col w-full lg:w-10/12 items-center gap-y-4 mb-4">
      <QrcodeStream @detect="onDecode"/>
      <p>Статус: {{ status }} Расшифровано: {{ decoded }}</p>
    </div>
  </div>
</template>
<script lang="ts">
import { QrcodeStream } from 'vue-qrcode-reader';
export default {
  components: {
    QrcodeStream,
  },
  data(){
    return {
      decoded: undefined as any,
      status: 'waiting...',
    }
  },
  methods: {
    onDecode(decodedString: string){
      this.decoded = decodedString;
      this.status = 'decoded';
    },
  }
};
</script>