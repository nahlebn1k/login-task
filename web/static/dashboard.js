function TestAuth(){
    let access="Bearer "+window.localStorage.getItem("accessToken");
    let refresh="Bearer "+window.localStorage.getItem("refreshToken");
    let user = window.localStorage.getItem("id");
    let url=window.location.origin+"/user";
    let text=document.querySelector('.welc')
    fetch (url,{
        method: "POST",
        headers: {
            'Content-Type': 'application/json; charset=UTF-8',
            'Authorization': access,
            'User' : user
        }
    }).then(response =>{
        if (!response.ok){
            fetch (url,{
                method: "POST",
                headers: {
                    'Content-Type': 'application/json; charset=UTF-8',
                    'Authorization': refresh,
                    'User':user
                }
            }).then(response =>{
                if (!response.ok){
                    Logout()
                    throw new Error("HTTP status" + response.status);
                }
                return response.text();
            }).then(obj=>{
                Refresh();
                text.innerHTML=`<h1>WELCOME USER #${obj}</h1>`;
                //Refresh();
            })
        }
        return response.text();
    }).then(obj=>text.innerHTML=`<h1>WELCOME USER #${obj}</h1>`)
}

window.addEventListener('load', TestAuth)

function Logout(){
    window.localStorage.clear();
    window.location.href='login.html';
}

function Refresh() {
    let user = window.localStorage.getItem("id");
    let url = window.location.origin + "/refresh";
    fetch(url, {
        method: "POST",
        headers: {
            'Content-Type': 'application/json; charset=UTF-8',
            'User': user
        }
    }).then(response => {
        if (!response.ok) {
            throw new Error("HTTP status" + response.status);
        }
        return response.json();
    }).then(object=>{
        window.localStorage.setItem("accessToken", object.accessToken)
        window.localStorage.setItem("refreshToken", object.refreshToken)
    })
}