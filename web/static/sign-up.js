function CreateUser(){
    event.preventDefault();
    let login = document.querySelector("#reg-login");
    let pass = document.querySelector("#reg-pass");
    let url=window.location.origin+"/signup";
    let data=JSON.stringify({"login":login.value, "password":pass.value});

    if ((login.value).length===0 || (pass.value).length===0){
        alert("Empty input");
        return;
    }
    if (validateEmail(login.value)===false){
        alert("Incorrect email");
        return;
    }

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

function validateEmail(email) {
    let re = /\S+@\S+\.\S+/;
    return re.test(email);
}