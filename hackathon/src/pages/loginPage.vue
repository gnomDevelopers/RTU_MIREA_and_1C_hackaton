<template>
  <div class="flex flex-grow flex-col items-center justify-center bg-slate-50">
    <div class="bg-slate-50 py-5 px-7 rounded-lg flex flex-col gap-y-6 login-shadow w-96 text-gray-900">
      <div class="w-full text-center text-2xl font-semibold">Log In</div>
      <div class="flex flex-col gap-y-5">
        <loginInput type="text" text="Login" @input-change="checkLogin"/>
        <loginInput type="password" text="Password" @input-change="checkPassword"/>
        <div class=" text-sm text-sky-500 w-full text-right cursor-pointer">Forgot Password?</div>
      </div>
      <submitButton value="Log In" class="mt-6"/>
      <div class="w-full text-center text-gray-400 text-sm" >or Log In with:</div>
      <div class="flex flex-row w-full justify-center gap-x-4">
        <div><img class="w-10 h-10 cursor-pointer rounded-full" src="../assets/icons/icon-facebook.svg"/></div>
        <div><img class="w-10 h-10 cursor-pointer rounded-full" src="../assets/icons/icon-vk.svg" /></div>
        <div><img class="w-10 h-10 cursor-pointer rounded-full" src="../assets/icons/icon-google.svg"></div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">

import loginInput from '../shared/loginInput.vue';
import submitButton from '../shared/submitButton.vue';
import { validUserLogin, validUserPassword, type IValidAnswer } from '../helpers/validator';

export default{
  components:{
    loginInput,
    submitButton,
  },
  data(){
    return{
      loginValue: {value: '', error: ''} as IValidAnswer,
      passwordValue: {value: '', error: ''} as IValidAnswer,
    }
  },
  methods: {
    checkLogin(value: string){
      this.loginValue = validUserLogin(value);
    },
    checkPassword(value: string){
      this.passwordValue = validUserPassword(value);
    }
  }
}

</script>