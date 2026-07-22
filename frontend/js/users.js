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
const pictureInput = document.getElementById("picture");
const groupPicture = document.getElementById("groupPicture");
const preview = document.getElementById("previewPicture");
groupPicture

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
    
    if (res.status === 401 || res.status === 400) {
        localStorage.removeItem("token");
        alert("Please Login First")
        location.href = "login.html";
    }
    
    const data = await res.json();
    tbody.innerHTML = "";
    let i = 1

    data.results.forEach(user => {
        const picture = user.picture 
            ? `${API}/${user.picture}`
            : `${API}/default-image.jpeg`
        tbody.innerHTML += `
            <tr>
                <td>${i++}</td>
                <td>
                    <img src="${picture}" class="avatar"
                    alt="Picture-Not Found">
                </td>
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
    groupPicture.style.display = "none"
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
        const user = data.results;
        userId.value = user.id;
        fullname.value = user.fullname;
        email.value = user.email;

        password.value = "";
        retypePassword.value = "";
        groupPicture.style.display = "block"
        passwordGroup.style.display = "none";
        retypePasswordGroup.style.display = "none";
        password.required = false;

        modalTitle.innerHTML = "Edit User";
        submitBtn.innerHTML = "Update";

        preview.src = user.picture
            ? `${API}/${user.picture}`
            : "img/default-avatar.png";

        pictureInput.value = "";

        modal.style.display = "block";
    } catch (err) {

        console.error(err);

        alert("Gagal mengambil data user.");
    }
}

async function uploadPicture(id) {

    if (pictureInput.files.length === 0) {
        return true;
    }

    const formData = new FormData();
    formData.append(
        "picture",
        pictureInput.files[0]
    );

    try {

        const res = await fetch(
            `${API}/users/${id}/picture`,
            {
                method: "PATCH",
                headers: {
                    Authorization: token
                },
                body: formData
            }
        );
        const data = await res.json();

        if (!data.success) {
            alert(data.message);
            return false;
        }
        return true;

    } catch (err) {
        console.error(err);
        alert("Gagal upload picture.");
        return false;
    }

}


pictureInput.addEventListener("change", function () {
    const file = this.files[0];
    if (!file) return;
    preview.src = URL.createObjectURL(file);

});

// =========================
// CLOSE MODAL
// =========================

function closeModal() {
    form.reset();
    modal.style.display = "none";
    groupPicture.style.display = "none"
    preview.src = ""
    pictureInput.value = ""
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

            if (method === "PATCH") {
                const success = await uploadPicture(id);
                if (!success) return;
            }
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