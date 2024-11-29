<template>
  <div class="scrollable flex flex-col justify-center h-auto">
    <div class=" m-8 flex justify-center w-full h-auto text-3xl">
    Календарь мероприятий
    </div> 
    <!--прогресс бар по месяцам -->
    <div class="b-someclass">
      <div class="b-someclass-inner">
        <input type="radio" name="check" id="check-1" @change="showContainer = 'container1'">
        <label class="text-xl" for="check-1"><br><br>Июль</label>    
        
        <input type="radio" name="check" id="check-2" @change="showContainer = 'container2'">
        <label class="text-xl" for="check-2"><br><br>Август</label>
        
        <input type="radio" name="check" id="check-3" @change="showContainer = 'container3'">
        <label class="text-xl" for="check-3"><br><br>Сентябрь</label>
      
        <input type="radio" name="check" id="check-4" @change="showContainer = 'container4'">
        <label class="text-xl" for="check-4"><br><br>Октябрь</label>
        
        <input type="radio" name="check" id="check-5" @change="showContainer = 'container5'" checked> 
        <label class="text-xl" for="check-5"><br><br>Ноябрь</label>
        
        <input type="radio" name="check" id="check-6" @change="showContainer = 'container6'">
        <label class="text-xl" for="check-6"><br><br>Декабрь</label>
    
        <div class="b-someclass-line"></div>
      </div>     
    </div>
    <div v-if="showContainer==='container5'" class="flex flex-row w-2/5 h-auto " style="margin:50px">
      <div class="border-4 rounded-xl border-ev-color">
        <div class="flex flex-row">
          <img class="md:max-w-28 md:max-h-28 uus:max-w-24 uus:max-h-24 mr-4" src="../assets/icons/icon-company.svg">
          <div class="flex flex-col">
            <p class="font-bold md:text-2xl uus:text-xl">1C</p>
            <p class="md:text-xl uus:text-base">Младший разработчик</p>
            <div class="flex flex-col">
              <button class="sml-btn mb-2 mt-2">Принять</button>
              <button class="sml-btn">Отказаться</button>
            </div>
          </div>
        </div>
      </div>
      <div class="border-4 rounded-xl border-ev-color">
        <div class="flex flex-row">
          <img class="md:max-w-28 md:max-h-28 uus:max-w-24 uus:max-h-24 mr-4" src="../assets/icons/icon-company.svg">
          <div class="flex flex-col">
            <p class="font-bold md:text-2xl uus:text-xl">1C</p>
            <p class="md:text-xl uus:text-base">Младший разработчик</p>
            <div class="flex flex-col">
              <button class="sml-btn mb-2 mt-2">Принять</button>
              <button class="sml-btn">Отказаться</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script lang="ts">
import { useUserFormStore } from '@/stores/userFormStore';

export default {
  data() {
    return {
      showContainer: 'container5', 
    };
  },
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
<style scoped>
* {
  padding: 0;
  margin: 0;
}
.b-someclass {
  text-align: center;
}
.b-someclass-inner {
  display: inline-block;
  vertical-align: top;
  position: relative;
}
.b-someclass-inner:before {
  content: '';
  position: absolute;
  top: 50%;
  left: 15px;
  width: calc(100% - 30px);
  height: 6px;
  background: #C9C2F4;
  margin-top: -3px;
  z-index: 2;
}
.b-someclass-inner input {
  display: none;
}
.b-someclass-inner input + label {
  display: inline-block;
  vertical-align: top;
  margin: 35px;
  width: 50px;
  height: 50px;
  background: #C9C2F4;
  border-radius: 50%;
  position: relative;
  z-index: 9;
}
.b-someclass-inner input + label:after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  height: 50%;
  width: 50%;
  margin: -25% 0 0 -25%;
  border-radius: 50%;
  background: #52026A;
  transform: scale(0);
  transition: .25s;
  transition-delay: .2s;
}
.b-someclass-inner input:checked + label:after {
  background: #52026A;
  transform: scale(1);
}
.b-someclass-line {
  position: absolute;
  top: 50%;
  margin-top: -3px;
  left: 15px;
  width: 0;
  height: 6px;
  background: #52026A;
  transition: .35s;
  z-index: 2;
  pointer-events: none;
}
.b-someclass-inner input:checked + label ~ .b-someclass-line {
  left: 50px;
}
.b-someclass-inner #check-2:checked + label ~ .b-someclass-line {
  width: calc(100% / 6);
}
.b-someclass-inner #check-3:checked + label ~ .b-someclass-line {
  width: calc(100% / 6 * 2);
}
.b-someclass-inner #check-4:checked + label ~ .b-someclass-line {
  width: calc(100% / 6 * 3);
}
.b-someclass-inner #check-5:checked + label ~ .b-someclass-line {
  width: calc(100% / 6 * 4);
}
.b-someclass-inner #check-6:checked + label ~ .b-someclass-line {
  width: calc(100% / 6 * 5);
}
</style>
