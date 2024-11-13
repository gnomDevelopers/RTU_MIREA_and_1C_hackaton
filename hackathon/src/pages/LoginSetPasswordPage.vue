<template>
  <div class="flex flex-grow flex-col gap-y-7 items-center justify-center bg-slate-50 login-bg">
    <div class="w-full text-center text-3xl mb:text-4xl">Вход в систему</div>
    <div class="flex flex-col gap-y-6 w-11/12 mb:w-96  py-5 px-7 rounded-lg login-shadow bg-slate-50 text-gray-900">
      <div class="w-full text-center text-4xl font-semibold">VUZ+</div>
      <div class="flex flex-col gap-y-5">
        <loginInput type="password" text="Пароль" @input-change="checkPassword"/>
        <loginInput type="password" text="Повторите пароль" @input-change="checkSecondPassword"/>
      </div>
      <submitButton value="Сохранить пароль" class="my-6"/>
    </div>
  </div>
</template>
<script lang="ts">

import loginInput from '../shared/loginInput.vue';
import submitButton from '../shared/submitButton.vue';
import { validUserLogin, validUserPassword } from '../helpers/validator';
import { type IValidAnswer } from '@/helpers/constants';

export default{
  components:{
    loginInput,
    submitButton,
  },
  data(){
    return{
      passwordValue: {value: '', error: ''} as IValidAnswer,
      isSamePassword: false as boolean,
    }
  },
  methods: {
    checkPassword(value: string){
      this.passwordValue = validUserPassword(value);
    },
    checkSecondPassword(value: string){
      if(this.passwordValue.value !== value) this.isSamePassword = false;
      else this.isSamePassword = true;
    }
  }
}

</script>