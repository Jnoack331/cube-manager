window.addEventListener('load', function () {

    console.log('loaded');
    $fileList = document.querySelector('#file-list');
    $fileInput = document.querySelector("#file");
    $form = document.querySelector("form#upload");

    const onDrag = function(e) {
        e.preventDefault();
        e.stopPropagation();
    };

    $fileList.addEventListener('drag', onDrag)
    $fileList.addEventListener('dragstart', onDrag)
    $fileList.addEventListener('dragend', onDrag)
    $fileList.addEventListener('dragover', onDrag)
    $fileList.addEventListener('dragenter', onDrag)
    $fileList.addEventListener('dragleave', onDrag)
    $fileList.addEventListener('drop', onDrag);

    $fileList.addEventListener('dragover', function() {
        $fileList.classList.add('is-dragover');
    })
    $fileList.addEventListener('dragenter', function() {
        $fileList.classList.add('is-dragover');
    })

    $fileList.addEventListener('dragleave', function() {
        $fileList.classList.remove('is-dragover');
    })
    $fileList.addEventListener('dragend', function() {
        $fileList.classList.remove('is-dragover');
    })
    $fileList.addEventListener('drop', function() {
        $fileList.classList.remove('is-dragover');
    })

    $fileList.addEventListener('drop', function(e) {
        e.preventDefault();
        e.stopPropagation();

        $fileInput.files = e.dataTransfer.files;
        $form.submit();
    });
})
