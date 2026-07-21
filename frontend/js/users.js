const API = "http://localhost:8080";
const token = localStorage.getItem("token");

const tbody = document.getElementById("tbody");
const modal = document.getElementById("userModal");
const form = document.getElementById("userForm");
const modalTitle = document.getElementById("modalTitle");
const submitBtn = document.getElementById("submitBtn");
const passwordGroup = document.getElementById("passwordGroup");
const retypePasswordGroup = document.getElementById("retypePasswordGroup");
const fullname = document.getElementById("fullname");
const email = document.getElementById("email");
const password = document.getElementById("password");
const retypePassword = document.getElementById("retypePassword");
const userId = document.getElementById("userId");

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
                <td class="action-td">
                    <button
                        class="btn btn-warning"
                        onclick="openEditModal(${user.id})">
                        Edit
                    </button>
                    <button
                        class="btn btn-danger"
                        onclick="deleteUser(${user.id})">
                        Delete
                    </button>
                </td>
            </tr>
        `;

    });

}


// =========================
// CREATE MODAL
// =========================

function openCreateModal() {
    form.reset();
    userId.value = "";
    modalTitle.innerHTML = "Tambah User";
    submitBtn.innerHTML = "Simpan";
    passwordGroup.style.display = "block";
    retypePasswordGroup.style.display = "block";
    password.required = true;
    modal.style.display = "block";
}

// =========================
// EDIT MODAL
// =========================

async function openEditModal(id) {

    try {
        const res = await fetch(`${API}/users/${id}`, {
            headers: {
                Authorization: token
            }
        });
        const data = await res.json();
        if (!data.success) {
            alert(data.message);
            return;
        }
        userId.value = data.results.id;
        fullname.value = data.results.fullname;
        email.value = data.results.email;
        password.value = "";
        retypePassword.value = "";
        modalTitle.innerHTML = "Edit User";
        submitBtn.innerHTML = "Update";
        passwordGroup.style.display = "none";
        retypePasswordGroup.style.display = "none";
        password.required = false;
        modal.style.display = "block";

    } catch (err) {
        console.error(err);
    }
}

// =========================
// CLOSE MODAL
// =========================

function closeModal() {
    form.reset();
    modal.style.display = "none";
}

// =========================
// SUBMIT FORM
// =========================

form.addEventListener("submit", async function (e) {
    e.preventDefault();
    const formData = new FormData(form);
    const id = formData.get("id");
    let url = `${API}/users`;
    let method = "POST";

    if (id == "") {
        if (password.value !== retypePassword.value) {
            alert("Password dan Retype Password tidak sama.");
            return;
        }
    } else {
        method = "PATCH";
        url += `/${id}`;
    }
    try {
        const res = await fetch(url, {
            method,
            headers: {
                Authorization: token
            },
            body: new URLSearchParams(formData)
        });
        const data = await res.json();
        alert(data.message);
        if (data.success) {
            closeModal();
            getUsers();
        }
    } catch (err) {
        console.error(err);
    }
});

// =========================
// DELETE
// =========================

async function deleteUser(id) {

    const confirmDelete = confirm("Yakin ingin menghapus user?");
    if (!confirmDelete) return;
    try {
        const res = await fetch(`${API}/users/${id}`, {
            method: "DELETE",
            headers: {
                Authorization: token
            }
        });
        const data = await res.json();
        alert(data.message);
        getUsers();
    } catch (err) {
        console.error(err);
    }
}

// =========================
// LOGOUT
// =========================

function logout() {
    localStorage.removeItem("token");
    window.location.href = "login.html";
}

// =========================
// CLOSE MODAL OUTSIDE
// =========================

window.onclick = function (event) {
    if (event.target == modal) {
        closeModal();
    }
}