function init() {
	//Dispalay albums
	showAlbums();
	$('#add-button').on('click', addAlbum);
};

function showAlbums() {
	$.getJSON({
		url: '/getGoods',
	}).done(function (data) {
		let out = '';
        for (let id in data){
            out += '<tr>' +
				`<td>${id}</td>` +
				`<td>${data[id].title}</td>` +
				`<td>${data[id].performer}</td>` +
				`<td>${data[id].cost}$</td>` +
				`<td><img src="${data[id].image}" height="75px" alt="..."></td>` +
				`<td>` +
				`<a class="edit" title="Edit" data-id="${id}" data-bs-toggle="modal" data-bs-target="#editAlbumModal"><i class="material-icons">&#xE254;</i></a>` +
				`<a class="delete" title="Delete" data-id="${id}" data-bs-toggle="modal" data-bs-target="#deleteAlbumModal"><i class="material-icons">&#xE872;</i></a>` +
				`</td>` +
			`</tr>`
		}
		$('#admin-goods').html(out);

		// Кнопки изменения и удаления товара
		$('#delete-button').on('click', deleteAlbum);
		$('#edit-button').on('click', editAlbum);
		
		// Передача в модальное окно id товара
		$('#editAlbumModal').on('show.bs.modal', function(event) {
			let id = $(event.relatedTarget).data('id');
			$('#edit-button').attr('data-id', id)
		})
		$('#deleteAlbumModal').on('show.bs.modal', function(event) {
			let id = $(event.relatedTarget).data('id');
			$('#delete-button').attr('data-id', id)
		})
	})
}

function addAlbum() {
	alb = {
		'title': $('#add-title').val(),
		'performer': $('#add-performer').val(),
		'cost': $('#add-cost').val(),
		'category': $('#add-category').val(),
		'image': $('#add-image').val()
	}
	$.post('/admin', alb)
}

function editAlbum() {
	let id = $(this).attr('data-id');
	changedVals = {
		'title': $('#edit-title').val(),
		'performer': $('#edit-performer').val(),
		'cost': $('#edit-cost').val(),
		'category': $('#edit-category').val(),
		'image': $('#edit-image').val()
	}
	$.ajax({
		url: `/admin/${id}`,
		type: 'PUT',
		data: changedVals,
	  });
}

function deleteAlbum() {
	let id = $(this).attr('data-id');
	$.ajax({
		url: `/admin/${id}`,
		type: 'delete',
		success: function(result) {
			console.log(result);
		},
	  });	
}


$(document).ready( () =>{
    init();
})