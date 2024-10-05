Todo App with Go and MongoDB
這是一個使用 Go 和 MongoDB 構建的簡單待辦事項應用程式，並通過 Docker 運行。該應用程式允許用戶添加和刪除待辦事項，並將資料存儲在 MongoDB 中。

專案結構
.\
├── Dockerfile\
├── docker-compose.yml\
├── go.mod\
├── go.sum\
├── main.go\
├── models/\
│   └── todo.go\
└── templates/\
    └── todos.html
    
main.go: 應用程式的主文件，定義了服務器啟動、路由和 MongoDB 連接。
models/todo.go: 定義了 Todo 結構體和與 MongoDB 的交互。
templates/todos.html: 用於顯示待辦事項列表的 HTML 模板。
docker-compose.yml: 定義了 Docker 服務，包括應用程式和 MongoDB。
Dockerfile: 定義了如何構建應用程式的 Docker 映像。
環境要求
Docker: 安裝 Docker 以便運行容器化的應用程式。
Docker Compose: 用於同時運行 Go 應用程式和 MongoDB 服務。
安裝與運行
請按照以下步驟克隆此倉庫並啟動應用程式：

1. 構建並啟動應用程式
首先進入專案的目錄：

cd $專案目錄


然後使用 docker compose 命令來構建並啟動應用程式和 MongoDB：

docker compose up --build
這將啟動兩個 Docker 容器：

todo-app: 包含 Go 應用程式，運行在 http://localhost:3000。
todo-database: 包含 MongoDB 資料庫，運行在 localhost:27017。

2. 開啟應用程式
在瀏覽器中打開 http://localhost:3000，您可以在此添加新的待辦事項和查看現有的事項。

3. 停止應用程式
要停止並移除容器，請使用以下命令：

docker compose down