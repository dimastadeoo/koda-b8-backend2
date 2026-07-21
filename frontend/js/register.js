const API = "http://localhost:8080";

const form = document.getElementById("registerForm");
const password = document.getElementById("password");
const retypePassword = document.getElementById("retypePassword");

form.addEventListener("submit", async function (e) {
    e.preventDefault();
    if (password.value !== retypePassword.value) {
        alert("Password dan Retype Password tidak sama.");
        return;
    }
    const formData = new FormData(form);

    try {
        const res = await fetch(`${API}/auth/register`, {
            method: "POST",
            body: new URLSearchParams(formData)
        });
        const data = await res.json();
        alert(data.message);

        if (data.success) {
            window.location.href = "login.html";
        }
    } catch (err) {
        console.error(err);
        alert("Terjadi kesalahan.");

    }

});