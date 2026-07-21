const form= document.getElementById("registerForm");

form.addEventListener("submit",async(e)=>{
    e.preventDefault();

    const formData=new FormData(form);
    const res = await fetch("http://localhost:8080/auth/register",{

        method:"POST",
        body: new URLSearchParams(formData)
    });

    const data = await res.json();

    alert(data.message);
    if(data.success){
        window.location="login.html";
    }

});