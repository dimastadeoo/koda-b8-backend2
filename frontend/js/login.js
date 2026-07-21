const API = "http://localhost:8080";

const form = document.getElementById("loginForm");

form.addEventListener("submit", async function (e) {
    e.preventDefault();
    const formData = new FormData(form);
    try {

        const res = await fetch(`${API}/auth/login`, {
            method: "POST",
            body: new URLSearchParams(formData)
        });
        const data = await res.json();

        if (!data.success) {
            alert(data.message);
            return;
        }
        // Simpan token
        localStorage.setItem("token", data.results.token);
        alert(data.message);
        window.location.href = "users.html";

    } catch (err) {
        console.error(err);
        alert("Terjadi kesalahan.");
    }
});