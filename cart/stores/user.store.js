import {defineStore} from "pinia";

const baseURL = `${import.meta.env.VITE_API_URL}/users`

export const useUserStore = defineStore('user',{

    states:()=>({
       user:{}
    }),
    actions: {
        async getAll(){
            this.user = {loading:true}
        }
    }
})