
function search() {
    var searchText = document.getElementById('SearchBar').value;
    console.log(searchText);
    var xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://127.0.0.1:3000/User', true);
    xhr.setRequestHeader('Accept', 'application/json');
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var data = JSON.parse(xhr.responseText);
            var container = document.getElementById('cardContainer');
            container.innerHTML = ''; // Clear the container before adding new cards
            var userFound = false; // Flag to track if any user was found
            data.forEach(obj => {
                if (obj.Username.toLowerCase().startsWith(searchText.toLowerCase())) {
                    var card = document.createElement('div');
                    card.className = 'col-sm-4 pt-3';
                    card.innerHTML = `
                        <div class="card">
                            <img src="./imgs/user.png" class="card-img-top" alt="...">
                            <div class="card-body">
                                <h5 class="card-title">${obj.Username}</h5>
                                <p class="card-text">${obj.Email}</p>
                                <a href="http://127.0.0.1:5500/Userdetail.html?ID=${obj.ID}" class="btn btn-primary">See Details</a>
                            </div>
                        </div>`;
                    container.appendChild(card);
                    userFound = true; // Set flag to true since a user was found
                }
            });
            if (!userFound) {
                getusers() // Display message if no user was found
            }
        }
    };
    xhr.send();
}

function getusers() {
    var searchText = document.getElementById('SearchBar').value;
    if (searchText==""){
        var xhr = new XMLHttpRequest();
        xhr.open('GET', 'http://127.0.0.1:3000/User', true);
        xhr.setRequestHeader('Accept', 'application/json');
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4 && xhr.status === 200) {
                var data = JSON.parse(xhr.responseText);
                var container = document.getElementById('cardContainer');

                // Clear existing cards


                // Create a card for each user
                data.forEach(obj => {
                    var card = document.createElement('div');
                    card.className = 'col-sm-4 pt-3';
                    card.innerHTML = `
                            <div class="card">
                                <img src="./imgs/user.png" class="card-img-top" alt="...">
                                <div class="card-body">
                                    <h5 class="card-title">${obj.Username}</h5>
                                    <p class="card-text">${obj.Email}</p>
                                    <a href="http://127.0.0.1:5500/Userdetail.html?ID=${obj.ID}" class="btn btn-primary">See Details</a>
                                </div>
                            </div>
                        `;
                    container.appendChild(card);
                });
            }
        }
        xhr.send();
    }
   

}
function query() {
    var urlParams = new URLSearchParams(window.location.search);
    var ID = urlParams.get('ID');
    console.log(ID);
    var xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://127.0.0.1:3000/Usersdetails?ID=' + ID, true);
    xhr.setRequestHeader('Accept', 'application/json');
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var data = JSON.parse(xhr.responseText);
            console.log(data);
            var container = document.querySelector('.row');
            if (container) {
                container.innerHTML = '';
                data.forEach(obj => {
                    var table = document.createElement('table');
                    table.className = 'table table-dark';
                    table.innerHTML = `
                                <thead>
                                    <tr>
                                        <th scope="col">ID</th>
                                        <th scope="col">Address</th>
                                        <th scope="col">Education</th>
                                        <th scope="col">Gmail</th>
                                        <th scope="col">Department</th>
                                        <th scope="col">Action</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr>
                                        <th scope="row">${obj.ID}</th>
                                        <td><small>${obj.Address}</small></td>
                                        <td><small>${obj.Education}</small></td>
                                        <td><small>${obj.Gmail}</small></td>
                                        <td><small>${obj.Department}</small></td>
                                        <td>
                                        <div class="d-flex justify-content-between">
                                    <button class="btn btn-success" style="width: 45%;">
                                        Edit
                                    </button>
                                    <button class="btn btn-danger text-center" style="width: 50%;">
                                        Delete
                                    </button>
                                </div>
                                        </td>
                                    </tr>
                                </tbody>`;
                    container.appendChild(table);
                });
            } else {
                console.error('Container element not found');
            }
        }
    };
    xhr.send();
}

function getemail() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://127.0.0.1:3000/Users?Username=Email', true);
    xhr.setRequestHeader('Accept', 'application/json');
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var obj = JSON.parse(xhr.responseText);
            console.log(obj);
            Object.keys(obj).forEach((key) => {
                document.getElementById('email').innerHTML += obj[key] + '<br>';
            });

        }
    }
    xhr.send();
}
function submitdata() {
    submitdata
    document.getElementById("myForm").addEventListener("submit", function (event) {
        event.preventDefault(); // Prevent form submission

        fetch('http://127.0.0.1:3000/Users', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                ID: document.getElementById('ID').value,
                Username: document.getElementById('Username').value,
                Email: document.getElementById('Email').value
            })
        })
            .then(response => response.json())
            .then(data => {
                alert("Success!");
                console.log(data);
            })
            .catch(error => console.error('Error:', error));
    });
}
function getusername() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://127.0.0.1:3000/Users?Username=Username', true);
    xhr.setRequestHeader('Accept', 'application/json');
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var obj = JSON.parse(xhr.responseText);
            console.log(obj);
            Object.keys(obj).forEach((key) => {
                document.getElementById('username').innerHTML += obj[key] + '<br>';
            });

        }
    }
    xhr.send();
}
