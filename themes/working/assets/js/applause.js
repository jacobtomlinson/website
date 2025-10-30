function updateApplause(applause) {
    let container = document.querySelector('#applause');
    if (!container || !document.querySelector("article")) {
        return;
    }
    container.addEventListener("click", giveApplause);
    container.style.removeProperty('display');
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
