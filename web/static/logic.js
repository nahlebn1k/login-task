function CreateUser(){
    event.preventDefault();
    let login = document.querySelector("#user-login");
    let pass = document.querySelector("#user-pass");
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