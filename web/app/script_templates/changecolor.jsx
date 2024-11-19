// 设置动作集和动作名称
var actionSetName = "地图调色"; // 动作集的名称
var actionName = "调整";     // 动作的名称

// 输入和输出文件夹路径
var inputFolderPath = "./";

if (inputFolderPath) {
    var inputFolder = new Folder(inputFolderPath);
    
    // 获取文件夹中所有图像文件
    var files = inputFolder.getFiles(/\.(jpg|jpeg|png|tif|psd)$/i);

    for (var i = 0; i < files.length; i++) {
        var file = files[i];
        
        if (file instanceof File) {
            // 打开图像
            var doc = app.open(file);
            // 执行动作
            app.doAction(actionName, actionSetName);
        }
    }

    alert("批量处理完成！");
} else {
    alert("未选择有效的文件夹！");
}


var xhr = new XMLHttpRequest();
xhr.open("POST", "https://jsonplaceholder.typicode.com/posts", true);  // 假设的 API 地址

// 设置请求头，指定 Content-Type 为 JSON
xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

// 请求完成时的回调函数
xhr.onreadystatechange = function() {
    if (xhr.readyState === 4 && xhr.status === 201) {  // 201 表示请求成功并创建资源
        try {
            // 解析 JSON 数据
            var responseData = JSON.parse(xhr.responseText);  // 将响应文本解析为 JSON 对象
            alert("返回的 JSON 数据: " + JSON.stringify(responseData, null, 2));  // 格式化并显示
        } catch (e) {
            alert("JSON 解析失败: " + e.message);
        }
    } else if (xhr.readyState === 4) {
        alert("请求失败，状态码：" + xhr.status);
    }
};

// 准备要发送的 JSON 数据
var requestData = JSON.stringify({
    title: 'foo',
    body: 'bar',
    userId: 1
});

// 发送 POST 请求
xhr.send(requestData);
