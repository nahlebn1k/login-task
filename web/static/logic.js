function CreateUser(){
    let login = document.querySelector("#user-login");
    let pass = document.querySelector("#user-pass");
    let url=window.location.origin+"/sign-up";
    let data=JSON.stringify({"login":login.value, "pass":pass.value});

    fetch(url, {
        method:'POST',
        headers: {
            'Content-Type': 'application/json; charset=UTF-8'
        },
        body: data
    }).then(response =>{
        if (!response.ok){
            throw new Error("Http Error");
        }
        return response.text();
    })
}