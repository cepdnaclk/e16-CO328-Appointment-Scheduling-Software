import {createStore} from 'vuex'

const store =createStore({
    state :{
            isLogged:false,
            isClient:false,
            dashboardAlert:false
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
        },
        setDashBoardAlert(state){
            state.dashboardAlert=true
        },
        removeDashBoardAlert(state){
            state.dashboardAlert=false
        }
    }
});

export default store;