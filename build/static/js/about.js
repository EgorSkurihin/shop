function init() {
    if (localStorage.getItem('cart')){
        cart = JSON.parse(localStorage.getItem('cart'));
        $('#cart-items-number').html(cart.amount);
    }
}


$(document).ready( () => {
    init();
})