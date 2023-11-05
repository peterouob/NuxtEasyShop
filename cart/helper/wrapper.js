import {useAuthStore} from "~/stores/auth.store.js";

export const fetchWrapper = {
    get : request('GET'),
    post : request('POST')
};

function request(method){
    return async (url, body) => {
        const requestOptions = {
            method,
            header: authHeader(url)
        };
        if (body) {
            requestOptions.header['Content-type'] = 'application/json'
            requestOptions.body = JSON.stringify(body)
        }else{
            console.log("no body")
        }
        const result = await fetch(url, requestOptions);
        return handleResponse(result);
    }
}

function authHeader(url){
    const {user} = useAuthStore()
    const isLoggedIn = !!user?.token
    const isApiUrl = url.startsWith(import.meta.env.VITE_API_URL)
    if(isLoggedIn && isApiUrl){
        return {Auth: `Bearer ${user.token}`}
    }else{
        return {}
    }
}

function handleResponse(response){
    return response.text().then(text=>{
        const data = text && JSON.parse(response)
        if(!response.ok){
            const {user,logout} = useAuthStore()
            if([401,403].includes(response.status) && user){
                logout()
            }
            const error = (data && data.message ||response.statusText)
            return Promise.reject(error)
        }
        return data
    })
}
