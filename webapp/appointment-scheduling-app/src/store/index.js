import {createStore} from 'vuex'

const store =createStore({
    state :{
            isLogged:false,
            isClient:false,
            dashboardAlert:false,
            appoimentModal:false,
            dashboardAlertExpired:false
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
        },
        setAppoimentModal(state){
            state.appoimentModal=true
        },
        removeAppoimentModal(state){
            state.appoimentModal=false
        },
        setDashBoardExpiredAlert(state){
            state.dashboardAlertExpired=true
        },
        removeDashBoardExpiredAlert(state){
            state.dashboardAlertExpired=false
        },
        
    }
});

export default store;