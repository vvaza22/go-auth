(function() {
    function showErrors(errors) {
        console.log(errors);

        let valid = ["first_name", "last_name", "username", "password"];

        for(let i=0; i<errors.length; i++) {
            let error = errors[i];
            valid = valid.filter(elem => error.element !== elem);
            let elem = document.getElementById(error.element);
            console.log(error.element + "_invalid");
            let elemError = document.getElementById(error.element + "_invalid");
            elemError.innerText = error.message;
            elem.className = "form-control is-invalid";
        }

        for(let i=0; i<valid.length; i++) {
            let node = document.getElementById(valid[i]);
            node.className = "form-control is-valid";
        }

    }

    function sendRequest(formData) {
        // Create XHR object
        const xhr = new XMLHttpRequest();

        xhr.open("POST", "/post/register", true);

        xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');

        // Listen for the state change
        xhr.onreadystatechange = function() {
            if(this.readyState === XMLHttpRequest.DONE && this.status === 200) {
                // Read the server response
                let response = JSON.parse(this.responseText)
                if(response.status) {
                    location.href = "/";
                } else {
                    showErrors(response.errors);
                }
            }
        }

        // Finally, send the request
        xhr.send(new URLSearchParams(formData).toString());
    }

    function hook() {
        const form = document.getElementById("register-form");
        form.onsubmit = function (e) {
            e.preventDefault();
            let formData = new FormData(this);
            sendRequest(formData)
        }
        const inputs = document.getElementsByClassName("form-control");
        for(let i=0; i<inputs.length; i++) {
            inputs[i].onchange = function() {
                // Reset error warnings
                this.className = "form-control";
            }
        }
    }

    hook();
})()