var cart = {
    goods: {},
    amount: 0,
};

function loadCart(){
    // Проверка карзины при открытии страницы
    if (localStorage.getItem('cart')){
        cart = JSON.parse(localStorage.getItem('cart'));
        changeMiniCart();
        if (cart.amount > 0) {
            showCart();
            $('#button-send-email').on('click', sendEmail)
            return
        }
    }
    $('#main-container').html("<h1 class='pt-5 mt-2'>Cart is empty!</h1> <a href='/'>Back to shop</a>")
}

function showCart(){
    // Вывод товаров корзины
    $.getJSON('/getGoods', (goods) =>{

        //Check deleted goods
        for (let id in cart.goods) {
            if (!(id in goods)){
                delete(cart.goods[id]);
                cart.amount -= 1;
                console.log(id)
            }
        }
        changeMiniCart();
        saveCart();

        //display cart
        let out = '';
        let totalCost = 0;
        for (let id in cart.goods){
            out += '<tr>' +
                `<td><img src="${goods[id].image}" width="100" class="img-cart"></td>` +
                `<td><strong>${goods[id].title}</strong><p>${goods[id].performer}</p></td>` +
                '<td>' +
                '<form class="form-inline">' +
                '<div class="input-group ">' +
                '<span class="input-group-btn">' +
                `<button type="button" good-id="${id}" class="btn btn-danger btn-number minus"  data-type="minus"">` +
                '<i class="fas fa-minus"></i>' +
                '</button>' +
                '</span>' +
                `<input type="number" good-id="${id}"" class="form-control input-number goods-number" value="${cart.goods[id]}" min="1" disabled style="width:50px">` +
                '<span class="input-group-btn">' +
                `<button type="button" good-id="${id}" class="btn btn-success btn-number plus" data-type="plus"">` +
                '<i class="fas fa-plus"></i>' +
                '</button>' +
                '</span>' +
                '</div>' +
                '</form>' +
                '<td>' +
                `<button type="button" good-id='${id}' class="btn btn-primary btn-number del-goods">` +
                '<i class="fa fa-trash" aria-hidden="true"></i>' +
                '</button>' +
                '</td>' +
                '</td>' +
                `<td>$${goods[id].cost}</td>` +
                `<td>$${goods[id].cost*cart.goods[id]}</td>` +
                '</tr>'

            totalCost += goods[id].cost * cart.goods[id]

        }
        out += '<tr>' +
                '<td colspan="6">&nbsp;</td>' +
                '<td colspan="4" class="text-right"><strong>Total</strong></td>' +
                `<td>${totalCost}$</td>` +
                '</tr>'
        
        $('#cart-goods').html(out);
        $('#trigger-modal').on('click', triggerModal)
        $('.del-goods').on('click', delGoods);
        $('.goods-number').change(changeGoodsNumber);
        $('.plus').on('click', plusVal);
        $('.minus').on('click', minusVal);
    });
};

function triggerModal(){
    $('#form-total-cost').html(`${cart.amount}`)
}

function plusVal() {
    // Увеличение количества товаров 
    let id = $(this).attr('good-id');
    input = $(`input[good-id=${id}]`);
    input.val(parseInt(input.val(), 10) + 1);
    input.change();
}

function minusVal() {
    // Уменьшение количества товаров
    let id = $(this).attr('good-id');
    let input = $(`input[good-id=${id}]`);
    let val = parseInt(input.val(), 10);
    if (val > 1) {
        input.val(parseInt(input.val(), 10) - 1);
        input.change();
    } 
}

function changeGoodsNumber() {
    // Реакция на изменение количества товаров
    let val =parseInt($(this).val(), 10);
    let id = $(this).attr('good-id');
    cart.goods[id] = val;
    countCartAmount();
    changeMiniCart();
    saveCart();
    showCart();
}

function delGoods() {
    // Удаление товара из корзины
    let id = $(this).attr('good-id');
    cart.amount -= cart.goods[id];
    delete cart.goods[id];
    showCart();
    changeMiniCart();
    saveCart();
}

function countCartAmount() {
    // Подсчет количества товаров в корзине
    let counter = 0
    for (id in cart.goods){
        counter += cart.goods[id]
    }
    cart.amount = counter
}

function changeMiniCart() {
    // Вывод количества товаров корзины
    $('#cart-items-number').html(cart.amount);
}

function saveCart(){
    // Сохранение корзины в localStorage
    localStorage.setItem('cart', JSON.stringify(cart));
}

function sendEmail() {
    var answer = window.confirm("Are you shure?");
    if (answer) {
        uname = $('#user-name').val();
        email = $('#user-email').val();
        tel = $('#user-tel').val();
        if (uname != '' && email != ''){
            if (!$.isEmptyObject(cart.goods)){
                $.post(
                    '/mail',
                    {
                        'name': uname,
                        'email': email,
                        'tel': tel,
                        'cart': JSON.stringify(cart.goods)
                    },
                    () => {
                        cart = {
                            goods: {},
                            amount: 0,
                        };
                        saveCart();
                        changeMiniCart();
                        $('#trigger-modal').click();
                        $('#main-container').html("<h1 class='pt-5 mt-2'>Success! Check your email</h1>");
                    })
                    .fail(() =>{
                        $('#trigger-modal').click();
                        $('#main-container').html("<h1 class='pt-5 mt-2'>Error! Repeat the order</h1>");
                    })
            } else {
                alert('Корзина пуста!')
            }
        } else {
            alert("Поля не заполнены!");
        }
    }
}

$(document).ready( () => {
    loadCart();
})