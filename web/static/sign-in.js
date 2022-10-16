function SignIn() {
    event.preventDefault();
    let login = document.querySelector("#log-login");
    let pass = document.querySelector("#log-pass");
    let url=window.location.origin+"/signin";
    let chbox=document.querySelector('#remember-me')
    let data=JSON.stringify({"login":login.value, "password":pass.value});

    if ((login.value).length===0 || (pass.value).length===0){
        alert("Empty input");
        return;
    }
    if (validateEmail(login.value)===false){
        alert("Incorrect email");
        return;
    }

    if (chbox.checked){
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
            window.localStorage.setItem("accessToken", object.accessToken)
            window.localStorage.setItem("refreshToken", object.refreshToken)
            window.localStorage.setItem("id", object.id)
            window.location.href = 'user.html'
        })
    } else{
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
            window.localStorage.setItem("accessToken", object.accessToken)
            //window.localStorage.setItem("refreshToken", object.refreshToken)
            window.localStorage.setItem("id", object.id)
            window.location.href = 'user.html'
        })
    }

}

function validateEmail(email) {
    let re = /\S+@\S+\.\S+/;
    return re.test(email);
}