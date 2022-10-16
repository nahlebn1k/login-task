function TestAuth(){
    let data="Bearer "+window.localStorage.getItem("token");
    let url=window.location.origin+"/api";
    let text=document.querySelector('.welc')
    let res;
    fetch (url,{
        method: "POST",
        headers: {
            'Content-Type': 'application/json; charset=UTF-8',
            'Authorization': data
        }
    }).then(response =>{
        if (!response.ok){
            throw new Error("HTTP status" + response.status);
        }
        return response.text();
    }).then(obj=>text.innerHTML=`<h1>WELCOME USER #${obj}</h1>`)
}

window.addEventListener('load', TestAuth)

function Logout(){
    window.localStorage.clear();
    window.location.href='index.html';
}