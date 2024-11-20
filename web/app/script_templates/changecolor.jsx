// 接口地址
var apiUrl = "API_URL";

// 设置动作集和动作名称
// var actionSetName = "地图调色"; // 动作集的名称
// var actionName = "调整";     // 动作的名称
var actionSetName = "ACTION_SET_NAME"; // 动作集的名称
var actionName = "ACTION_NAME";     // 动作的名称

// 输入和输出文件夹路径
// var inputFolderPath = "./";
var inputFolderPath = "INPUT_FOLDER_PATH";

function httpGet(host, path) {
    var socket = new Socket();
    var response = "";

    try {
        // 连接到服务器 (主机名和端口)
        if (socket.open(host)) {
            // 构造 HTTP GET 请求
            var request = "GET " + path + " HTTP/1.1\r\n" +
                          "Host: " + host + "\r\n" +
                          "Connection: close\r\n\r\n";

            // 发送请求
            socket.write(request);

            // 接收响应
            while (!socket.eof) {
                response += socket.read(1024); // 每次读取 1024 字节
            }
        } else {
            throw new Error("Unable to connect to " + host);
        }
    } catch (e) {
        alert("Error: " + e.message);
    } finally {
        socket.close(); // 关闭连接
    }

    return response;
}

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

    // 通知服务器任务完成
    var response = httpGet(apiUrl, "/api/home/taskDone?taskPath="+inputFolderPath);
} else {
    alert("未选择有效的文件夹！");
}