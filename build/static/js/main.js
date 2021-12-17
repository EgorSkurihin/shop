var cart = {
    goods: {},
    amount: 0,
};

function init() {
    // Получение файла goods.json
    $.getJSON("/getGoods", goodsOut);
}

function goodsOut(data) {
    //Вывод товара на страницу
    let out = '';
    for (let key in data) {
        out += 
            '<div class="col-lg-3 col-md-6 mb-4">' +
                '<div class="card bg-light">' +
                    `<img src="${data[key].image}" class="card-img-top" alt="...">` +
                    '<div class="card-body">' +
                    `<h5 class="card-title">${data[key].title}</h5>` +
                    `<p>${data[key].performer}</p>` +
                    '<h4 class="font-weight-bold text-warning">' +
                        `<strong>${data[key].cost}$</strong>` +
                    '</h4>' +
                    `<button class="btn btn-secondary add-to-cart" item-id="${key}">Add to cart</button>` +
                    '</div>' +
                '</div>' +
            '</div>'
    }
    $('#goods-out').html(out);
    $('.add-to-cart').on('click', addToCart);
}

function addToCart(){
    //Добавление товара в корзину на нажтие кнопки
    let id =$(this).attr('item-id');
    if (cart.goods[id] == undefined) {
        cart.goods[id] = 1;
    } else {
        cart.goods[id] += 1;
    }
    cart.amount += 1;
    $('#cart-items-number').html(cart.amount);
    localStorage.setItem('cart', JSON.stringify(cart)); 
}

function loadCart(){
    // Проверка карзины при открытии страницы
    if (localStorage.getItem('cart')){
        cart = JSON.parse(localStorage.getItem('cart'));
        $('#cart-items-number').html(cart.amount);
    }
}

$(document).ready( () => {
    init();
    loadCart();
})