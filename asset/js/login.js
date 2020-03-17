$("#frm-login").submit(function (e) {
    e.preventDefault();
    $.LoadingOverlay("show");
    $frm = $(this);
    $.ajax({
        url : $frm.attr("action"),
        type : $frm.attr("method"),
        data : $frm.serialize(),
        success : function (r) {
            $.LoadingOverlay("hide");
            if(r.code==200){
                toastr.success(r.msg);
                setTimeout(function () { document.location = '/dashboard'; },500);
                //$.LoadingOverlay("hide");
            }else{
                toastr.warning(r.msg);
            }

        },error : function (r) {

            $.LoadingOverlay("hide");
            if (!r.responseJSON.msg){
                toastr.warning(r.responseJSON.msg);
            } else {
                toastr.warning(r);
            }
        }
    });
});


$.LoadingOverlaySetup({
    zIndex  :   999999
});
