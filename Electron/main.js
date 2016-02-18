var app = require("app");
var BrowserWindow = require("browser-window");
var Menu = require("menu");
var mainWindow = null;

app.on("window-all-closed", function(){
    app.quit();
});

app.on("ready", function () {
    mainWindow = new BrowserWindow({
        width: 1280,
        height: 1024,
        'node-integration': false
    });
    // mainWindow.openDevTools();
    mainWindow.loadUrl('http://localhost:8080/');
    mainWindow.on("closed", function () {
        mainWindow =  null;
    });

    // Create the Application's main menu
    var template = [{
        label: "Application",
        submenu: [
            { label: "About enno_go", selector: "orderFrontStandardAboutPanel:" },
            { type: "separator" },
            { label: "Quit", accelerator: "Command+Q", click: function() { app.quit(); }}
        ]}, {
        label: "Edit",
        submenu: [
            { label: "Undo", accelerator: "CmdOrCtrl+Z", selector: "undo:" },
            { label: "Redo", accelerator: "Shift+CmdOrCtrl+Z", selector: "redo:" },
            { type: "separator" },
            { label: "Cut", accelerator: "CmdOrCtrl+X", selector: "cut:" },
            { label: "Copy", accelerator: "CmdOrCtrl+C", selector: "copy:" },
            { label: "Paste", accelerator: "CmdOrCtrl+V", selector: "paste:" },
            { label: "Select All", accelerator: "CmdOrCtrl+A", selector: "selectAll:" }
        ]}
    ];

    Menu.setApplicationMenu(Menu.buildFromTemplate(template));
});