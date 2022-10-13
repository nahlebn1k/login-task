function CreateUser(){
    event.preventDefault();
    let login = document.querySelector("#reg-login");
    let pass = document.querySelector("#reg-pass");
    let url=window.location.origin+"/signup";
    let data=JSON.stringify({"login":login.value, "password":pass.value});

    fetch(url, {
        method:'POST',
        headers: {
            'Content-Type': 'application/json; charset=UTF-8'
        },
        body: data
    }).then(response =>{
        if (!response.ok){
            throw new Error(`Http Error ${response.status}`);
        }
        return response.text();
    })
}

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
            throw new Error(`Http Error ${response.status}`);
        }
        return response.text();
    })
}