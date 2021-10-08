import {createStore} from 'vuex'

const store =createStore({
    state :{
            isLogged:false,
            isClient:false,
    },
    mutations:{
        setLogedIn(state){
            state.isLogged = true;
        },
        setLogedOut(state){
            state.isLogged=false
        },
        setClient(state){
            state.isClient=true;
        },
        setServiceProvider(state){
            state.isClient=false;
        }
    }
});

export default store;