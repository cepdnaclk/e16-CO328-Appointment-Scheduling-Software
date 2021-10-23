import axios from 'axios'

export  function getDataFromApiPOST(obj,url,data,doneTitle,doneText,pushLink,redirect,doneNotification) {
    return new Promise((resolve,reject)=>{
      let loader = obj.$loading.show();
      axios.post(url, data,{
          
          headers: {
            'Authorization': obj.$cookies.get("user_token")
          }
      }).then( (response) =>{
          loader.hide()
          if (doneNotification) {
            obj.$notify({
              group: "done",
              title: doneTitle,
              text: doneText
              }, 4000); 
          }
          if (redirect) {
            obj.$router.push(pushLink)
          }else{
            return resolve(response.data)
          }
         
      }).catch(error=> {
         loader.hide()
          if (error.response && error.response.status==403) {
              obj.$cookies.remove("user_token")
              obj.$cookies.remove("user_type")
              obj.$notify({
                group: "generic",
                title: "Session Expires ",
                text: "Session Expired pleace login"
                }, 4000)
                obj.$store.commit('setLogedOut')
                obj.$router.push('/')
              } else {
                obj.$notify({
                  group: "error",
                  title: "Error",
                  text: "Something happened in setting up the request that triggered a Error!",
                }, 4000)
            }
            reject(error) 
               
    });
    })
}

export  function getDataFromApiGET(obj,url,doneTitle,doneText,pushLink,redirect,doneNotification) {
  return new Promise((resolve,reject)=>{
    let loader = obj.$loading.show();
    axios.get(url,{
        headers: {
          'Authorization': obj.$cookies.get("user_token")
        }
    }).then( (response) =>{
        loader.hide()
        if (doneNotification) {
          obj.$notify({
          group: "done",
          title: doneTitle,
          text: doneText
          }, 4000); 
        }
        
        if (redirect) {
          obj.$router.push(pushLink)
        }else{
          return resolve(response.data)
        }
    }).catch(error=> {
       loader.hide()
        if (error.response && error.response.status==403) {
            obj.$cookies.remove("user_token")
            obj.$cookies.remove("user_type")
            obj.$notify({
              group: "generic",
              title: "Session Expires ",
                text: "Session Expired pleace login"
              }, 4000)
              obj.$store.commit('setLogedOut')
              obj.$router.push('/')
            } else {
              obj.$notify({
                group: "error",
                title: "Error",
                text: "Something happened in setting up the request that triggered a Error!",
              }, 4000)
          }
          reject(error)       
  });
 }) 
}
