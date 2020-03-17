let goback = document.getElementById("goback")
goback.onclick = function (e) {
    goBack();
}
function goBack() {
    window.history.back();
}