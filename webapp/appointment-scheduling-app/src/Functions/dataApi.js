import axios from 'axios'

export  function getDataFromApiPOST(obj,url,data,doneTitle,doneText,pushLink) {
    let loader = obj.$loading.show();
      axios.post(url, data,{
          
          headers: {
            'Authorization': obj.$cookies.get("user_token")
          }
      }).then( (response) =>{
          loader.hide()
          obj.$notify({
          group: "done",
          title: doneTitle,
          text: doneText
          }, 4000);
        obj.$router.push(pushLink)
        console.log(response);
        return response.data
      }).catch(function (error) {
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
               
    });
}

export  function getDataFromApiGET(obj,url,doneTitle,doneText,pushLink) {
  let loader = obj.$loading.show();
    axios.get(url,{
        
        headers: {
          'Authorization': obj.$cookies.get("user_token")
        }
    }).then( (response) =>{
        loader.hide()
        obj.$notify({
        group: "done",
        title: doneTitle,
        text: doneText
        }, 4000);
      obj.$router.push(pushLink)
      console.log(response);
      return response.data
    }).catch(function (error) {
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
             
  });
}
