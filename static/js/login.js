(function() {

    function sendRequest(formData) {
        // Create XHR object
        const xhr = new XMLHttpRequest();

        xhr.open("POST", "/post/login", true);
        xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');

        // Listen for the state change
        xhr.onreadystatechange = function() {
            if(this.readyState === XMLHttpRequest.DONE && this.status === 200) {
                // Read the server response
                let response = JSON.parse(this.responseText)
                if(response.status) {
                    location.href = "/";
                } else {
                    let errorCont = document.getElementById("error_cont");
                    let errorMsg = document.getElementById("error_msg");
                    errorMsg.innerText = response.message;
                    errorCont.classList.add("d-flex");
                    errorCont.classList.remove("d-none");
                }
            }
        }

        // Finally, send the request
        xhr.send(new URLSearchParams(formData).toString());
    }

    function hook() {
        let form = document.getElementById("login-form")
        form.onsubmit = function(e) {
            e.preventDefault();
            let formData = new FormData(form)
            sendRequest(formData);
        }
    }

    hook();
})()