import {defineStore} from 'pinia'
import {fetchWrapper} from '~/helper/wrapper.js'
import router from "#app/plugins/router.js";
import axios from "axios";
import data from "~/data.js";

const baseURL = `${import.meta.env.VITE_API_URL}/user`

export const useAuthStore = defineStore('auth',{
    states:()=>({
        user: JSON.parse(localStorage.getItem('user')),
        returnUrl: null,
        isLogin: false
    }),
    actions: {
      async login(username,password) {
          console.log(username,password)
        // const user = await fetchWrapper.post(`${baseURL}/login`,{username,password})
          const {data} = await axios.post(`${baseURL}/login`, {username:username, password:password})
        // this.user = user
          localStorage.setItem('user',JSON.stringify(data))
          this.isLogin = true
          console.log(this.isLogin)
          // router.push({path:this.returnUrl || "/"})
          window.location = "/"
      },
        async register(firstname,lastname,phone,email,username,password){
            const {data} = await axios.post(`${baseURL}/register`, {
                firstname:firstname,
                lastname:lastname,
                phone:phone,
                email:email,
                username:username,
                password:password,
            })
            // this.user = user
            localStorage.setItem('user',JSON.stringify(data))
            this.isLogin = true
            console.log(this.isLogin)
            // router.push({path:this.returnUrl || "/"})
            window.location = "/"
        },
        logout(){
            this.user = null
            localStorage.removeItem('user')
            router.push("/")
        }
    }
})