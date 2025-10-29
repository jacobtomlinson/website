function halloweenTheme() {
    document.querySelectorAll('.halloween').forEach(el => {
        if (el instanceof HTMLElement) el.style.display = '';
    });
}

window.addEventListener("load", function () {
    if (typeof Date === "function") {
        const now = new Date();
        if (now.getMonth() === 9 && now.getDate() >= 25 && now.getDate() <= 31) {
            halloweenTheme();
        }
    }
});
