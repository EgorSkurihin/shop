// Определение страницы, на которой находимся для отображения в навбаре
function checkPathMenu(){
    pathname = window.location.pathname
    switch(pathname) {
        case '/':
            menuItem = $("#nav-home");
            break
        case '/about':
            menuItem = $("#nav-about");
            break
        case '/cart':
            menuItem = $("#nav-cart");
            break
        default:
            return
    }
    menuItem.addClass("active");
}


$(document).ready( () => {
    checkPathMenu();
})