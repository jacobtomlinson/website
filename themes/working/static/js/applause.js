function updateApplause(applause) {
    let container = document.getElementById('applause');
    if (!container) {
        const attributes = document.getElementsByClassName("attributes")[0];
        attributes.innerHTML += '<small id="applause"></small>';
        document.getElementById("applause").addEventListener("click", giveApplause);
    }
    container = document.getElementById('applause');
    container.innerHTML = '<i class="fas fa-hands-clapping"></i> ' + applause;
}

function getApplause() {
    var xhr = new XMLHttpRequest();
    var url = 'https://applause.cloud.jacobtomlinson.dev' + window.location.pathname;
    xhr.open("GET", url, true);
    xhr.onreadystatechange = function () {
        if (this.readyState == 4 && this.status == 200) {
            updateApplause(this.responseText);
        }
    }
    xhr.send();
}

function giveApplause() {
    var xhr = new XMLHttpRequest();
    var url = 'https://applause.cloud.jacobtomlinson.dev' + window.location.pathname;
    xhr.open("POST", url, true);
    xhr.onreadystatechange = function () {
        if (this.readyState == 4 && this.status == 200) {
            updateApplause(this.responseText);
        }
    }
    xhr.send();
}

window.addEventListener("load", function () {
    getApplause();
});
