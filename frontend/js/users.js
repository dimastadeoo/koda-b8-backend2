const API = "http://localhost:8080";
const token = localStorage.getItem("token");

const tbody = document.getElementById("tbody");
const form = document.getElementById("userForm");
const submitBtn = document.getElementById("submitBtn");

getUsers();
// =======================
// GET ALL
// =======================

async function getUsers() {
    const res = await fetch(`${API}/users`, {

        headers: {
            Authorization: token
        }

    });
    const data = await res.json();
    tbody.innerHTML = "";
    let i = 1

    if (!token) {
        alert(data.message + " Please Login First")
        location.href = "login.html";
    }       
    data.results.forEach(user => {
        tbody.innerHTML += `
            <tr>
                <td>${i++}</td>
                <td>${user.fullname}</td>
                <td>${user.email}</td>
                <td>${user.created_at}</td>
                <td>${user.updated_at}</td>
                <td>
                    <button onclick="editUser(${user.id})">
                        Edit
                    </button>
                    <button onclick="deleteUser(${user.id})">
                        Delete
                    </button>
                </td>
            </tr>
        `;

    });

}


// =======================
// SUBMIT
// =======================

form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const formData = new FormData(form);
    const id = formData.get("id");
    let url = `${API}/users`;
    let method = "POST";
    if (id) {
        url = `${API}/users/${id}`;
        method = "PATCH";
    }
    const res = await fetch(url, {
        method,
        headers: {
            Authorization: token
        },
        body: new URLSearchParams(formData)
    });
    const data = await res.json();
    alert(data.message);
    form.reset();
    document.getElementById("id").value = "";
    submitBtn.innerHTML = "Register";
    getUsers();

});


// =======================
// GET BY ID
// =======================

async function editUser(id) {
    const res = await fetch(`${API}/users/${id}`, {
        headers: {
            Authorization: token
        }

    });
    const data = await res.json();
    document.getElementById("id").value = data.results.id;
    document.querySelector("[name=fullname]").value =
        data.results.fullname;
    document.querySelector("[name=email]").value =
        data.results.email;
    document.querySelector("[name=password]").value = "";
    submitBtn.innerHTML = "Update User";

}


// =======================
// DELETE
// =======================

async function deleteUser(id) {
    const confirmDelete = confirm("Yakin ingin menghapus user?");
    if (!confirmDelete)
        return;
    const res = await fetch(`${API}/users/${id}`, {
        method: "DELETE",
        headers: {
            Authorization: token
        }

    });

    const data = await res.json();
    alert(data.message);
    getUsers();

}


// =======================
// CANCEL
// =======================

function cancelEdit() {
    form.reset();
    document.getElementById("id").value = "";
    submitBtn.innerHTML = "Register";
}


// =======================
// LOGOUT
// =======================

function logout() {
    localStorage.removeItem("token");
    location.href = "login.html";
}