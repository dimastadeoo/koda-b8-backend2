const form=document.getElementById("loginForm");

form.addEventListener("submit",async(e)=>{
    e.preventDefault();

    const formData=new FormData(form);

    const res=await fetch("http://localhost:8080/auth/login",{

        method:"POST",
        body: new URLSearchParams(formData)
    });

    const data=await res.json();
    alert(data.message);

    if(data.success){
        localStorage.setItem("token", data.results.token)
        location.href="users.html";
    }

});