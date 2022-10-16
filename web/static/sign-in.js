function SignIn() {
    event.preventDefault();
    let login = document.querySelector("#log-login");
    let pass = document.querySelector("#log-pass");
    let url=window.location.origin+"/signin";
    let data=JSON.stringify({"login":login.value, "password":pass.value});


    fetch(url, {
        method:'POST',
        headers: {
            'Content-Type': 'application/json; charset=UTF-8'
        },
        body: data
    }).then(response =>{
        if (!response.ok){
            throw new Error("HTTP status" + response.status);
        }
        return response.json();
    }).then(object => {
        window.localStorage.setItem("token", object.token)
        window.location.href = 'logout.html'
    })

}