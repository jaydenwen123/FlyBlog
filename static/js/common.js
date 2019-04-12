function showCategory() {
    //如果不设置false的话，会先提交表单，然后导致无法跳转
    window.location.href = "/category/show/"
    return false
}

function showArticle() {
    window.location.href = "/article/show/"
    return false;
}

function goAddCategory() {
    window.location.href = "/category/"
    return false;
}

function goAddArticle() {
    window.location.href = "/article/"
    return false;
}

function showTopBlock() {
    layui.use(['layer', "jquery", "util"], function () { //独立版的layer无需执行这一句
        layer = layui.layer; //独立版的layer无需执行这一句
        $ = layui.jquery;
        util = layui.util
        util.fixbar({
            top: true,
            bar1: false,
            bar2: false,
            css: {right: 100, bottom: 100},
            // bar1: true,
            click: function (type) {
                if (type === "bar2") {
                    window.location.href = "/about"
                }
            }
        });
        // sync=layui.sync;
        //总页数大于页码总数
    });

}

