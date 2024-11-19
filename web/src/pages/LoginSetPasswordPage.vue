<template>
  <div class="flex flex-grow flex-col gap-y-7 items-center justify-center bg-slate-50 login-bg">
    <div class="w-full text-center text-3xl mb:text-4xl">Смена временного пароля</div>
    <div class="flex flex-col gap-y-6 w-11/12 mb:w-96  py-5 px-7 rounded-lg login-shadow bg-slate-50 text-gray-900">
      <div class="logo" style="font-family: BacasimeAntique">VUZ+</div>
      <div class="flex flex-col gap-y-5">
        <loginInput type="password" text="Пароль" @input-change="checkPassword"/>
        <loginInput type="password" text="Повторите пароль" @input-change="checkSecondPassword" ref="repeatPassword"/>
      </div>
      <submitButton value="Сохранить пароль" class="btn" @click="sendPasswords"/>
    </div>
  </div>
</template>
<script lang="ts">

import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { ValidUserPassword } from '../helpers/validator';
import { type IValidAnswer, StatusCodes } from '@/helpers/constants';

import loginInput from '../shared/loginInput.vue';
import submitButton from '../shared/submitButton.vue';

export default{
  components:{
    loginInput,
    submitButton,
  },
  data(){
    return{
      password: {value: '', error: ''} as IValidAnswer,
      secondPassword: {value: '', error: ''} as IValidAnswer,
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore),
  },
  methods: {
    sendPasswords(){
      if(this.password.value !== '' && this.secondPassword.value !== '' && this.password.value === this.secondPassword.value){
        //API
        return;
      }

      if(this.password.value === ''){
        if(this.password.error === '') this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите новый пароль!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.password.error);
      }
      if(this.secondPassword.value === ''){
        if(this.secondPassword.error === '') this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Повторите пароль!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.password.error);
      }
      else{
        if(this.password.value !== '') this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Пароли не совпадают!');
      }
    },
    checkPassword(value: string){
      this.password = ValidUserPassword(value);
      if(value === '') this.password.error = '';
    },
    checkSecondPassword(value: string){
      this.secondPassword = ValidUserPassword(value);
      if(value === '') this.secondPassword.error = '';
    }
  }
}

</script>